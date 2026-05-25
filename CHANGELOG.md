# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

When a new release is proposed:

1. Create a new branch `bump/x.x.x` (this isn't a long-lived branch!!!);
2. The Unreleased section on `CHANGELOG.md` gets a version number and date;
3. Open a Pull Request with the bump version changes targeting the `main` branch;
4. When the Pull Request is merged, a new Git tag must be created using [GitHub environment](https://github.com/rios0rios0/testkit/tags).

Releases to productive environments should run from a tagged version.
Exceptions are acceptable depending on the circumstances (critical bug fixes that can be cherry-picked, etc.).

## [Unreleased]

### Changed

- refreshed `.github/copilot-instructions.md` to fix stale Go version references (`1.26.0+` → `1.26.3+`)

## [0.2.1] - 2026-05-08

### Changed

- changed the Go version to `1.26.3`
- refreshed `.github/copilot-instructions.md` and `CLAUDE.md` to reflect the Go `1.26.3` version requirement

## [0.2.0] - 2026-04-28

### Added

- added `CLAUDE.md` with project guidance for Claude Code sessions

### Changed

- refreshed `.github/copilot-instructions.md` to reflect Go 1.26.2 version requirement

## [0.1.2] - 2026-04-15

### Changed

- changed the Go version to `1.26.2` and updated all module dependencies

## [0.1.1] - 2026-03-12

### Changed

- changed the Go version to `1.26.1` and updated all module dependencies

## [0.1.0] - 2026-03-06

### Added

- added clone and reset capabilities for deep copy and state management
- added Configuration Management for applying default values and settings
- added Factory System for dynamic builder creation and management
- added modular Builder Pattern with extensible base builder
- added Tag System for metadata support
- added Validation Framework with built-in error accumulation

### Changed

- changed the Go version to `1.26.0` and updated all module dependencies
- moved tests to same package for internal field access
- restructured project to align with standard package format (`pkg/test/`)

### Fixed

- fixed 41 lint findings including `forbidigo`, `errcheck`, `govet`, `funlen`, `cyclop`, `mnd`, and `nestif` violations
- fixed test compilation error caused by `errors` variable shadowing `errors` package import in builder tests
