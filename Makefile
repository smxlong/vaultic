# Copyright 2024 Scott Long
#
# SPDX-License-Identifier: MIT

.PHONY: check
check: fmt vet lint test

.PHONY: check-pr
check-pr: check
	echo ">> checking pull request"
	git fetch origin main
	! git diff --exit-code origin/main..HEAD version/version.go || (echo "Version number has not been updated"; exit 1)
	grep -q "## v$$(grep -oP '(?<=Version = ").*(?=")' version/version.go)" CHANGELOG.md || (echo "Changelog has not been updated"; exit 1)

.PHONY: fmt
fmt: generate
	@echo ">> formatting code"
	@go fmt ./...

.PHONY: vet
vet: generate
	@echo ">> vetting code"
	@go vet ./...

.PHONY: lint
lint: generate
	@echo ">> linting code"
	@golangci-lint run

.PHONY: test
test: generate
	@echo ">> running tests"
	@go test -coverprofile=coverage.txt -covermode=atomic -v ./...

.PHONY: generate
generate:
	@echo ">> generating code"
	@go generate ./...

.PHONY: build
build: check bin
	@echo ">> building binary"
	@go build -o bin/vaultic ./cmd/vaultic

bin:
	@mkdir -p bin

.PHONY: docker-build
docker-build:
	@echo ">> building docker image"
	@docker build -t smxlong/vaultic:latest .
