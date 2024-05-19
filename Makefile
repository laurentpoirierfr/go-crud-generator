generate:
	go run cmd/main.go schemas/schema-test.sql repository internal/repository

test:
	go test ./... -v
	