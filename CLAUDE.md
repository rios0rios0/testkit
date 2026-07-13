# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project

Go testing utility library (`github.com/rios0rios0/testkit`). Provides a modular builder framework for test environment setup. Pure library -- no deployable artifacts. Requires Go 1.26.5+.

## Commands

```bash
go build ./...              # build all packages
go test ./... -v            # run all tests
go fmt ./...                # format (required before commit)
go vet ./...                # static analysis (required before commit)
golangci-lint run ./...     # lint (matches CI)
go run cmd/example/main.go  # run example app -- use to validate changes
```

The `Makefile` pulls shared targets from `rios0rios0/pipelines` via `$SCRIPTS_DIR`. Use `go` commands directly for local development.

## Architecture

All library code is in `pkg/test/` (Go package name: `testkit`).

| File | Purpose |
|------|---------|
| `builder.go` | `BaseBuilder` struct and `Builder` interface |
| `factory.go` | `BuilderFactory`, `BuilderConfig`, global registry |
| `examples.go` | `UserBuilder` reference implementation, `TestUser` entity |
| `doc.go` | Package-level documentation |

Tests live in the same package (`package testkit`) for internal field access.

## Conventions

- Custom builders embed `*BaseBuilder` and implement `Builder` (`Build`, `Reset`, `Clone`).
- `With*` methods return the builder (method chaining).
- Validation guards with `IsValidationEnabled()`; errors accumulate via `AddError()` and surface from `Build()`.
- `Build()` returns a defensive copy of the entity.
- Builders are not thread-safe.
- Update `CHANGELOG.md` under `[Unreleased]` for user-facing changes.

## CI

`.github/workflows/default.yaml` delegates to `rios0rios0/pipelines/.github/workflows/go-library.yaml`. Runs on push/PR to `main` and tags.
