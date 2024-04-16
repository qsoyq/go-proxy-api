.PHONY: default precommit build

default: precommit

install:
	@pip install pre-commit

build:
	@go build -o bin/app src/main.go
	@chmod +x bin/app

precommit:
	@pre-commit install
	@pre-commit run --all-file
