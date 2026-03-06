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

## [1.0.0] - 2026-03-06

### Added

- Modular Builder Pattern with extensible base builder
- Factory System for dynamic builder creation and management
- Configuration Management for applying default values and settings
- Validation Framework with built-in error accumulation
- Tag System for metadata support
- Clone & Reset capabilities for deep copy and state management

### Changed

- Restructured project to align with standard package format
- Moved source code to `pkg/test/` directory
- Moved tests to same package for internal field access
- changed the Go version to `1.26.0` and updated all module dependencies

### Fixed

- fixed test compilation error caused by `errors` variable shadowing `errors` package import in builder tests
- fixed `errcheck` findings by handling unchecked error returns in examples and factory code
- fixed `govet` shadow findings by using distinct variable names and type switches
- fixed `funlen` and `cyclop` findings by splitting monolithic `main()` into helper functions
- fixed `mnd` findings by extracting magic numbers into named constants
- fixed `nestif` finding by extracting `applyDefaults()` method to reduce nesting depth

