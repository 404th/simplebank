postgres:
	docker run --name postgres12 -p 2345:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine 

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrateup:
	migrate -path ./db/migrations/ -database "postgresql://root:secret@localhost:2345/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path ./db/migrations/ -database "postgresql://root:secret@localhost:2345/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres, createdb, dropdb, migrateup, migratedown, test
