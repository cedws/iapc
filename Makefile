LDFLAGS = -s -w
GOFLAGS = -trimpath -ldflags "$(LDFLAGS)"

define build
	mkdir -p bin
	CGO_ENABLED=0 GOOS=$(1) GOARCH=$(2) go build -o bin/iapc-$(1)-$(2) $(GOFLAGS)
endef

all: release

.PHONY: release
release:
	$(call build,linux,amd64)
	$(call build,linux,arm64)
	$(call build,darwin,amd64)
	$(call build,darwin,arm64)

.PHONY: clean
clean:
	rm -rf bin
