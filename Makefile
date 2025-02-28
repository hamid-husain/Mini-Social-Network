include .env

APP_NAME=project
PKG=./...
MIGRATIONS_DIR=db/migrations

GOLANGCI_LINT=golangci-lint
GOOSE=goose

install:
	go mod tidy
	go mod download

build:
	go build -o $(APP_NAME) main.go

run: build
	./$(APP_NAME)

lint:
	$(GOLANGCI_LINT) run

fmt:
	go fmt $(PKG)

test:
	go test -v $(PKG)
