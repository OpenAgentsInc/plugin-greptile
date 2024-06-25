# Load environment variables from .env file
ifneq (,$(wildcard .env))
    include .env
    export
endif

build:
	tinygo build -target wasi -o plugin.wasm main.go

test:
	@if [ -z "$(GREPTILE_API_KEY)" ] || [ -z "$(GITHUB_TOKEN)" ]; then \
		echo "Error: GREPTILE_API_KEY and GITHUB_TOKEN must be set in .env file"; \
		exit 1; \
	fi
	extism call plugin.wasm run --input '{"repository": "openagentsinc/plugin-greptile", "remote": "github", "branch": "main", "api_key": "$(GREPTILE_API_KEY)", "github_token": "$(GITHUB_TOKEN)"}' --wasi --allow-host="*.greptile.com"

.PHONY: build test
