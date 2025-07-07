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
	# Build binaries for all platforms
	@for plat in $(PLATFORMS); do \
		GOOS=$${plat%/*} GOARCH=$${plat#*/} go build -o dist/gitzy-$${plat%/*}-$${plat#*/} ./cmd; \
	done
	# Copy install scripts
	cp scripts/install_windows.bat dist/
	cp scripts/install_linux.sh dist/
	# Package for Windows
	cd dist && zip gitzy-windows-amd64.zip gitzy-windows-amd64.exe install_windows.bat && sha256sum gitzy-windows-amd64.zip > gitzy-windows-amd64.zip.sha256
	# Package for Linux
	cd dist && tar czf gitzy-linux-amd64.tar.gz gitzy-linux-amd64 install_linux.sh && sha256sum gitzy-linux-amd64.tar.gz > gitzy-linux-amd64.tar.gz.sha256
	# Package for macOS
	cd dist && tar czf gitzy-darwin-amd64.tar.gz gitzy-darwin-amd64 && sha256sum gitzy-darwin-amd64.tar.gz > gitzy-darwin-amd64.tar.gz.sha256
