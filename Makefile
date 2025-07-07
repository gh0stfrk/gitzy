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
	GOOS=windows GOARCH=amd64 go build -o dist/gitzy-windows-amd64.exe ./cmd
	GOOS=linux GOARCH=amd64 go build -o dist/gitzy-linux-amd64 ./cmd
	GOOS=darwin GOARCH=amd64 go build -o dist/gitzy-darwin-amd64 ./cmd
	# Copy install scripts
	cp scripts/install_windows.bat dist/
	cp scripts/install_linux.sh dist/
	# Package for Windows (binary + install script)
	cd dist && zip gitzy-windows-amd64.zip gitzy-windows-amd64.exe install_windows.bat && sha256sum gitzy-windows-amd64.zip > gitzy-windows-amd64.zip.sha256
	# Package for Linux (binary + install script)
	cd dist && zip gitzy-linux-amd64.zip gitzy-linux-amd64 install_linux.sh && sha256sum gitzy-linux-amd64.zip > gitzy-linux-amd64.zip.sha256
	# Package for macOS (binary only)
	cd dist && zip gitzy-darwin-amd64.zip gitzy-darwin-amd64 && sha256sum gitzy-darwin-amd64.zip > gitzy-darwin-amd64.zip.sha256
