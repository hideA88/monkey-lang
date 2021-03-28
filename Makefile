TARGETS = monkey-lang

setup:
	go get github.com/Songmu/make2help/cmd/make2help
	go get github.com/kyoh86/richgo
deps:
	go get ./...

clean:
	rm -rf $(TARGETS)

build: deps
	go build ./...

test:
	richgo test -v ./...
