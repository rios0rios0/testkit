# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed
- Restructured project to align with standard package format
- Moved source code to `pkg/test/` directory
- Moved tests to same package for internal field access

## [1.0.0] - Initial Release

### Added
- Modular Builder Pattern with extensible base builder
- Factory System for dynamic builder creation and management
- Configuration Management for applying default values and settings
- Validation Framework with built-in error accumulation
- Tag System for metadata support
- Clone & Reset capabilities for deep copy and state management