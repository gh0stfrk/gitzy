# Contributing to gitzy

Thank you for your interest!

- Please open issues for bugs or feature requests.
- All code must be formatted (`go fmt`), linted, and tested.
- Run `go vet`, `staticcheck`, and `go test ./...` before submitting a PR.
- Add or update unit tests for new features.
- For major changes, open a discussion first.

## Documentation Generation

- CLI documentation is auto-generated from the code using Cobra.
- Before every commit, the pre-commit hook will run `go run -tags tools ./tools/gendocs.go` to update the docs in `docs/cli/`.
- If you want to manually generate docs, run:
  ```sh
  go run -tags tools ./tools/gendocs.go
  ```
- Make sure to commit any changes in `docs/cli/` if you update command help, usage, or examples.

Happy hacking!
