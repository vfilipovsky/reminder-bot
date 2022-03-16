GO=$(shell which go)

.DEFAULT_GOAL := build-and-run

install:
	${GO} get ./...

dev:
	${GO} run cmd/server/main.go

build:
	${GO} build -o .bin/build ./cmd/server/main.go

build-and-run:
	${GO} build -o .bin/build ./main.go && ./.bin/build

fmt:
	${GO} fmt ./...

test:
	${GO} test -v ./...