BINS=gitzy
PLATFORMS=linux/amd64 darwin/amd64 windows/amd64

.PHONY: all build test vet staticcheck release

all: build

build:
	go build -o bin/$(BINS) ./cmd

test:
	go test ./...

vet:
	go vet ./...

staticcheck:
	staticcheck ./...

release:
	rm -rf dist && mkdir dist
	@for plat in $(PLATFORMS); do \
		GOOS=$${plat%/*} GOARCH=$${plat#*/} go build -o dist/gitzy-$${plat%/*}-$${plat#*/} ./cmd; \
	done
	cd dist && for f in *; do \
		if [[ $$f == *.exe ]]; then \
			zip "$$f.zip" "$$f"; \
			sha256sum "$$f.zip" > "$$f.zip.sha256"; \
		else \
			tar czf "$$f.tar.gz" "$$f"; \
			sha256sum "$$f.tar.gz" > "$$f.tar.gz.sha256"; \
		fi \
	done
