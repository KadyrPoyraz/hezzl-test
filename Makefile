run.app:
	go run cmd/main.go

run.db:
	docker-compose -f docker/docker-compse.yml up -d db

run.test.db:
	docker-compose -f docker/docker-compse.yml up -d db-test

run.db.migrate.up:
	goose -dir internal/migrations -allow-missing postgres "host=localhost port=1337 user=user password=228 dbname=db sslmode=disable" up

run.db.migrate.down:
	goose -dir internal/migrations -allow-missing postgres "host=localhost port=1337 user=user password=228 dbname=db sslmode=disable" down

stop.db:
	docker-compose -f docker/docker-compse.yml  db
