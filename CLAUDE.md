# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project

Go testing utility library (`github.com/rios0rios0/testkit`). Provides a modular builder framework for test environment setup. Pure library -- no deployable artifacts. Requires Go 1.27.0+.

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
- Write a changelog fragment for every change — `chlog new --kind <Kind> --body "..."`, committed from `.changes/unreleased/`. Never edit `CHANGELOG.md`: it is generated from the fragments at release time by `chlog batch auto && chlog merge`.

## CI

`.github/workflows/default.yaml` delegates to `rios0rios0/pipelines/.github/workflows/go-library.yaml`. Runs on push/PR to `main` and tags.

<!-- chlog:start -->
## Changelog (chlog) — MANDATORY

If the repository you are working in uses chlog (a `.chlog.yaml` or `.chlog.yml`
config file, or a `.changes/` directory, exists at the project root), the
following is binding and ALWAYS applies: whenever you make ANY change, you MUST
create a changelog fragment as part of the same change — automatically, without
being asked, before committing.

- Do NOT edit CHANGELOG.md directly; it is generated from fragments.
- Create the fragment with:
  `chlog new --kind <Kind> --body "<imperative description>"`
- Valid kinds: Added, Changed, Deprecated, Removed, Fixed, Security
- Choose the kind that best matches the change (e.g., new feature → Added,
  bug fix → Fixed, behavior change → Changed, removal → Removed, security fix → Security).
- If the change is backward-INCOMPATIBLE with the public API (a breaking
  change), you MUST add the `--breaking` flag:
  `chlog new --kind <Kind> --breaking --body "<description>"`.
  This is the ONLY thing that triggers a major version bump — the kind alone
  never does (per SemVer, major = incompatible change). When unsure whether a
  change breaks compatibility, ask the user instead of guessing.
- Fragments are YAML files in `.changes/unreleased/`; stage them with your commit.
- `chlog check` fails the build when a fragment is missing — never skip it.
<!-- chlog:end -->
