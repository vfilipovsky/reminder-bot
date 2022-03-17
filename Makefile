GO=$(shell which go)

.DEFAULT_GOAL := build-and-run

install:
	${GO} get ./...

dev:
	${GO} run .

build:
	${GO} build -o .bin/build .

build-and-run:
	${GO} build -o .bin/build . && ./.bin/build

fmt:
	${GO} fmt ./...

test:
	${GO} test -v ./...