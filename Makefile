default: all

all: go-rest-api

.PHONY: go-rest-api
go-rest-api:
	@echo "building go-rest-api"
	docker compose up -d