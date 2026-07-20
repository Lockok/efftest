-include .env

DSN=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

app-run:
	go run cmd/main.go

migration-create:
	goose -dir migrations create $(name) sql

migrate-up:
	goose -dir migrations postgres "$(DSN)" up

migrate-down:
	goose -dir migrations postgres "$(DSN)" down

migrate-status:
	goose -dir migrations postgres "$(DSN)" status

deploy:
	docker build -t efftest:dev .
	minikube image load efftest:dev
	helm upgrade dev ./chart
	kubectl rollout status deployment/dev-app

swagger-gen:
	@docker compose run --rm swagger \
		init \
		-g cmd/main.go \
		-o docs \
		--parseInternal \
		--parseDependency

fmt:
	gofmt -w cmd internal

vet:
	go vet ./...

lint:
	golangci-lint run

test:
	go test ./...

build:
	go build -o bin/app ./cmd

check: vet test build