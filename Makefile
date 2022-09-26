.PHONY: build
build:
	go build -o ./bin/fcio

.PHONY: run-io
run-io:
	go run main.go clockin
