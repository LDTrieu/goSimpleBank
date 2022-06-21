postgres:
	docker run --name postgresk12 -d -p 5432:5432 -e POSTGRES_USER=root -e  POSTGRES_PASSWORD=secret postgres:12-alpine
createdb:
	docker exec -it postgresk12 createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it postgresk12 dropdb simple_bank
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
dockerexec:
	docker exec -it postgresk12 psql -U root -d simple_bank
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/LDTrieu/goSimpleBank/db/sqlc Store
.PHONY: postgres createdb dropdb migrateup migratedown sqlc test dockerexec mock