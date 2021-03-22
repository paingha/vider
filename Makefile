PROJECT_NAME := "vider"
PKG := "github.com/paingha/$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

lint: ## Lint Golang files
	@golint -set_exit_status ${PKG_LIST}
	
test: ## Run unit tests
	@go test -short ${PKG_LIST}

test-coverage: ## Run tests with coverage
	@go test
