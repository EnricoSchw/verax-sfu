run:
	go run *.go

build:
	 go build

build-linux:
	GOOS=linux GOARCH=amd64 go build

install:
	go install .

clean:
	go clean -modcache

init:
	go mod download
	go generate ./...

