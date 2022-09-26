.PHONY: build
build:
	go build -o ./bin/fcio cmd/f-clock-io/main.go

.PHONY: run-io
run-io:
	go run main.go clockin
