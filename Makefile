BIN     := bin/joist
PKG     := ./cmd/joist
VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
LDFLAGS := -X main.version=$(VERSION)

.DEFAULT_GOAL := build
.PHONY: build dev web test lint fmt fmt-check tidy clean

# Release build. Runs the web build first because the spa tag requires dist.
build: web
	go build -tags spa -ldflags '$(LDFLAGS)' -o $(BIN) $(PKG)

# Development build. No Node required; the viewer serves a placeholder page.
dev:
	go build -ldflags '$(LDFLAGS)' -o $(BIN) $(PKG)

web:
	cd web && npm ci && npm run build

test:
	go test ./...
	cd schema && go test ./...

lint:
	go vet ./...
	cd schema && go vet ./...

fmt:
	gofmt -w cmd internal schema

# Fails when any file needs formatting. Used by CI.
fmt-check:
	@out=$$(gofmt -l cmd internal schema); \
	if [ -n "$$out" ]; then echo "gofmt needed:"; echo "$$out"; exit 1; fi

tidy:
	go mod tidy
	cd schema && go mod tidy

clean:
	rm -rf bin internal/viewer/dist web/node_modules
