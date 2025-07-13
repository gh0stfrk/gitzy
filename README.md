<p align="center">
<a href="https://github.com/gh0stfrk/gitzy">
    <img src="./docs/icon.svg">
</a>
</p>
<h1 align=center>Gitzy</h1>
<h2 align="center">Quickly zap into a new profile ⚡</h2>

<p align="center">
<a href="https://github.com/gh0stfrk/gitzy/releases"><img src="https://img.shields.io/github/v/release/gh0stfrk/gitzy"></a>
<a href="https://github.com/gh0stfrk/gitzy/blob/main/LICENSE"><img src="https://img.shields.io/github/license/gh0stfrk/gitzy"></a>
<a href="https://github.com/gh0stfrk/gitzy/actions"><img src="https://img.shields.io/github/actions/workflow/status/gh0stfrk/gitzy/ci.yml"></a>
<a href="https://goreportcard.com/report/github.com/gh0stfrk/gitzy"><img src="https://goreportcard.com/badge/github.com/gh0stfrk/gitzy"></a>
</p>

## About

Gitzy is a cross-platform CLI tool that helps developers manage multiple GitHub identities and switch between them seamlessly. Whether you're working on personal projects, client work, or contributing to different organizations, Gitzy makes it easy to maintain separate Git configurations and credentials.

### Features

- 🔄 **Quick Profile Switching**: Switch between different Git identities instantly
- 👤 **Multiple Identity Management**: Store and manage unlimited GitHub profiles
- 🔐 **Secure Credential Storage**: Safely store GitHub tokens and credentials
- 🖥️ **Cross-Platform**: Works on Windows, macOS, and Linux
- ⚡ **Fast & Lightweight**: Built with Go for optimal performance
- 🛠️ **Easy Setup**: Simple commands to add, switch, and manage profiles

## Installation

### Option 1: Install from GitHub Releases (Recommended)

1. Go to the [Releases page](https://github.com/gh0stfrk/gitzy/releases)
2. Download the appropriate binary for your platform:
   - **Windows**: `gitzy-windows-amd64.zip`
   - **Linux**: `gitzy-linux-amd64.zip`
   - **macOS**: `gitzy-darwin-amd64.zip`
3. Extract the archive and run the included install script:
   - **Windows**: Run `install_windows.ps1` in PowerShell
   - **Linux/macOS**: Run `chmod +x install_linux.sh && ./install_linux.sh`

### Option 2: Install with Go

```bash
go install github.com/gh0stfrk/gitzy@latest
```

### Option 3: Build from Source

```bash
git clone https://github.com/gh0stfrk/gitzy.git
cd gitzy
make build
```

## Quick Start

1. **Add a new profile**:
   ```bash
   gitzy add work
   ```

2. **Switch to a profile**:
   ```bash
   gitzy switch work
   ```

3. **List all profiles**:
   ```bash
   gitzy list
   ```

4. **Remove a profile**:
   ```bash
   gitzy wipe work
   ```

## Usage

### Adding a Profile

```bash
gitzy add <profile-name>
```

This will prompt you to enter:
- GitHub username
- Email address
- Personal access token (optional)

### Switching Profiles

```bash
gitzy switch <profile-name>
```

This updates your global Git configuration to use the selected profile's credentials.

### Listing Profiles

```bash
gitzy list
```

Shows all configured profiles with their details.

### Removing Profiles

```bash
gitzy wipe <profile-name>
```

Permanently removes a profile and its associated data.

### Troubleshooting

```bash
gitzy doctor --docs
```

Provides diagnostic information and troubleshooting guidance.

## Development

### Prerequisites

- Go 1.22 or later
- Make (optional, for using Makefile commands)

### Setting Up Development Environment

1. **Clone the repository**:
   ```bash
   git clone https://github.com/gh0stfrk/gitzy.git
   cd gitzy
   ```

2. **Install dependencies**:
   ```bash
   go mod download
   ```

3. **Install development tools**:
   ```bash
   go install honnef.co/go/tools/cmd/staticcheck@latest
   ```

4. **Run tests**:
   ```bash
   make test
   # or
   go test ./...
   ```

5. **Build the project**:
   ```bash
   make build
   # or
   go build -o bin/gitzy .
   ```

### Development Commands

- `make build` - Build the binary
- `make test` - Run all tests
- `make vet` - Run go vet
- `make staticcheck` - Run static analysis
- `make release` - Build release binaries for all platforms

### Project Structure

```
├── cmd/           # CLI commands and subcommands
├── internal/      # Internal packages
│   ├── git/       # Git profile management
│   └── ui/        # User interface utilities
├── docs/          # Documentation
├── scripts/       # Installation scripts
├── tools/         # Development tools
└── examples/      # Usage examples
```

## Contributing

We welcome contributions! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

### Quick Contribution Guide

1. Fork the repository
2. Create a feature branch: `git checkout -b feature-name`
3. Make your changes and add tests
4. Run the test suite: `make test`
5. Run linting: `make vet && make staticcheck`
6. Commit your changes: `git commit -am 'Add feature'`
7. Push to the branch: `git push origin feature-name`
8. Submit a pull request

## CLI Reference

Detailed CLI documentation is available in [docs/cli/](docs/cli/) or visit our [documentation site](https://gh0stfrk.github.io/gitzy/).

## Configuration

### Storage Locations

- **Windows**: `%USERPROFILE%\.gitzy`
- **Linux/macOS**: `~/.gitzy` (respects `XDG_CONFIG_HOME` if set)
- **Custom**: Set `GITZY_HOME` environment variable

### Security Notes

- Tokens are stored unencrypted in the configuration directory
- Follow [GitHub token best practices](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token)
- Consider using tokens with minimal required scopes

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

- 📖 [Documentation](https://gh0stfrk.github.io/gitzy/)
- 🐛 [Report Issues](https://github.com/gh0stfrk/gitzy/issues)
- 💬 [Discussions](https://github.com/gh0stfrk/gitzy/discussions)
- ⭐ Star this repo if you find it helpful!

## Acknowledgments

- Built with [Cobra](https://github.com/spf13/cobra) for CLI framework
- Uses [Survey](https://github.com/AlecAivazis/survey) for interactive prompts
- Inspired by the need for seamless Git identity management
