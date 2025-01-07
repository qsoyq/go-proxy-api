.PHONY: default precommit build swag run

default: precommit

run:
	@go run src/main.go

install:
	@pip install pre-commit
	@go install github.com/swaggo/swag/cmd/swag@latest

build:
	@go build -o bin/app src/main.go
	@chmod +x bin/app

precommit:
	@pre-commit install
	@pre-commit run --all-file

swag:
	@swag init -d src -o src/docs

test:
	@go test ./src...
