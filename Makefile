.PHONY: build
build:
	docker-compose build

.PHONY: up
up:
	docker-compose up

.PHONY: upd
upd:
	docker-compose up -d

.PHONY: down
down:
	docker-compose down

.PHONY: exec
exec:
	docker-compose exec sgo bash

.PHONY: gobuild
gobuild:
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o ./bin/fcio cmd/f-clock-io/main.go

.PHONY: m-gobuild
m-gobuild:
	go build -o ./bin/fcio cmd/f-clock-io/main.go

.PHONY: cpcfg
cpcfg:
	cp config.sample.yml config.yml
	cp config.yml ./bin/config.yml

.PHONY: run-in
run-in:
	go run cmd/f-clock-io/main.go clockin

.PHONY: run-out
run-out:
	go run cmd/f-clock-io/main.go clockout

.PHONY: run-sb
run-sb:
	go run cmd/f-clock-io/main.go startbreak
	
.PHONY: run-fb
run-fb:
	go run cmd/f-clock-io/main.go finishbreak

.PHONY: run-fix
run-fix:
	go run cmd/f-clock-io/main.go fixmemo
