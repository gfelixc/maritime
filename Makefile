MOCKERY_VERSION_REQUIRED=v2.20.0
GOLANGCI_LINT_VERSION_REQUIRED=1.52.2
BUF_VERSION_REQUIRED=1.17.0
GRPCURL_VERSION_REQUIRED=latest

init: check-buf check-golangci-lint check-mockery check-grpcurl
PHONY: init

check-grpcurl:
	@which grpcurl > /dev/null 2>&1 || make install-grpcurl
PHONY: install-grpcurl

install-grpcurl:
	@echo "Installing grpcurl $(GRPCURL_VERSION_REQUIRED)"
	@GOBIN=$$(go env GOPATH)/bin go install github.com/fullstorydev/grpcurl/cmd/grpcurl@v$(GRPCURL_VERSION_REQUIRED)
PHONY: install-grpcurl

check-buf:
	@which buf > /dev/null 2>&1 || make install-buf
	@BUF_VERSION_INSTALLED=$$(buf --version); \
	if [[ $$BUF_VERSION_INSTALLED != $(BUF_VERSION_REQUIRED) ]]; then \
  		make install-buf;\
	fi
PHONY: check-buf

install-buf:
	@echo "Installing buf $(BUF_VERSION_REQUIRED)"
	@GOBIN=$$(go env GOPATH)/bin go install github.com/bufbuild/buf/cmd/buf@v$(BUF_VERSION_REQUIRED)
PHONY: install-buf

check-golangci-lint:
	@which golangci-lint > /dev/null 2>&1 || make install-golangci-lint
	@GOLANGCI_LINT_VERSION_INSTALLED=$$(golangci-lint --version | cut -d ' ' -f4); \
	if [[ $$GOLANGCI_LINT_VERSION_INSTALLED != $(GOLANGCI_LINT_VERSION_REQUIRED) ]]; then \
	  make install-golangci-lint;\
	fi
PHONY: check-golangci-lint

install-golangci-lint:
	@echo "Installing golangci-lint $(GOLANGCI_LINT_VERSION_REQUIRED)"
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v$(GOLANGCI_LINT_VERSION_REQUIRED);\
PHONY: install-golangci-lint

check-mockery:
	@which mockery > /dev/null 2>&1 || make install-mockery
	@MOCKERY_VERSION_INSTALLED=$$(mockery --version --quiet); \
	if [[ $$MOCKERY_VERSION_INSTALLED != $(MOCKERY_VERSION_REQUIRED) ]]; then \
	  make install-mockery;\
	fi
PHONY: check-mockery

install-mockery:
	@echo "Installing mockery $(MOCKERY_VERSION_REQUIRED)"
	@go install github.com/vektra/mockery/v2@$(MOCKERY_VERSION_REQUIRED)
PHONY: install-mockery

run:
	go run cmd/port-domain/main.go
PHONY: run

lints:
	golangci-lint run --new-from-rev=HEAD~1
PHONY: lints

tests:
	go test -count=1 ./...
PHONY: tests

generate-mocks:
	go generate ./...
PHONY: generate-mocks

generate-protos:
	buf generate port
PHONY: generate-protos
