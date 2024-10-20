all: build

.PHONY: build
build:
	go build ./...

.PHONY: test
test:
	go test -v ./...

.PHONY: release
release: test
	goreleaser release

.PHONY: clean
clean:
	rm -rf dist
