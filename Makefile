export GOOSE_DRIVER := postgres
export GOOSE_DBSTRING := user=root password=root host=127.0.0.1 dbname=pos sslmode=disable
export GOOSE_MIGRATION_DIR := internal/app/database/migrations

up:
	goose up

down:
	goose down

drop:
	goose down-to 0

status:
	goose status
