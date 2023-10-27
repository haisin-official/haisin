.PHONY: all
all: dev

.PHONY: prod
prod:
	docker compose -f compose.yaml -f ./docker/compose.prod.yaml up -d --build

.PHONY: dev
dev:
	docker compose -f compose.yaml -f ./docker/compose.dev.yaml up -d --build

.PHONY: bash
bash:
	docker compose exec frontend bash

.PHONY: down
down:
	docker compose down
