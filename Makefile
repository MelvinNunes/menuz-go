# Makefile

# Default migration name if not provided
DEFAULT_NAME := unnamed_migration

# Migration directory
MIGRATION_DIR := internal/infrastructure/database/migrations

# Command to create a new migration
create-migration:
	@if [ -z "$(name)" ]; then \
		echo "Usage: make create-migration name=your_migration_name"; \
		echo "Using default name: $(DEFAULT_NAME)"; \
		migrate create -ext sql -dir $(MIGRATION_DIR) -format "20060102150405_$(DEFAULT_NAME)" $(DEFAULT_NAME); \
	else \
		migrate create -ext sql -dir $(MIGRATION_DIR) -format "20060102150405_$(name)" $(name); \
	fi

# Command to generate Swagger documentation
swagger:
	swag init -d internal/handlers -g ../../cmd/api/main.go

# Help command
help:
	@echo "Available commands:"
	@echo "  make create-migration name=your_migration_name  - Create a new migration with the specified name"
	@echo "  make swagger                           		 - Generate Swagger documentation"
	@echo "  make help                                       - Show this help message"