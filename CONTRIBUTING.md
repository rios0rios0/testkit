# Contributing

Contributions are welcome. By participating, you agree to maintain a respectful and constructive environment.

For coding standards, testing patterns, architecture guidelines, commit conventions, and all
development practices, refer to the **[Development Guide](https://github.com/rios0rios0/guide/wiki)**.

## Prerequisites

- [Go](https://go.dev/dl/) 1.26+
- [Make](https://www.gnu.org/software/make/)
- [golangci-lint](https://golangci-lint.run/) (for linting)
- [Pipelines repo](https://github.com/rios0rios0/pipelines) cloned at `~/Development/github.com/rios0rios0/pipelines` (for shared Makefile targets)

## Development Workflow

1. Fork and clone the repository
2. Create a branch: `git checkout -b feat/my-change`
3. Set up the shared pipelines (one-time):
   ```bash
   make setup
   ```
   This clones the [pipelines](https://github.com/rios0rios0/pipelines) repository with shared Makefile targets.
4. Install dependencies:
   ```bash
   go mod download
   ```
5. Run the linter:
   ```bash
   make lint
   ```
6. Run tests:
   ```bash
   make test
   ```
7. Run static analysis (SAST):
   ```bash
   make sast
   ```
8. Use this library in your project:
   ```bash
   go get github.com/rios0rios0/testkit
   ```
9. Commit following the [commit conventions](https://github.com/rios0rios0/guide/wiki/Life-Cycle/Git-Flow)
10. Open a pull request against `main`
