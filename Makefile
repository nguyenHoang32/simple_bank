createdb:
	docker exec -it postgres16 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres16 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

startpostgres:
	@if sudo lsof -i :5432; then \
		echo "Killing process on port 5432"; \
		pid=$$(sudo lsof -t -i :5432); \
		if [ -n "$$pid" ]; then \
			echo $$pid; \
			sudo kill -9 $$pid; \
			sleep 2; \
		fi; \
		if lsof -i :5432; then \
			echo "Port 5432 is still in use. Exiting."; \
			exit 1; \
		fi; \
	fi
	docker start postgres16

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb --destination db/mock/store.go simple_bank/db/sqlc Store

.PHONY: createdb dropdb migrateup migratedown test server mock
