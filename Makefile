run:
	go run ./cmd/api

migrate:
	go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate \
		-path ./migrations \
		-database "postgres://erp:erp123@localhost:5432/erp_dev?sslmode=disable" up

migrate-down:
	go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate \
		-path ./migrations \
		-database "postgres://erp:erp123@localhost:5432/erp_dev?sslmode=disable" down

test:
	go test ./...
