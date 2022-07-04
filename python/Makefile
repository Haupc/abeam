chain_name = aurora
chain_id = 1313161554
run:
	GOOGLE_APPLICATION_CREDENTIALS=~/.secret/dataflow.json python3 negative.py \
	--bigtable_instance haupc-instance \
	--bigtable_table balance \
	--bigtable_project dev-krystal-wallet \
	--chain_id $(chain_id) \
	--input gs://krystal-log-warehouse-v1/$(chain_name)-transfer-events-v8/log_0000000000_0000000999.avro \
	--job_name $(chain_name)
rundataflow:
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
	--input gs://krystal-log-warehouse-v1/$(chain_name)-transfer-events-v8/log_*.avro \
