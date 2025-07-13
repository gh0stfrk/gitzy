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
	rm -rf dist && mkdir -p dist/windows dist/linux dist/darwin

	# Build binaries into platform-specific subdirs with consistent names
	GOOS=windows GOARCH=amd64 go build -o dist/windows/gitzy.exe .
	GOOS=linux   GOARCH=amd64 go build -o dist/linux/gitzy .
	GOOS=darwin  GOARCH=amd64 go build -o dist/darwin/gitzy .

	# Copy install scripts into respective folders
	cp scripts/install_windows.ps1 dist/windows/
	cp scripts/install_linux.sh dist/linux/

	# Zip and checksum each platform's build
	cd dist/windows && zip gitzy-windows-amd64.zip gitzy.exe install_windows.ps1 && sha256sum gitzy-windows-amd64.zip > gitzy-windows-amd64.zip.sha256
	cd dist/linux   && zip gitzy-linux-amd64.zip gitzy install_linux.sh       && sha256sum gitzy-linux-amd64.zip   > gitzy-linux-amd64.zip.sha256
	cd dist/darwin  && zip gitzy-darwin-amd64.zip gitzy                       && sha256sum gitzy-darwin-amd64.zip  > gitzy-darwin-amd64.zip.sha256
