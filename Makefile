LDFLAGS = -s -w
GOFLAGS = -trimpath -ldflags "$(LDFLAGS)"

define build-binary
	mkdir -p bin
	CGO_ENABLED=0 GOOS=$(1) GOARCH=$(2) go build -o bin/iapc-$(1)-$(2) $(GOFLAGS)
endef

all: build

.PHONY: build
build:
	go build ./...

.PHONY: test
test:
	go test -v ./...

.PHONY: release
release: test
	$(call build-binary,linux,amd64)
	$(call build-binary,linux,arm64)
	$(call build-binary,darwin,amd64)
	$(call build-binary,darwin,arm64)

.PHONY: clean
clean:
	rm -rf bin
