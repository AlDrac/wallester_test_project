run: build start

build:
	go build -o ./bin -v ./

start:
	./bin/api

run-tests:
	go test -v -race -timeout 30s ./tests/...

.DEFAULT_GOAL := run
.PHONY: run build start run-tests