#!/bin/bash

if [ -f ./.env ]; then
    export $(cat ./.env | xargs)
fi

MIGRATIONS_DIR="internal/migrations/clickhouse"

for migration in $(ls $MIGRATIONS_DIR | sort -V); do
    echo "Applying migration: $migration"
    # echo $CLICKHOUSE_USER
    # echo $CLICKHOUSE_PASSWORD
    # echo $CLICKHOUSE_DB
    # echo "$(cat $MIGRATIONS_DIR/$migration)"
    docker exec clickhouse clickhouse-client --user=$CLICKHOUSE_USER --password=$CLICKHOUSE_PASSWORD --database=$CLICKHOUSE_DB --multiquery "$(cat $MIGRATIONS_DIR/$migration)"
    if [ $? -ne 0 ]; then
        echo "Migration failed: $migration"
        exit 1
    fi
done

echo "All migrations applied successfully."
