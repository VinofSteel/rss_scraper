.DEFAULT_GOAL := run

ifneq (,$(wildcard ./.env))
    include .env
    export
endif

PG_CONN_STRING=postgres://$(PGUSER):$(PGPASSWORD)@$(PGHOST):$(PGPORT)/$(PGDATABASE)?sslmode=disable

fmt:
	gofmt -w .
.PHONY: fmt

test: fmt
	go test ./... -count=1
.PHONY: test

run: fmt
	go build && rssscraper.exe
.PHONY: run

migrate-up:
	goose -dir sql/schema postgres "$(PG_CONN_STRING)" up
.PHONY: migrate-up

migrate-down:
	goose -dir sql/schema postgres "$(PG_CONN_STRING)" down
.PHONY: migrate-down

migrate-status:
	goose -dir sql/schema postgres "$(PG_CONN_STRING)" status
.PHONY: migrate-status

migrate-create:
	goose -dir sql/schema create $(name) sql
.PHONY: migrate-create
