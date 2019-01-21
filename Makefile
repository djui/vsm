SRC_PKG := ./cmd/... ./pkg/...

.PHONY: help
help:
	@grep -E '^[a-zA-Z0-9_/%\-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "%-20s %s\n", $$1, $$2}'

.PHONY: build
build: ## Build command-line binary
	GO111MODULE=on go build -race ./cmd/...

.PHONY: deps
deps: ## Fetch dependencies
	go get -u github.com/golang/dep/cmd/dep

	curl -L -o /tmp/gometalinter.sh https://raw.githubusercontent.com/alecthomas/gometalinter/master/scripts/install.sh
	sh /tmp/gometalinter.sh -b $$GOPATH/bin v2.0.12

.PHONY: install
install: ## Install command-line binary
	GO111MODULE=on go install -race ./cmd/...

.PHONY: lint
lint: ## Run linters
	gometalinter ./...

.PHONY: test
test: ## Run unit tests
	GO111MODULE=on go test -race ${SRC_PKG}
