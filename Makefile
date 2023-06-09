
postgres:
	docker run --name postgres15 --network ticketx-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root ticketx
dropdb:
	docker exec -it postgres15 dropdb ticketx
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/ticketx?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/ticketx?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
aws:
	aws secretsmanager get-secret-value --secret-id ticketx --query SecretString --output text | jq -r 'to_entries|map("\(.key)=\(.value)")|.[]' > app.env
.PHONY: createdb, dropdb, postgres,migrateup,migratedown,sqlc,test,server,aws