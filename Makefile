createdb:
	docker exec -it infra-pg-image-1 createdb --username=postgres --owner=postgres simple_bank
dropdb:
	docker exec -it infra-pg-image-1 dropdb --username=postgres simple_bank
migrate:
	docker run -v ./migration:/migration --network host migrate/migrate -path=/migration/ -database "postgres://postgres:postgres@localhost:5439/simple_bank?sslmode=disable" up 1
rollback:
	docker run -v ./migration:/migration --network host migrate/migrate -path=/migration/ -database "postgres://postgres:postgres@localhost:5439/simple_bank?sslmode=disable" down 1
sqlc:
	sqlc generate
test:
	go test -v -count=1 ./...
default:
	$(MAKE) dropdb && $(MAKE) createdb && $(MAKE) migrate && $(MAKE) sqlc

.PHONY: createdb
