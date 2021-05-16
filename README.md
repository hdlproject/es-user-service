# ES User Service
## Description
## How to Run
## DB Migration
```shell script
$ migrate create -ext sql -dir db/migrations -seq <migration_name>
$ migrate -database "postgresql://postgres:postgres@localhost:5433/es-user-service?sslmode=disable" -path db/migrations up
```
