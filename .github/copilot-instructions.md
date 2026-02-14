# Testkit - Go Testing Utility Library

Testkit is a Go module providing a modular builder framework to streamline test environment setup. It offers interfaces, ready-to-use structs, and patterns for creating reusable builders, fixtures, and mocks.

**ALWAYS follow these instructions first and fallback to search or bash commands only when you encounter unexpected information that does not match the info here.**

## Working Effectively

### Bootstrap and Dependencies
- Ensure Go 1.24.7+ is installed
- `cd /home/runner/work/testkit/testkit` (repository root)
- `go mod tidy` -- downloads and organizes dependencies (takes ~2 seconds)
- `go mod download` -- ensures all dependencies are available

### Build Process
- `go build ./...` -- builds all packages. Takes ~0.13 seconds. NEVER CANCEL. Set timeout to 30+ seconds.
- No external build dependencies required - pure Go project
- Build artifacts are not saved to disk for library projects

### Testing
- `go test ./... -v` -- runs all tests with verbose output. Takes ~0.07 seconds. NEVER CANCEL. Set timeout to 30+ seconds.
- `go test ./... -cover -coverprofile=coverage.out` -- runs tests with coverage (~95.9% coverage). Takes ~0.16 seconds.
- `go tool cover -html=coverage.out -o coverage.html` -- generates HTML coverage report
- All tests must pass before making changes

### Code Quality
- `go fmt ./...` -- formats all Go code (required before committing)
- `go vet ./...` -- runs static analysis (required before committing)  
- Linting: golangci-lint has compatibility issues with Go 1.24.7 - do NOT attempt to install or run it

### Run the Example Application
- `go run cmd/example/main.go` -- demonstrates all library features. Takes ~0.04 seconds.
- Expected output: Creates users via different patterns (basic, factory, config, validation, state management)
- ALWAYS run this after making changes to validate functionality

## Validation Scenarios

**ALWAYS run these complete validation scenarios after making any changes:**

1. **Basic Builder Functionality Test**:
   ```bash
   go run cmd/example/main.go
   ```
   - Verify all 6 examples run successfully
   - Check that user objects are created with expected data
   - Confirm validation errors are properly handled
   - Look for "Demo Complete" message at the end

2. **Library Integration Test**:
   ```bash
   go test ./... -v
   ```
   - All 26 tests must pass
   - No test failures or panics allowed
   - Coverage should remain at ~95%+

3. **Code Quality Validation**:
   ```bash
   go fmt ./...
   go vet ./...
   ```
   - go fmt should produce minimal or no output after first run
   - go vet must report no issues

## Repository Structure

**Key directories and files:**
- `/` -- root module directory with all source code
- `cmd/example/main.go` -- example application demonstrating all features
- `*.go` -- main library files (builder.go, factory.go, examples.go, doc.go)
- `*_test.go` -- comprehensive test files with 95.9% coverage
- `go.mod` -- module definition (requires Go 1.24.7+)
- `README.md` -- project documentation and usage examples
- `CHANGELOG.md` -- version history and release notes
- `CONTRIBUTING.md` -- contributing guidelines and development process
- `LICENSE` -- MIT license

**Core components:**
- `builder.go` -- BaseBuilder and Builder interface
- `examples.go` -- UserBuilder example implementation  
- `factory.go` -- BuilderFactory and BuilderConfig
- `doc.go` -- package documentation

## Contributing and Documentation

### Project Documentation
- `README.md` -- comprehensive project overview, installation, and usage examples
- `CHANGELOG.md` -- track version history and document breaking changes
- `CONTRIBUTING.md` -- development guidelines, code style, and contribution process
- **ALWAYS** update CHANGELOG.md when making user-facing changes
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
- Follow the builder pattern established in examples.go
- Embed BaseBuilder for common functionality
- Implement validation in With* methods when IsValidationEnabled() is true
- Return errors from Build() when HasErrors() is true
- ALWAYS test Clone() and Reset() functionality

### Code Style
- Follow existing patterns: no comments except for package documentation
- Use method chaining pattern: return builder instance from With* methods
- Validate inputs only when IsValidationEnabled() returns true
- Return copies from Build() to prevent mutation

## Build Times and Timeouts

**NEVER CANCEL any of these commands - use these minimum timeouts:**
- `go build ./...` -- ~0.13 seconds, set timeout to 30+ seconds
- `go test ./...` -- ~0.07 seconds, set timeout to 30+ seconds  
- `go test ./... -cover` -- ~0.16 seconds, set timeout to 30+ seconds
- `go run cmd/example/main.go` -- ~0.04 seconds, set timeout to 10+ seconds
- `go mod tidy` -- ~2 seconds, set timeout to 30+ seconds

## Frequently Used Commands Output

### Repository Root Contents
```
.git/          -- git repository
.github/       -- GitHub configuration
.gitignore     -- Go-specific gitignore
LICENSE        -- MIT license
README.md      -- project documentation and usage examples
CHANGELOG.md   -- version history and release notes
CONTRIBUTING.md -- contributing guidelines and development process
go.mod         -- Go module file
*.go           -- source files (builder.go, examples.go, factory.go, doc.go)
*_test.go      -- test files
cmd/           -- example application
```

### go.mod Contents
```go
module github.com/rios0rios0/testkit

go 1.24.7
```

### Example Application Output
```
=== Testkit Library Demo ===
1. Basic UserBuilder Usage:
   Created user: &{ID:0 Name:Alice Smith Email:alice@example.com Age:28 Active:true Tags:map[department:engineering] Metadata:map[hire_date:2023-01-15]}

2. Factory Pattern Usage:
   Factory-created user: &{ID:0 Name:Bob Wilson Email:bob@example.com Age:35 Active:false Tags:map[] Metadata:map[]}

3. Configuration System:
   Configured user: &{ID:0 Name:Test User Email:test@example.com Age:30 Active:true Tags:map[] Metadata:map[]}
   Builder tags: env=test, team=qa

4. Validation and Error Handling:
   Validation failed as expected: cannot build user due to validation errors: [user name cannot be empty user email cannot be empty user age must be non-negative]

5. Builder State Management:
   Original builder name: v1, tag: Original User
   Cloned builder name: v2, tag: Cloned User
   After reset, original builder has errors: false

6. Custom Factory Usage:
   Admin user: &{ID:0 Name:Admin User Email:admin@example.com Age:0 Active:true Tags:map[role:admin] Metadata:map[]}

=== Demo Complete ===
```

## Important Notes

- **No CI/CD pipeline exists** - manual validation required
- **No Makefile** - use go commands directly  
- **Library project** - no deployable artifacts, focus on API correctness
- **High test coverage** - maintain 95%+ coverage when adding features
- **Thread safety** - builders are NOT thread-safe by design
- **Error handling** - builders accumulate errors and report at build time
- **Extensibility** - designed for creating custom builder types

## Troubleshooting

**Build failures:**
- Run `go mod tidy` to fix dependency issues
- Check Go version is 1.24.7+

**Test failures:**  
- Ensure no validation logic changes without updating tests
- Check that Builder interface implementations are complete

**Example app issues:**
- Verify UserBuilder methods return expected data types
- Check that validation errors contain expected messages