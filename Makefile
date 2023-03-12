export GO11MODULE=true

ifneq (,$(wildcard .env))
	include .env
	export
endif

.PHONY: run-local
run-local:
	go build -o dist/app main.go && dist/app

build:
	docker-compose up -d --build