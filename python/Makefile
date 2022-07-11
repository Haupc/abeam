chain_name = arbitrum
chain_id = 42161
date_time=`date +'%H-%M-%S-%d-%m-%Y'`
test:
	echo $(date_time)
neg:
	GOOGLE_APPLICATION_CREDENTIALS=~/.secret/dataflow.json python3 negative.py \
	--bigtable_instance haupc-instance \
	--bigtable_table balance \
	--bigtable_project dev-krystal-wallet \
	--chain_id $(chain_id) \
	--input gs://krystal-log-warehouse-v1/$(chain_name)-transfer-events-v8/log_0000000000_0000000999.avro \
	--job_name $(chain_name)
neg-df:
	GOOGLE_APPLICATION_CREDENTIALS=~/.secret/dataflow.json python3 negative.py \
	--runner DataflowRunner \
	--project dev-krystal-wallet \
	--region asia-southeast1 \
	--staging_location gs://krystal-log-warehouse-v1/temp-dataflow/staging \
	--temp_location gs://krystal-log-warehouse-v1/temp-dataflow/temp \
	--bigtable_instance haupc-instance \
	--bigtable_table balance \
	--bigtable_project dev-krystal-wallet \
	--max_num_workers 200 \
	--chain_id $(chain_id) \
	--input gs://krystal-log-warehouse-v1/$(chain_name)-transfer-events-v8/log_*.avro

tradingview:
	GOOGLE_APPLICATION_CREDENTIALS=~/.secret/dataflow.json python3 tradingview.py \
	--bigtable_instance haupc-instance \
	--bigtable_table balance \
	--bigtable_project dev-krystal-wallet \
	--chain_id $(chain_id) \
	--input gs://krystal-log-warehouse-v1/$(chain_name)-sync-events-v10/log_0000219000_0000219999.avro \
	--job_name $(chain_name)
tradingview-df:
	GOOGLE_APPLICATION_CREDENTIALS=~/.secret/dataflow.json python3 tradingview.py \
	--runner DataflowRunner \
	--project dev-krystal-wallet \
	--region asia-southeast1 \
	--staging_location gs://krystal-log-warehouse-v1/temp-dataflow/staging \
	--temp_location gs://krystal-log-warehouse-v1/temp-dataflow/temp \
	--bigtable_instance haupc-instance \
	--bigtable_project dev-krystal-wallet \
	--max_num_workers 200 \
	--chain_id $(chain_id) \
	--job_name $(chain_name)-tradingview-$(date_time) \
	--input gs://krystal-log-warehouse-v1/$(chain_name)-sync-events-v10/log_*.avro