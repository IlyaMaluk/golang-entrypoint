DB_URL="postgres://user:password@localhost:5432/school_db?sslmode=disable"
MIGRATIONS_DIR=migrations

.PHONY: help build run test swagger migrate-up migrate-down docker-up docker-down clean

help:
	@echo "Доступні команди:"
	@echo "  make build         - Зібрати бінарний файл"
	@echo "  make run           - Запустити сервер локально (без Docker)"
	@echo "  make test          - Запустити всі тести"
	@echo "  make migrate-up    - Накотити міграції бази даних"
	@echo "  make migrate-down  - Відкотити всі міграції бази даних"
	@echo "  make docker-up     - Підняти проєкт у Docker Compose"
	@echo "  make docker-down   - Зупинити Docker Compose та видалити томи"
	@echo "  make clean         - Видалити зібрані файли"

build:
	go build -o bin/server cmd/server/main.go

run:
	go run cmd/server/main.go

test:
	go test -v ./...

migrate-up:
	migrate -path $(MIGRATIONS_DIR) -database $(DB_URL) up

migrate-down:
	migrate -path $(MIGRATIONS_DIR) -database $(DB_URL) down -all

docker-up:
	docker compose -f build/docker-compose.yml up --build

docker-down:
	docker compose -f build/docker-compose.yml down -v

clean:
	rm -rf bin/