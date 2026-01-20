include .env
export

help:
	@echo "Available commands:"
	@echo "  make up       	 - Start services"
	@echo "  make down     	 - Stop services"
	@echo "  make db       	 - Connect to database"
	@echo "  make redis    	 - Connect to Redis"
	@echo "  make up-build 	 - Build with flag --build"

up-build:
	docker-compose up --build -d

up:
	docker-compose up -d

down:
	docker-compose down

db:
	docker-compose exec postgres psql -U postgres -d messenger

redis:
	docker-compose exec redis redis-cli

migrate-up:
	migrate -path migrations -database ${DATA_BASE_URL} up

migrate-down:
	migrate -path migrations -database ${DATA_BASE_URL} down

