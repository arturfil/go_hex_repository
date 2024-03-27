include .env

BINARY=serverbin

build:
	go build -o ${BINARY} ./cmd/main.go 

run:
	@echo "run  binary"
	@env DB_HOST=${DB_HOST} ./${BINARY}

up:
	docker-compose up -d 

start: build run

