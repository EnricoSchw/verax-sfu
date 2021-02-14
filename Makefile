run: init
	go run *.go

build:
	 go build

install:
	go install .

clean:
	go clean -modcache

init:
	go mod download
	go generate ./...

