# Testkit - Go Testing Utility Library

Testkit is a Go module providing a modular builder framework to streamline test environment setup. It offers interfaces, ready-to-use structs, and patterns for creating reusable builders, fixtures, and mocks.

## Working Effectively

### Bootstrap and Dependencies
- Ensure the Go version declared in `go.mod` is installed
- Work from the repository root
- `go mod tidy` -- downloads and organizes dependencies (takes ~2 seconds)
- `go mod download` -- ensures all dependencies are available

### Build Process
- `go build ./...` -- builds all packages
- No external build dependencies required - pure Go project
- Build artifacts are not saved to disk for library projects

### Testing
- `go test ./... -v` -- runs all tests with verbose output
- `go test ./... -cover -coverprofile=coverage.out` -- runs tests with coverage (~95.9% coverage). Takes ~0.16 seconds.
- `go tool cover -html=coverage.out -o coverage.html` -- generates HTML coverage report
- All tests must pass before making changes

### Code Quality
- `go fmt ./...` -- formats all Go code (required before committing)
- `go vet ./...` -- runs static analysis (required before committing)  
- Linting: `golangci-lint run ./...` -- runs linters with local defaults; CI runs them with the shared config from `rios0rios0/pipelines`, which also enables `paralleltest`

### Run the Example Application
- `go run cmd/example/main.go` -- demonstrates all library features
- Expected output: Creates users via different patterns (basic, factory, config, validation, state management)
- Run it after changes to validate functionality; it should finish with `=== Demo Complete ===`

## Repository Structure

**Key directories and files:**
- `/` -- root module directory
- `pkg/test/` -- main library package (`github.com/rios0rios0/testkit/pkg/test`)
- `cmd/example/main.go` -- example application demonstrating all features
- `pkg/test/*.go` -- main library files (builder.go, factory.go, examples.go, doc.go)
- `pkg/test/*_test.go` -- comprehensive test files with 95.9% coverage
- `go.mod` -- module definition and required Go version
- `Makefile` -- build targets provided via shared pipeline makefiles
- `README.md` -- project documentation and usage examples
- `CHANGELOG.md` -- version history and release notes
- `CONTRIBUTING.md` -- contributing guidelines and development process
- `CODE_OF_CONDUCT.md` -- community code of conduct
- `LICENSE` -- MIT license

**Core components (all in `pkg/test/`):**
- `builder.go` -- BaseBuilder and Builder interface
- `examples.go` -- UserBuilder example implementation  
- `factory.go` -- BuilderFactory and BuilderConfig
- `doc.go` -- package documentation

## Contributing and Documentation

### Project Documentation
- `README.md` -- comprehensive project overview, installation, and usage examples
- `CHANGELOG.md` -- track version history and document breaking changes
- `CONTRIBUTING.md` -- development guidelines, code style, and contribution process
- **ALWAYS** write a changelog fragment with `chlog new --kind <Kind> --body '...'` and commit it from `.changes/unreleased/` -- CHANGELOG.md is generated from the fragments and is never edited by hand
- **ALWAYS** follow contributing guidelines for code style and testing requirements

### Making Changes
- Read CONTRIBUTING.md before making modifications
- Follow semantic versioning guidelines documented in CHANGELOG.md
- Update documentation when adding new features or changing APIs
- Ensure all examples in README.md continue to work after changes

## Common Tasks

### Testing Changes
- ALWAYS run: `go test ./... -v && go run cmd/example/main.go`
- For validation changes: Create a UserBuilder and test both valid and invalid inputs
- For new builder types: Test Build(), Clone(), Reset(), and error handling

### Adding New Features
- Follow the builder pattern established in `pkg/test/examples.go`
- Embed BaseBuilder for common functionality
- Implement validation in With* methods when IsValidationEnabled() is true
- Return errors from Build() when HasErrors() is true
- ALWAYS test Clone() and Reset() functionality

### Code Style
- Follow existing patterns: a doc comment on every exported symbol, and a justification on every `//nolint`
- Use method chaining pattern: return builder instance from With* methods
- Validate inputs only when IsValidationEnabled() returns true
- Return copies from Build() to prevent mutation
- Every test calls `t.Parallel()` -- CI's `paralleltest` linter requires it

## Important Notes

- **CI/CD pipeline** - `.github/workflows/default.yaml` uses the shared `rios0rios0/pipelines` Go library workflow; runs on push/PR to `main` and tags
- **Makefile** - provides targets via `rios0rios0/pipelines` shared makefiles (optional, not required for local development)
- **Library project** - no deployable artifacts, focus on API correctness
- **High test coverage** - maintain 95%+ coverage when adding features
- **Thread safety** - individual builders are NOT thread-safe by design, but `BuilderFactory` (and the `DefaultFactory` singleton) IS safe for concurrent use — its registry is guarded by a `sync.RWMutex`, so parallel tests can register and create through it
- **Error handling** - builders accumulate errors and report at build time
- **Extensibility** - designed for creating custom builder types

## Troubleshooting

**Build failures:**
- Run `go mod tidy` to fix dependency issues
- Check the installed Go matches the version in `go.mod`

**Test failures:**  
- Ensure no validation logic changes without updating tests
- Check that Builder interface implementations are complete

**Example app issues:**
- Verify UserBuilder methods return expected data types
- Check that validation errors contain expected messages

<!-- chlog:start -->
## Changelog (chlog) — MANDATORY

If the repository you are working in uses chlog (a `.chlog.yaml` or `.chlog.yml`
config file, or a `.changes/` directory, exists at the project root), the
following is binding and ALWAYS applies: whenever you make ANY change, you MUST
create a changelog fragment as part of the same change — automatically, without
being asked, before committing.

- Do NOT edit CHANGELOG.md directly; it is generated from fragments.
- Create the fragment with:
  `chlog new --kind <Kind> --body '<past-tense description>'`
- Write an apostrophe inside the single-quoted body as `'\''`.
- Valid kinds: Added, Changed, Deprecated, Removed, Fixed, Security
- Choose the kind that best matches the change (e.g., new feature → Added,
  bug fix → Fixed, behavior change → Changed, removal → Removed, security fix → Security).
- If the change is backward-INCOMPATIBLE with the public API (a breaking
  change), you MUST add the `--breaking` flag:
  `chlog new --kind <Kind> --breaking --body '<past-tense description>'`.
  This is the ONLY thing that triggers a major version bump — the kind alone
  never does (per SemVer, major = incompatible change). When unsure whether a
  change breaks compatibility, ask the user instead of guessing.
- Fragments are YAML files in `.changes/unreleased/`; stage them with your commit.
- `chlog check` fails the build when a fragment is missing — never skip it.
<!-- chlog:end -->
