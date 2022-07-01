import argparse
import logging
import re

import apache_beam as beam
from apache_beam.io import ReadFromText
from apache_beam.io import WriteToText
from apache_beam.io.gcp.bigtableio import WriteToBigTable
from apache_beam.options.pipeline_options import PipelineOptions
from apache_beam.options.pipeline_options import SetupOptions
from apache_beam.io.avroio import ReadFromAvro

# out gs://krystal-log-warehouse-v1/neg/%s.txt


class Balance:
    def __init__(self, address, token, tokenId, tokenType, amount):
        self.address = address
        self.token = token
        self.tokenId = tokenId
        self.tokenType = tokenType
        self.amount = amount

    def key(self, chainId: str):
        return "%s:%s:%s:%s:%s" % (self.address, chainId, self.token, self.tokenId, self.tokenType)


class ToCbtRowFn(beam.DoFn):
    def process(self, kv):

        from google.cloud.bigtable import row
        import datetime

        direct_row = row.DirectRow(row_key=kv[0])
        direct_row.set_cell(
            'balance',
            'amount',
            str(kv[1]),
            timestamp=datetime.datetime.now())

        yield direct_row


class ToKvBalanceFn(beam.DoFn):
    def __init__(self, chainId: str):
        self.chainId = chainId

    def process(self, elem):
        balance1 = Balance(elem['to'], elem['token'], elem['tokenId'],
                           elem['eventType'], int(elem['amount']))
        if balance1.address != '0x0000000000000000000000000000000000000000':
            yield balance1.key(self.chainId), balance1.amount

        balance2 = Balance(elem['from'], elem['token'],
                           elem['tokenId'], elem['eventType'], -int(elem['amount']))
        if balance2.address != '0x0000000000000000000000000000000000000000':
            yield balance2.key(self.chainId), balance2.amount


def main(argv=None, save_main_session=True):
    """Main entry point; defines and runs the wordcount pipeline."""

    parser = argparse.ArgumentParser()
    parser.add_argument(
        '--bigtable_instance',
        dest='bigtable_instance',
        required=True,
        help='bigtable_instance to write.')
    parser.add_argument(
        '--bigtable_table',
        dest='bigtable_table',
        required=True,
        help='bigtable_table to write.')
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
        recordPCol = (p | 'readFromAvro' >> ReadFromAvro(known_args.input)
                        | 'parseAvro' >> (beam.ParDo(ToKvBalanceFn(known_args.chain_id)))
                        | 'sumBalance' >> beam.CombinePerKey(sum)
                        | 'toCbtRow' >> beam.ParDo(ToCbtRowFn())
                      )
        recordPCol | 'writeToBigTable' >> WriteToBigTable(project_id=known_args.bigtable_project,
                                                          instance_id=known_args.bigtable_instance,
                                                          table_id=known_args.bigtable_table)


if __name__ == '__main__':
    logging.getLogger().setLevel(logging.INFO)
    main()
