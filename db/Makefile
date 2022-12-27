POD = $(shell kubectl get pods | grep postgres | awk '{ print $$1 }')
DB = simple_bank

createdb:
	kubectl exec -it $(POD) -- createdb --username=admin --owner=admin $(DB)

dropdb:
	kubectl exec -it $(POD) -- dropdb --username=admin $(DB)

migrate-up:
	migrate -path ./migration -database "postgresql://admin:password@localhost:5432/$(DB)?sslmode=disable" --verbose up

migrate-down:
	migrate -path ./migration -database "postgresql://admin:password@localhost:5432/$(DB)?sslmode=disable" --verbose down


sqlc:
	sqlc generate

install-deps:
	sudo snap install sqlc

.PHONY: createdb dropdb migrate-up migrate-down sqlc