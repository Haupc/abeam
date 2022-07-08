import argparse
import logging
import re

import apache_beam as beam
from apache_beam import pvalue
from apache_beam.io import ReadFromText
from apache_beam.io import WriteToText
from apache_beam.io.gcp.bigtableio import WriteToBigTable
from apache_beam.options.pipeline_options import PipelineOptions
from apache_beam.options.pipeline_options import SetupOptions
from apache_beam.io.avroio import ReadFromAvro

# out gs://krystal-log-warehouse-v1/neg/%s.txt

TAG_1M = '1m'
TAG_3M = '3m'
TAG_5M = '5m'
TAG_15M = '15m'
TAG_30M = '30m'
TAG_45M = '45m'
TAG_1H = '1h'
TAG_2H = '2h'
TAG_3H = '3h'
TAG_4H = '4h'
TAG_1D = '1d'

TAGS = [TAG_1M, TAG_3M, TAG_5M, TAG_15M, TAG_30M,
        TAG_45M, TAG_1H, TAG_2H, TAG_3H, TAG_4H, TAG_1D]
COLUMN_FAMILY_ID = 'data_point'
TABLE_PREFIX = 'trading_view_'


def truncateTime(timeStamp: int, time: str):
    switcher = {
        TAG_1M: 60,
        TAG_3M: 3*60,
        TAG_5M: 5*60,
        TAG_15M: 15*60,
        TAG_30M: 30*60,
        TAG_45M: 45*60,
        TAG_1H: 60*60,
        TAG_2H: 2*60*60,
        TAG_3H: 3*60*60,
        TAG_4H: 4*60*60,
        TAG_1D: 24*60*60,
    }
    timeBlock = switcher.get(time, 0)
    if timeBlock == 0:
        raise Exception('Invalid time block')
    return (timeStamp - (timeStamp % timeBlock))*1000


class Candle:
    def __init__(self, blockTime, logIndex, poolAddress, rev0, rev1, open, close, high, low):
        self.blockTime = blockTime
        self.openLogIndex = logIndex
        self.closeLogIndex = logIndex
        self.poolAddress = poolAddress
        self.rev0 = rev0
        self.rev1 = rev1
        self.open = open
        self.close = close
        self.high = high
        self.low = low

    def key(self, timeInterval: str):
        return "{0}:{1}".format(self.poolAddress, truncateTime(self.blockTime, timeInterval))

    def CbtKey(self, chainId: str, openTime: str):
        return "{0}:{1}:{2}".format(self.poolAddress, chainId, openTime)


def syncEventToCandle(event):
    blockTime = event['blockTime']
    logIndex = event['logIndex']
    poolAddress = event['poolAddress']
    rev0 = int(event['token0Reserve'])/(10**event['token0Decimals'])
    rev1 = int(event['token1Reserve'])/(10**event['token1Decimals'])
    o = rev0/rev1
    return Candle(blockTime, logIndex, poolAddress, rev0, rev1, o, o, o, o)


def mergeCandle(existingCandle: Candle, incomingCandle: Candle):
    if (existingCandle.blockTime == -1):
        return incomingCandle
    if (incomingCandle.blockTime < existingCandle.blockTime
            or (incomingCandle.blockTime == existingCandle.blockTime and incomingCandle.openLogIndex < existingCandle.openLogIndex)):
        existingCandle.open = incomingCandle.open
        existingCandle.openLogIndex = incomingCandle.openLogIndex

    if (incomingCandle.blockTime > existingCandle.blockTime
            or (incomingCandle.blockTime == existingCandle.blockTime and incomingCandle.closeLogIndex > existingCandle.closeLogIndex)):
        existingCandle.close = incomingCandle.close
        existingCandle.closeLogIndex = incomingCandle.closeLogIndex
        existingCandle.rev0 = incomingCandle.rev0
        existingCandle.rev1 = incomingCandle.rev1

    existingCandle.high = max(existingCandle.high, incomingCandle.high)
    existingCandle.low = min(existingCandle.low, incomingCandle.low)
    return existingCandle


class ToKVAndSplitFn(beam.DoFn):
    def process(self, elem):
        candle = syncEventToCandle(elem)
        for tag in TAGS:
            yield pvalue.TaggedOutput(tag, (candle.key(tag), candle))


class CombineCandle(beam.CombineFn):
    def create_accumulator(self):
        return Candle(-1, -1, '', 0, 0, 0, 0, 0, 0)

    def add_input(self, accumulator, input):
        if (accumulator.blockTime == -1):
            return input
        candle = accumulator
        return mergeCandle(candle, input)

    def merge_accumulators(self, accumulators):
        # accumulators = [candle1, candle2, candle3,...]
        initCanlde = Candle(-1, -1, '', 0, 0, 0, 0, 0, 0)
        for candle in accumulators:
            initCanlde = mergeCandle(initCanlde, candle)
        return initCanlde

    def extract_output(self, accumulator):
        return accumulator


class ToCbtRowFn(beam.DoFn):
    def __init__(self, chainId: str):
        self.chainId = chainId

    def process(self, kv):

        from google.cloud.bigtable import row
        from datetime import datetime

        candle = kv[1]
        openTimeStr = kv[0].split(':')[1]
        openTime = datetime.utcfromtimestamp(int(openTimeStr)/1000)
        direct_row = row.DirectRow(
            row_key=candle.CbtKey(self.chainId, openTimeStr))
        direct_row.set_cell(
            COLUMN_FAMILY_ID,
            'open',
            str(candle.open),
            timestamp=openTime)
        direct_row.set_cell(
            COLUMN_FAMILY_ID,
            'close',
            str(candle.close),
            timestamp=openTime)
        direct_row.set_cell(
            COLUMN_FAMILY_ID,
            'high',
            str(candle.high),
            timestamp=openTime)
        direct_row.set_cell(
            COLUMN_FAMILY_ID,
            'low',
            str(candle.low),
            timestamp=openTime)
        direct_row.set_cell(
            COLUMN_FAMILY_ID,
            'rev0',
            str(candle.rev0),
            timestamp=openTime)
        direct_row.set_cell(
            COLUMN_FAMILY_ID,
            'rev1',
            str(candle.rev1),
            timestamp=openTime)
        yield direct_row


def main(argv=None, save_main_session=True):
    """Main entry point; defines and runs the wordcount pipeline."""

    parser = argparse.ArgumentParser()
    parser.add_argument(
        '--bigtable_instance',
        dest='bigtable_instance',
        required=True,
        help='bigtable_instance to write.')
    # parser.add_argument(
    #     '--bigtable_table',
    #     dest='bigtable_table',
    #     required=True,
    #     help='bigtable_table to write.')
    parser.add_argument(
        '--bigtable_project',
        dest='bigtable_project',
        required=True,
        help='bigtable_project to write.')
    parser.add_argument(
        '--input',
        dest='input',
        required=True,
        help='Input file to process.')
    parser.add_argument(
        '--chain_id',
        dest='chain_id',
        required=True,
        help='chain_id to write.')

    known_args, pipeline_args = parser.parse_known_args(argv)

    print("---------------------> pipeline_args: {}".format(pipeline_args))

    # We use the save_main_session option because one or more DoFn's in this
    # workflow rely on global context (e.g., a module imported at module level).
    pipeline_options = PipelineOptions(pipeline_args)
    print("---------------------> pipeline_options: {}".format(pipeline_options))
    pipeline_options.view_as(
        SetupOptions).save_main_session = save_main_session
    with beam.Pipeline(options=pipeline_options) as p:  # auto call p.run()

        # Read the text file[pattern] into a PCollection.
        splitedPCols = (p | 'readFromAvro' >> ReadFromAvro(known_args.input)
                        | 'parseAndSplitAvro' >> (beam.ParDo(ToKVAndSplitFn())).with_outputs(
                            TAG_1M,
                            TAG_3M,
                            TAG_5M,
                            TAG_15M,
                            TAG_30M,
                            TAG_45M,
                            TAG_1H,
                            TAG_2H,
                            TAG_3H,
                            TAG_4H,
                            TAG_1D))
        for tag in TAGS:
            cbtRow = (splitedPCols[tag] | 'combineCandle-%s' % tag >> beam.CombinePerKey(
                CombineCandle())
                | 'toCbtRow-%s' % tag >> beam.ParDo(ToCbtRowFn(known_args.chain_id))
            )
            cbtRow | 'writeToBigTable-%s' % tag >> WriteToBigTable(project_id=known_args.bigtable_project,
                                                                   instance_id=known_args.bigtable_instance,
                                                                   table_id=TABLE_PREFIX + tag)


if __name__ == '__main__':
    logging.getLogger().setLevel(logging.INFO)
    main()
