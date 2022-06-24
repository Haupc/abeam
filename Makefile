buildamd:
	GOARCH=amd64 go build -o bin-amd64 main.go
runamd:
	GOARCH=amd64 go run .
build:
	go build -o bin-arm64 main.go
run:
	go run main.go