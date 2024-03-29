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


# prod-createdb:
# 	kubectl exec -n db -it $(POSTGRES_POD_NAME) -- createdb --username=$(DB_USER) --owner=$(DB_USER) $(DB_NAME)

# prod-dropdb:
# 	kubectl exec -n db -it $(POSTGRES_POD_NAME) -- dropdb --username=$(DB_USER) $(DB_NAME)

# Required Dependency - https://github.com/golang-migrate/migrate
# To install - https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
create-migration:
	test $(MIGRATION_NAME) != "default" && migrate create -ext sql -dir db/migration -seq $(MIGRATION_NAME)

# migrate-up:
# 	migrate -path ./db/migration -database "$(DB_DRIVER)://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" --verbose up

# migrate-up1:
# 	migrate -path ./db/migration -database "$(DB_DRIVER)://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" --verbose up 1

# migrate-down:
# 	migrate -path ./db/migration -database "$(DB_DRIVER)://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" --verbose down

# migrate-down1:
# 	migrate -path ./db/migration -database "$(DB_DRIVER)://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" --verbose down 1

sqlc:
	sqlc generate -f ./db/sqlc.yaml

test:
	go test -v -cover -count 1 ./...

prod-psql:
	kubectl exec -it -n db $(POSTGRES_POD_NAME) -- psql -h $(DB_HOST) -U $(DB_USER) -d $(DB_NAME)

prod-tail-pod:
	kubectl logs -f -n api-server $(API_SERVER_POD_NAME)

prod-deploy:
	earthly +final-image --ENV="prod"
	kind --name=jtrack load docker-image backend:latest
	kubectl -n backend rollout restart deployment backend-deployment


LOCAL_DB_DRIVER = postgresql
LOCAL_DB_NAME = db
LOCAL_DB_USER = admin
LOCAL_DB_PASS = password
LOCAL_DB_HOST = localhost
LOCAL_DB_PORT = 6666
LOCAL_MIGRATION_NAME = default
LOCAL_DB_CONTAINER = local-postgres

LOCAL_MINIO_ACCESS_KEY = root
LOCAL_MINIO_SECRET_KEY = password
LOCAL_MINIO_CONTAINER = local-minio

local-check-db:
	curl $(LOCAL_DB_HOST):$(LOCAL_DB_PORT)

local-db-drop:
	docker container exec $(LOCAL_DB_CONTAINER) dropdb --username=$(DB_USER) $(DB_NAME)

local-db:
	./resources/build/start-local-db.sh $(LOCAL_DB_CONTAINER) $(LOCAL_DB_PORT) $(LOCAL_DB_NAME) $(LOCAL_DB_USER) $(LOCAL_DB_PASS)
	
local-minio:
	@./resources/build/start-local-minio.sh $(LOCAL_MINIO_CONTAINER) $(LOCAL_MINIO_ACCESS_KEY) $(LOCAL_MINIO_SECRET_KEY)

local-destroy-minio:
	docker rm --force $(LOCAL_MINIO_CONTAINER)

local-destroy-db:
	docker rm --force $(LOCAL_DB_CONTAINER)

local-psql:
	docker container exec -it $(LOCAL_DB_CONTAINER) /bin/bash -c "psql -h $(LOCAL_DB_HOST) -U $(LOCAL_DB_USER) -d $(LOCAL_DB_NAME)"

local: local-db local-minio
	go run main.go

local-migrate-up:
	migrate -path ./db/migration -database "$(LOCAL_DB_DRIVER)://$(LOCAL_DB_USER):$(LOCAL_DB_PASS)@$(LOCAL_DB_HOST):$(LOCAL_DB_PORT)/$(LOCAL_DB_NAME)?sslmode=disable" --verbose up

local-migrate-up1:
	migrate -path ./db/migration -database "$(LOCAL_DB_DRIVER)://$(LOCAL_DB_USER):$(LOCAL_DB_PASS)@$(LOCAL_DB_HOST):$(LOCAL_DB_PORT)/$(LOCAL_DB_NAME)?sslmode=disable" --verbose up 1

local-migrate-down:
	migrate -path ./db/migration -database "$(LOCAL_DB_DRIVER)://$(LOCAL_DB_USER):$(LOCAL_DB_PASS)@$(LOCAL_DB_HOST):$(LOCAL_DB_PORT)/$(LOCAL_DB_NAME)?sslmode=disable" --verbose down

local-migrate-down1:
	migrate -path ./db/migration -database "$(LOCAL_DB_DRIVER)://$(LOCAL_DB_USER):$(LOCAL_DB_PASS)@$(LOCAL_DB_HOST):$(LOCAL_DB_PORT)/$(LOCAL_DB_NAME)?sslmode=disable" --verbose down 1