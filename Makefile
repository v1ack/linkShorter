PG_CONTAINER := $(shell docker ps | grep postgres | cut -d ' ' -f 1)

generate-grpc-gateway:
	buf generate

generate-sqlc:
	sqlc generate

generate-migration:
	goose -dir sql/migrations create migration sql

migrate:
	goose -dir sql/migrations postgres "dbname=shorter password=admin user=admin sslmode=disable" up

build:
	go build -o shorter_cli ./cmd

dev:
	docker-compose -f .\docker-compose.dev.yaml up --build app

clean_test_db:
	docker exec -it $(PG_CONTAINER) psql -U admin postgres -w -c "drop database if exists shorter_test"
	docker exec -it $(PG_CONTAINER) psql -U admin postgres -w -c "create database shorter_test"


test: clean_test_db
	go test ./...