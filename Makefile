TARGETS = monkey-lang

deps:
	go get ./...

clean:
	rm -rf $(TARGETS)

build: deps
	go build ./...

test:
	go test ./...
