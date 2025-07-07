# gitzy

**gitzy** is a cross-platform CLI for managing and switching between multiple GitHub identities.

## Install

```
go install github.com/gh0stfrk/gitzy@latest
```

Or download a pre-built binary from [Releases](https://github.com/gh0stfrk/gitzy/releases).

## Usage

```
gitzy add work
gitzy switch work
gitzy list
gitzy wipe work
```

See `gitzy doctor --docs` for troubleshooting.

## Examples

See [docs/](docs/) for usage gifs.

## Troubleshooting

- On Windows, uses `%USERPROFILE%\.gitzy` and the "store" credential helper.
- On Linux/macOS, uses `~/.gitzy` or respects `XDG_CONFIG_HOME`.
- Tokens are stored unencrypted. See [GitHub token best practices](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token).
