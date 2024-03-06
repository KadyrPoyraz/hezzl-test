run.app:
	go run cmd/main.go

run.all:
	docker-compose -f docker/docker-compse.yml up -d

run.db:
	docker-compose -f docker/docker-compse.yml up -d db

run.nuts:
	docker-compose -f docker/docker-compse.yml up -d nats

run.clickhouse:
	docker-compose -f docker/docker-compse.yml up -d clickhouse

run.initTestDB:
	docker exec -it db psql -U user db -c "CREATE DATABASE test"
	goose -dir internal/migrations/postgresql -allow-missing postgres "host=localhost port=1337 user=user password=228 dbname=test sslmode=disable" up

run.test.db:
	docker-compose -f docker/docker-compse.yml up -d db-test

run.db.migrate.up:
	goose -dir internal/migrations/postgresql -allow-missing postgres "host=localhost port=1337 user=user password=228 dbname=db sslmode=disable" up

run.db.migrate.down:
	goose -dir internal/migrations/postgresql -allow-missing postgres "host=localhost port=1337 user=user password=228 dbname=db sslmode=disable" down

stop.db:
	docker-compose -f docker/docker-compse.yml  db
