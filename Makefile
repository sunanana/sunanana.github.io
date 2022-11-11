.PHONY: help
help: ## show this help message.
	@grep -hE '^\S+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## build wasm file
	GOOS=js GOARCH=wasm go build -o ast.wasm main.go