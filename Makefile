.SILENT .PHONY: all start

start:
	if test -f .env; then echo "Using .env"; else echo "Creating .env from env.dist" && cp env.dist .env; fi
	docker compose up -d

serve:
	export $$(cat .env | xargs) && cd ./app && go run . serve