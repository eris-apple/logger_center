.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: create_migraion
create_migration:
	migrate create -ext sql -dir migrations <name>

.PHONY: upload_migraion
upload_migration:
	migrate -path migrations -database "postgres://root:password@localhost:5432/logger_center?sslmode=disable" up