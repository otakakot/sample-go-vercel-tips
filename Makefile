SHELL := /bin/bash
include .env
export

.PHONY: help
help: ## display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: deploy
deploy: ## deploy to vercel
	@vercel --prod

.PHONY: destroy
destroy: ## destroy the vercel deployment
	@vercel project rm sample-go-vercel-tips

.PHONY: gen
gen: ## generate the types
	@oapi-codegen -generate types -package openapi api/openapi.yaml > pkg/openapi/types.gen.go
	@go mod tidy

.PHONY: test
test: ## run the tests
	@go test -v ./...
