include .env

export $(shell sed 's/=.*//' .env)

DATABASE_URL=postgresql://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

run: 
	go run cmd/*.go

test: 
	go test ./... -v

migrate-up:
	migrate -path=database/migrations -database "$(DATABASE_URL)" -verbose up

migrate-down:
	migrate -path=database/migrations -database "$(DATABASE_URL)" -verbose down
