BACKEND_DIR=backend
FRONTEND_DIR=src
ENV_FILE=.env

.PHONY: help
help:
	@echo "Makefile for MovieCollection project"
	@echo ""
	@echo "Available targets:"
	@echo "  backend-setup      - Set up the backend environment"
	@echo "  backend-run        - Run the backend server"
	@echo "  frontend-setup     - Set up the frontend environment"
	@echo "  frontend-start     - Start the frontend development server"
	@echo "  clean              - Clean up build artifacts"

.PHONY: backend-setup
backend-setup:
	@echo "Setting up the backend environment..."
	cd $(BACKEND_DIR) && go get -u ./... && go build

.PHONY: backend-run
backend-run:
	@echo "Running the backend server..."
	cd $(BACKEND_DIR) && go run main.go

.PHONY: frontend-setup
frontend-setup:
	@echo "Setting up the frontend environment..."
	cd $(FRONTEND_DIR) && npm install

.PHONY: frontend-start
frontend-start:
	@echo "Starting the frontend development server..."
	cd $(FRONTEND_DIR) && npm start

.PHONY: cle
