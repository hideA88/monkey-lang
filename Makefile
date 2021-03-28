export PATH := $(CURDIR)/.bin:$(PATH)

TARGETS = monkey-lang
TEST =  ./...

default: setup

setup:
	GOBIN=$(CURDIR)/.bin $(MAKE) -C tools

clean:
	rm -rf $(TARGETS)

build: setup
	go build ./...

test: setup
	richgo test -v $(TEST)
