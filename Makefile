.PHONY: build run dev clean install lint check fmt vet test help rename version

BUILD_DIR    := build/bin
APP_NAME     := AnagataSentinel
FRONTEND_DIR := frontend
VERSION      := $(shell git describe --tags --always --dirty 2>nul || echo dev)
COMMIT       := $(shell git rev-parse --short HEAD 2>nul || echo unknown)
BUILD_TIME   := $(shell powershell -Command "Get-Date -Format 'yyyy-MM-ddTHH:mm:ssZ'")
GO_VERSION   := $(shell go version 2>nul)

LDFLAGS := -X 'AnagataSentinel/internal/version.Version=$(VERSION)' \
           -X 'AnagataSentinel/internal/version.GitCommit=$(COMMIT)' \
           -X 'AnagataSentinel/internal/version.BuildTime=$(BUILD_TIME)' \
           -X 'AnagataSentinel/internal/version.GoVersion=$(GO_VERSION)'

.DEFAULT_GOAL := help

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

build: ## Production build
	wails build -ldflags "$(LDFLAGS)" -clean

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
	wails build -ldflags "$(LDFLAGS)" -clean

test: ## Run tests
	go test ./...

tidy: ## Tidy go modules
	go mod tidy

version: ## Show current version
	@echo "Version:    $(VERSION)"
	@echo "Commit:     $(COMMIT)"
	@echo "Build Time: $(BUILD_TIME)"
	@echo "Go Version: $(GO_VERSION)"

rename: ## Rename project: make rename NEW=MyAppName
	@if [ -z "$(NEW)" ]; then echo "Usage: make rename NEW=MyAppName"; exit 1; fi
	@echo "Renaming $(APP_NAME) -> $(NEW)..."
	@# go.mod
	powershell -Command "(Get-Content go.mod -Raw) -replace '$(APP_NAME)','$(NEW)' | Set-Content go.mod -NoNewline"
	@# wails.json
	powershell -Command "$$c = Get-Content wails.json -Raw; $$c = $$c.replace('\"$(APP_NAME)\"', '\"$(NEW)\"'); Set-Content wails.json $$c -NoNewline"
	@# main.go (Title + import)
	powershell -Command "(Get-Content main.go -Raw) -replace '$(APP_NAME)','$(NEW)' | Set-Content main.go -NoNewline"
	@# internal/app/app.go (imports)
	powershell -Command "(Get-Content internal/app/app.go -Raw) -replace '$(APP_NAME)','$(NEW)' | Set-Content internal/app/app.go -NoNewline"
	@# internal/database/database.go (imports)
	powershell -Command "(Get-Content internal/database/database.go -Raw) -replace '$(APP_NAME)','$(NEW)' | Set-Content internal/database/database.go -NoNewline"
	@# frontend wailsjs import
	powershell -Command "$$old = 'go/$($$((Get-Content go.mod -First 1)).Split('/')[-1].ToLower())/App'; $$new = 'go/$(NEW.ToLower())/App'; (Get-Content $(FRONTEND_DIR)/src/App.svelte -Raw) -replace $$old,$$new | Set-Content $(FRONTEND_DIR)/src/App.svelte -NoNewline"
	@# internal/splash/splash_windows.go (window title)
	powershell -Command "(Get-Content internal/splash/splash_windows.go -Raw) -replace '$(APP_NAME)','$(NEW)' | Set-Content internal/splash/splash_windows.go -NoNewline"
	@go mod tidy
	@echo "Done! Renamed to $(NEW)"
