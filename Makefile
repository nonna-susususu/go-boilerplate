ENV_FILES := $(wildcard .env)
ifneq ($(ENV_FILES),)
  $(info Sourcing environment variables from $(ENV_FILES))
  # Read lines from .env, filter comments/blank lines, then eval export for each.
  # This ensures 'make' itself processes the export command for each variable.
  $(foreach env_file,$(ENV_FILES),$(foreach line,$(shell egrep -v '^#|^$$' $(env_file)),$(eval export $(line))))
endif

.PHONY: tidy dev gen-mock unit-test

tidy:
	@go mod tidy

dev:
	@go run ./cmd/main.go

gen-mock:
	@mockery

unit-test:
	@ginkgo ./...
