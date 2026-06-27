DB_DSN=postgres://postgres:postgres@localhost:5433/effdb?sslmode=disable

app-run:
	go run cmd/api/main.go

goose-download:
	go run github.com/pressly/goose/v3/cmd/goose@latest -h

migrate-up:
	goose -dir migrations postgres "$(DB_DSN)" up

migrate-down:
	goose -dir migrations postgres "$(DB_DSN)" down