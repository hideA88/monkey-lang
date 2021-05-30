export PATH := $(CURDIR)/.bin:$(PATH)

TARGETS = monkey-lang

GOLANGCI_LINT = golangci-lint run
TEST = ./...
PKGNAME = $(shell go list -m)
GIT_COMMIT = $(shell git rev-parse HEAD)
LDFLAGS = -ldflags "-X $(PKGNAME)/pkg/version.gitCommit=$(GIT_COMMIT)"

# command
defualt: help

## Install dependency tidy
tidy: deps.tidy tools.tidy

## Install dependency
deps:
	go get ./...
deps.update.minor:
	go get -t -u ./...
deps.update.patch:
	go get -t -u=patch ./...
deps.tidy:
	go mod tidy

## Install tools
tools: deps
	GOBIN=$(CURDIR)/.bin $(MAKE) -C tools
tools.update:
	GOBIN=$(CURDIR)/.bin $(MAKE) -C tools update
tools.tidy:
	GOBIN=$(CURDIR)/.bin $(MAKE) -C tools tidy

## Remove build target
clean:
	rm -f $(TARGETS)
	rm -rf tmp

## Build app
build: deps
	go build ./...
build.cmd: clean $(TARGETS)

$(TARGETS): deps
	go build $(LDFLAGS) ./cmd/$@

## Check code format
check: tools
	$(GOLANGCI_LINT) ./...

## Fix code
fix: tools
	$(GOLANGCI_LINT) --fix ./...

## Run test
test: deps tools
	mkdir -p tmp
	richgo test -race -coverprofile=tmp/coverage.txt -covermode=atomic $(TEST)

## Show help
help:
	@make2help $(MAKEFILE_LIST)

NO_PHONY = /^:/
PHONY := $(shell cat $(MAKEFILE_LIST) | awk -F':' '/^[a-z0-9_.-]+:/ && !$(NO_PHONY) {print $$1}')
.PHONY: $(PHONY)

show_phony:
	@echo $(PHONY)
