buildamd:
	GOARCH=amd64 go build -o bin-amd64 main.go
runamd:
	GOARCH=amd64 go run .
build:
	go build -o bin-arm64 main.go
run:
	go run main.go
rundataflow:
	GOOGLE_APPLICATION_CREDENTIALS=~/.secret/dataflow.json go run main.go \
	--runner dataflow \
	--project dev-krystal-wallet \
	--region asia-southeast1 \
	--staging_location gs://dev-krystal-wallet/krystal-public-data/test