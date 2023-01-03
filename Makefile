.SILENT .PHONY: all start serve migrate test.init test.integration

start.env.dev:
	if test -f .env; then echo "Using .env"; else echo "Creating .env from env.dist" && cp env.dist .env; fi
	docker compose up -d

start.env.test:
	docker compose --env-file .test.env up -d

serve:
	export $$(cat .env | xargs) && cd ./app && go run . serve

migrate:
	docker compose up migrations

test.unit:
	cd ./app && METRIC_DB_DISABLED=1 go test ./...

test.integration: start.env.test
	export $$(cat .test.env | xargs) && cd ./app && go test --tags=integration.test ./...
	docker compose stop