export PATH := $(CURDIR)/.bin:$(PATH)

TARGETS = monkey-lang
TEST =  ./...

default: setup

setup:
	GOBIN=$(CURDIR)/.bin $(MAKE) -C tools

clean:
	rm -rf $(TARGETS)

build:
	go build ./...

test: setup
	richgo test -v $(TEST)

run: build
	go run ./...
