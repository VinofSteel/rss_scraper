.DEFAULT_GOAL := build

ifneq (,$(wildcard ./.env))
    include .env
    export
endif

PG_CONN_STRING=postgres://$(PGUSER):$(PGPASSWORD)@$(PGHOST):$(PGPORT)/$(PGDATABASE)?sslmode=disable

fmt:
	gofmt -w .
.PHONY: fmt

run: fmt
	air
.PHONY: run

build: m-up
	go build && rssscraper.exe
.PHONY: build

m-up:
	goose -dir sql/schema postgres "$(PG_CONN_STRING)" up
.PHONY: m-up

m-down:
	goose -dir sql/schema postgres "$(PG_CONN_STRING)" down
.PHONY: m-down

m-status:
	goose -dir sql/schema postgres "$(PG_CONN_STRING)" status
.PHONY: m-status

m-redo:
	goose -dir sql/schema redo $(name) sql
.PHONY: m-redo
