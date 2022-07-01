run:
	GOOGLE_APPLICATION_CREDENTIALS=~/.secret/dataflow.json python3 balance.py \
	--input gs://krystal-log-warehouse-v1/arbitrum-transfer-events-v8/log_0000000000_0000000999.avro \
	--bigtable_instance haupc-instance \
	--bigtable_table balance \
	--bigtable_project dev-krystal-wallet \
	--chain_id 42161
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
	--chain_id 2501 \
	--max_num_workers 200 \
	--job_name fantom-n1-standard-1 \
	--machine_type n1-standard-1 \
	--input gs://krystal-log-warehouse-v1/fantom-transfer-events-v8/log_*.avro

rundataflow-aurora:
	GOOGLE_APPLICATION_CREDENTIALS=~/.secret/dataflow.json python3 balance.py \
	--runner DataflowRunner \
	--project dev-krystal-wallet \
	--region asia-southeast1 \
	--staging_location gs://krystal-log-warehouse-v1/temp-dataflow/staging \
	--temp_location gs://krystal-log-warehouse-v1/temp-dataflow/temp \
	--bigtable_instance haupc-instance \
	--bigtable_table balance \
	--bigtable_project dev-krystal-wallet \
	--chain_id 1313161554 \
	--max_num_workers 200 \
	--job_name aurora-n1-standard-1 \
	--machine_type n1-standard-1 \
	--input gs://krystal-log-warehouse-v1/aurora-transfer-events-v8/log_*.avro