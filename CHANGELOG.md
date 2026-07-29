# Changelog

All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Public API benchmarks with allocation reporting.
- FOSSA dependency license and security scanning on pushes to `main`.

## [1.1.0] - 2026-07-29

### Added

- `Compact` for normalized RUTs without separators and with their verification digit.
- `FormatWithDots` for conventional Chilean RUT display formatting.
- `ValidStrict` for syntax-aware validation at API and form boundaries.

### Fixed

- Run each fuzz target independently in CI.

## [1.0.0] - 2026-07-28

First stable release as an independent Go module:
`github.com/pzentenoe/chilerut`.

### Added

- `Valid`, `Format`, `VerificationDigit`, and `Compare` for Chilean RUTs.
- Standard módulo 11 verification-digit calculation.
- Fuzz tests for validation and formatting.
- GitHub Actions CI with Go 1.18 and stable, race detection, coverage, and golangci-lint.
- Automated GitHub releases with GoReleaser.

### Fixed

- Prevented `Valid` from panicking on malformed input with trailing garbage.

[Unreleased]: https://github.com/pzentenoe/chilerut/compare/v1.1.0...HEAD
[1.1.0]: https://github.com/pzentenoe/chilerut/compare/v1.0.0...v1.1.0
[1.0.0]: https://github.com/pzentenoe/chilerut/releases/tag/v1.0.0
