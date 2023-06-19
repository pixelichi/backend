.PHONY: createdb dropdb migrate-up migrate-up1  migrate-down migrate-down1 sqlc test server create-migration shell-postgres

POSTGRES_POD_NAME = "$(shell kubectl -n db get pods | grep db | awk '{ print $$1 }')"
API_SERVER_POD_NAME = "$(shell kubectl -n api-server get pods | grep api-server | awk '{ print $$1 }')"

DB_DRIVER = postgresql
DB_NAME = db
DB_USER = admin
DB_PASS = password
DB_HOST = localhost
DB_PORT = 5432
MIGRATION_NAME = "default"
ENV = "dev-cluster"


createdb:
	kubectl exec -n db -it $(POSTGRES_POD_NAME) -- createdb --username=$(DB_USER) --owner=$(DB_USER) $(DB_NAME)

dropdb:
	kubectl exec -n db -it $(POSTGRES_POD_NAME) -- dropdb --username=$(DB_USER) $(DB_NAME)

# Required Dependency - https://github.com/golang-migrate/migrate
# To install - https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
create-migration:
	test $(MIGRATION_NAME) != "default" && migrate create -ext sql -dir db/migration -seq $(MIGRATION_NAME)

migrate-up:
	migrate -path ./db/migration -database "$(DB_DRIVER)://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" --verbose up

migrate-up1:
	migrate -path ./db/migration -database "$(DB_DRIVER)://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" --verbose up 1

migrate-down:
	migrate -path ./db/migration -database "$(DB_DRIVER)://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" --verbose down

migrate-down1:
	migrate -path ./db/migration -database "$(DB_DRIVER)://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" --verbose down 1

sqlc:
	sqlc generate -f ./db/sqlc.yaml

test:
	go test -v -cover -count 1 ./...

psql:
	kubectl exec -it -n db $(POSTGRES_POD_NAME) -- psql -h $(DB_HOST) -U $(DB_USER) -d $(DB_NAME)

tail-pod:
	kubectl logs -f -n api-server $(API_SERVER_POD_NAME)

server:
	go run main.go

kind-dev:
	earthly +final-image --ENV=$(ENV)
	kind load docker-image backend:latest
	kubectl -n api-server rollout restart deployment api-server-deployment