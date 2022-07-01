run:
	GOOGLE_APPLICATION_CREDENTIALS=~/.secret/dataflow.json python3 balance.py \
	--input gs://krystal-log-warehouse-v1/arbitrum-transfer-events-v8/log_0000000000_0000000999.avro \
	--bigtable_instance haupc-instance \
	--bigtable_table balance \
	--bigtable_project dev-krystal-wallet \
	--chain_id 42161
chain_name = fantom
chain_id = 2522
rundataflow:
	GOOGLE_APPLICATION_CREDENTIALS=~/.secret/dataflow.json python3 balance.py \
	--runner DataflowRunner \
	--project dev-krystal-wallet \
	--region asia-southeast1 \
	--staging_location gs://krystal-log-warehouse-v1/temp-dataflow/staging \
	--temp_location gs://krystal-log-warehouse-v1/temp-dataflow/temp \
	--bigtable_instance haupc-instance \
	--bigtable_table balance \
	--bigtable_project dev-krystal-wallet \
	--chain_id $(chain_id) \
	--input gs://krystal-log-warehouse-v1/$(chain_name)-transfer-events-v8/log_*.avro \
	--max_num_workers 200 \
	--job_name $(chain_name)-n1-standard-2 \
	--machine_type n1-standard-2