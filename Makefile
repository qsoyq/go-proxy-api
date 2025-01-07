.PHONY: default precommit build generate

default: precommit

install:
	@pip install pre-commit
	@go install github.com/swaggo/swag/cmd/swag@latest

build:
	@go build -o bin/app src/main.go
	@chmod +x bin/app

precommit:
	@pre-commit install
	@pre-commit run --all-file

generate:
	@cd src 
	@swag init 
	@cd ..
