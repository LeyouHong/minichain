PACKAGES = $(shell go list ./... | grep -v '/vendor/')

test:
	@echo "====> Running go test"
	@go test $(PACKAGES)

.PHONY: test
