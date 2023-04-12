init:
	@echo "Linking git hooks..."
	@chmod +x $(CURDIR)/git-hooks/*
	@ln -s $(CURDIR)/git-hooks/* $(CURDIR)/.git/hooks
	@echo "Running go mod download..."
	@go mod download
PHONY: init

run:
	go run cmd/port-domain/main.go