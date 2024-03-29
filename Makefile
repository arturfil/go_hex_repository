include .env

BINARY=serverbin

build:
	go build -o ${BINARY} ./cmd/main.go 

run:
	@echo "run  binary"
	@env DB_HOST=${DB_HOST} ./${BINARY}

stop:
	@echo "stopping backend"
	@-pkill -SIGTERM -f "./${BINARY}"


up:
	docker-compose up -d 

down:
	docker-compose down

init.up:
	@echo "initiating tables..."
	docker exec -i ${DB_CONTAINER_NAME} mysql -u${DB_USER} -p${DB_PASSWORD} ${DB_NAME} < ./migrations/init.up.sql

init.down:
	@echo "Deleting initial creation..."
	docker exec -i ${DB_CONTAINER_NAME} mysql -u${DB_USER} -p${DB_PASSWORD} ${DB_NAME} < ./migrations/init.down.sql

start: build run

restart: stop build start
