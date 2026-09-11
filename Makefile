.PHONY: build run dev clean install lint check fmt vet test help

BUILD_DIR    := build/bin
APP_NAME     := AnagataSentinel
FRONTEND_DIR := frontend

.DEFAULT_GOAL := help

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

build: ## Production build
	wails build

run: build ## Build and run
	./$(BUILD_DIR)/$(APP_NAME).exe

dev: ## Development mode with hot-reload
	wails dev

clean: ## Remove build artifacts and node_modules
	rm -rf $(BUILD_DIR) $(FRONTEND_DIR)/dist $(FRONTEND_DIR)/node_modules data/

install: ## Install frontend dependencies
	cd $(FRONTEND_DIR) && pnpm install

fmt: ## Format Go code
	gofmt -s -w .

vet: ## Run go vet
	go vet ./...

lint: fmt vet ## Format and vet

check: lint ## Full check (lint + build)
	cd $(FRONTEND_DIR) && npx svelte-check
	go build -o /dev/null .

test: ## Run tests
	go test ./...

tidy: ## Tidy go modules
	go mod tidy
