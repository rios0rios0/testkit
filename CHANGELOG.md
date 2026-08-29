# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

This file is not edited by hand. Every change writes its own fragment under
`.changes/unreleased/` with [chlog](https://github.com/luizjhonata/chlog), and a release compiles
the pending fragments into a version section here — so two branches each adding an entry no
longer touch the same lines, and a rebase that used to conflict on this file now conflicts on
nothing.

When a new release is proposed:

1. Create a new branch `bump/x.x.x` (this isn't a long-lived branch!!!);
2. The fragments pending under `.changes/unreleased/` are compiled into a version section by `chlog batch auto && chlog merge` (AutoBump does this for you — it reads the fragments directly);
3. Open a Pull Request with the bump version changes targeting the `main` branch;
4. When the Pull Request is merged, a new Git tag must be created using [GitHub environment](https://github.com/rios0rios0/testkit/tags).

Releases to productive environments should run from a tagged version.
Exceptions are acceptable depending on the circumstances (critical bug fixes that can be cherry-picked, etc.).

## [Unreleased]

## [0.3.1] - 2026-08-28

### Changed

- changed the Claude workflows to call the reusable workflows in `rios0rios0/pipelines` instead of `rios0rios0/.github`, which is where every other reusable workflow and composite action already lives, and renamed them to `claude-review.yaml` and `claude-mention.yaml`, matching the `reusable-claude-review.yaml` / `reusable-claude-mention.yaml` definitions they call

### Fixed

- made `BuilderFactory` safe for concurrent use. Its map was unsynchronised while `DefaultFactory` is a package-level singleton, so any `t.Parallel()` tests in a package -- and any helper goroutines they start -- wrote the same map concurrently -- and an unsynchronised concurrent map access is a runtime throw rather than a panic, so `recover` cannot catch it and the whole test binary dies with `fatal error: concurrent map writes` and no failing test to point at. Reads take an `RWMutex` read lock, since registration typically happens once during setup and every lookup after it is a read. `Create` deliberately calls the creation function after releasing the lock, so a creation function that registers further builders -- a composite registering its parts -- does not deadlock against the non-reentrant lock.
- restored the `.changes/unreleased/` directory with a `.gitkeep`, so the release tooling keeps recognising this project as [chlog](https://github.com/luizjhonata/chlog)-based after a release consumes the last fragment. Git tracks files rather than directories, so the bump commit that removed the final fragment removed the directory too, and the next run read the empty `[Unreleased]` section as "nothing to release"
- restored the `id-token: write` permission on both Claude workflow callers. Without it the caller grants less than the reusable workflow declares, which GitHub rejects before the job starts -- runs ended in `startup_failure`. The action needs the scope because `setupGitHubToken()` exchanges a GitHub OIDC token for the GitHub App token it posts with, unless a `github_token` is passed explicitly.

### Removed

- removed the unused `id-token: write` permission from the Claude workflow callers, and changed `claude-review.yaml`'s display name to `Claude Review` so it matches its file name and its `Claude Mention` sibling. `anthropics/claude-code-action` needs `id-token: write` only for workload identity federation or the Bedrock / Vertex / Foundry OIDC paths; these authenticate with `claude_code_oauth_token`, so the scope allowed minting OIDC tokens for any audience without ever being used.

## [0.3.0] - 2026-08-26

### Added

- added a tailored `code-review` skill under `.github/skills/` so GitHub Copilot reviews changes against the [rios0rios0/guide](https://github.com/rios0rios0/guide/wiki) standards and this repository's own load-bearing invariants

### Changed

- changed the changelog to [chlog](https://github.com/luizjhonata/chlog) fragments: a change now writes its own YAML file under `.changes/unreleased/` through `chlog new --kind <Kind> --body "..."`, and `CHANGELOG.md` is GENERATED from them at release time by `chlog batch auto && chlog merge`. That is the one thing a single shared file cannot do — two branches each adding an entry no longer touch the same lines, so a rebase that used to conflict on `CHANGELOG.md` now conflicts on nothing. The `[Unreleased]` section was empty, so nothing had to be carried across. AutoBump already reads the fragments directly, so the release flow is unchanged.

### Fixed

- corrected the Go prerequisite in `CONTRIBUTING.md` to 1.27+, matching the version `go.mod` requires
- fixed the `main` pipeline, which every repository's `sast:gitleaks` job had been failing since the code-review skill landed: the skill's own security bullet listed credential prefixes verbatim to warn against writing them, and the scanner's second pass matches those prefixes on their own, so the warning tripped the rule it was describing. The bullet now names the vendors instead, and the commit that carried the original wording is allowlisted by fingerprint in `.gitleaksignore`, because the scan walks the whole history reachable from `HEAD` and no edit at the tip can clear a past commit. No credential was ever committed.

## [0.2.8] - 2026-08-24

### Changed

- changed the Go version to `1.27.0` and updated all module dependencies
- refreshed `CLAUDE.md` and `.github/copilot-instructions.md` to update the required Go version (`1.26.5+` → `1.27.0+`)

## [0.2.7] - 2026-08-15

### Changed

- changed the Go version to `1.26.6` and updated all module dependencies

## [0.2.6] - 2026-07-13

### Changed

- refreshed `CLAUDE.md` and `.github/copilot-instructions.md` to update the required Go version (`1.26.4+` → `1.26.5+`)

## [0.2.5] - 2026-07-10

### Changed

- changed the Go version to `1.26.5` and updated all module dependencies

### Security

- replaced `secrets: inherit` with an explicit `CLAUDE_CODE_OAUTH_TOKEN` secret in the Claude workflow callers, following the least-privilege principle

## [0.2.4] - 2026-06-09

### Changed

- refreshed `CLAUDE.md` and `.github/copilot-instructions.md` to update the required Go version (`1.26.3+` → `1.26.4+`)

## [0.2.3] - 2026-06-03

### Changed

- changed the Go version to `1.26.4` and updated all module dependencies

## [0.2.2] - 2026-05-25

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
