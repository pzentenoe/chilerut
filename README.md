# chilerut

# chilerut

[![CI](https://github.com/pzentenoe/chilerut/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/pzentenoe/chilerut/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/pzentenoe/chilerut/branch/main/graph/badge.svg)](https://codecov.io/gh/pzentenoe/chilerut)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/pzentenoe/chilerut.svg)](https://pkg.go.dev/github.com/pzentenoe/chilerut)
[![Go 1.18+](https://img.shields.io/badge/go-%3E%3D1.18-blue)](go.mod)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

[![FOSSA License](https://app.fossa.com/api/projects/custom%2B60042%2Fgithub.com%2Fpzentenoe%2Fchilerut.svg?type=shield&issueType=license)](https://app.fossa.com/projects/custom%2B60042%2Fgithub.com%2Fpzentenoe%2Fchilerut?ref=badge_shield&issueType=license)
[![FOSSA Security](https://app.fossa.com/api/projects/custom%2B60042%2Fgithub.com%2Fpzentenoe%2Fchilerut.svg?type=shield&issueType=security)](https://app.fossa.com/projects/custom%2B60042%2Fgithub.com%2Fpzentenoe%2Fchilerut?ref=badge_shield&issueType=security)

Go package to validate, format and compare Chilean RUTs. Zero dependencies, 100% test coverage, fuzz-tested.

## Installation

```sh
go get github.com/pzentenoe/chilerut
```

Requires Go 1.18+.

## Usage

```go
import "github.com/pzentenoe/chilerut"
```

### Validate RUT

```go
chilerut.Valid("12345678-9")   // false
chilerut.Valid("6265837-1")    // true

// It works with the following formats:
chilerut.Valid("98685030")     // true
chilerut.Valid("9868503-0")    // true
chilerut.Valid("9.868.503-0")  // true
chilerut.Valid("12-667-869-K") // true
chilerut.Valid("12*667*869-k") // true
```

`Valid` is permissive and normalizes common separators. Use `ValidStrict` at
API or form boundaries when arbitrary characters must be rejected:

```go
chilerut.ValidStrict("12.667.869-K") // true
chilerut.ValidStrict("12667869K")     // true
chilerut.ValidStrict("12*667*869*K")  // false
```

### Input rules

- `Valid` is permissive: it normalizes common separators and ignores other
  non-significant characters for compatibility.
- `ValidStrict` accepts only compact (`12667869K`), hyphenated
  (`12667869-K`), or conventionally dotted (`12.667.869-K`) input. It trims
  outer whitespace, accepts lowercase `k`, and rejects malformed separators,
  all-zero bodies, and lone verification digits.
- `Format`, `Compact`, and `FormatWithDots` normalize input but do not verify
  the verification digit. Use `Valid` or `ValidStrict` when validity matters.

### Get verification digit

```go
chilerut.VerificationDigit("9868503")  // "0"
chilerut.VerificationDigit("12667869") // "K"
chilerut.VerificationDigit("16647869") // "3"
```

Pass the RUT body without its verification digit. To build a compact RUT with
its digit:

```go
body := "12667869"
rut := body + chilerut.VerificationDigit(body) // "12667869K"
```

### Format RUT

```go
chilerut.Format("12.667.869.k")        // "12667869-K"
chilerut.Format("12-667-869-k")        // "12667869-K"
chilerut.Format("12*667*869*K")        // "12667869-K"
chilerut.Format("12 667 869 k")        // "12667869-K"
chilerut.Format("12667869k")           // "12667869-K"
chilerut.Format("   000012667869k   ") // "12667869-K"
```

### Compact and display formats

```go
chilerut.Compact("12.667.869-k")      // "12667869K"
chilerut.FormatWithDots("12667869K")  // "12.667.869-K"
```

`Compact` always includes the final significant character as the verification
digit. All-zero input normalizes to an empty string.

### Compare RUT

```go
chilerut.Compare("12667869k", "12.667.869-K")    // true
chilerut.Compare("12667869-k", "12.667.869-K")   // true
chilerut.Compare("12.667.861-K", "12.667.869-K") // false
```

## Development

Run the test suite:

```sh
go test ./...
```

Measure public API performance and allocations:

```sh
go test -run='^$' -bench=. -benchmem -count=5 .
```

### Baseline

Measured on an Apple M4 Pro with Go 1.26.5 on 2026-07-29. These numbers are
reference measurements, not performance guarantees; compare changes on the
same machine and Go version.

| API | Time | Allocations |
| --- | ---: | ---: |
| `VerificationDigit` | ~4 ns/op | 0 |
| `Valid` | ~35 ns/op | 1 |
| `ValidStrict` (compact) | ~124 ns/op | 1 |
| `ValidStrict` (dotted) | ~179 ns/op | 1 |
| `Format` | ~48 ns/op | 2 |
| `Compact` | ~46 ns/op | 2 |
| `FormatWithDots` | ~61 ns/op | 3 |
| `Compare` | ~83 ns/op | 2 |



Go package to validate, format and compare Chilean RUTs. Zero dependencies, 100% test coverage, fuzz-tested.

## Installation

```sh
go get github.com/pzentenoe/chilerut
```

Requires Go 1.18+.

## Usage

```go
import "github.com/pzentenoe/chilerut"
```

### Validate RUT

```go
chilerut.Valid("12345678-9")   // false
chilerut.Valid("6265837-1")    // true

// It works with the following formats:
chilerut.Valid("98685030")     // true
chilerut.Valid("9868503-0")    // true
chilerut.Valid("9.868.503-0")  // true
chilerut.Valid("12-667-869-K") // true
chilerut.Valid("12*667*869-k") // true
```

`Valid` is permissive and normalizes common separators. Use `ValidStrict` at
API or form boundaries when arbitrary characters must be rejected:

```go
chilerut.ValidStrict("12.667.869-K") // true
chilerut.ValidStrict("12667869K")     // true
chilerut.ValidStrict("12*667*869*K")  // false
```

### Input rules

- `Valid` is permissive: it normalizes common separators and ignores other
  non-significant characters for compatibility.
- `ValidStrict` accepts only compact (`12667869K`), hyphenated
  (`12667869-K`), or conventionally dotted (`12.667.869-K`) input. It trims
  outer whitespace, accepts lowercase `k`, and rejects malformed separators,
  all-zero bodies, and lone verification digits.
- `Format`, `Compact`, and `FormatWithDots` normalize input but do not verify
  the verification digit. Use `Valid` or `ValidStrict` when validity matters.

### Get verification digit

```go
chilerut.VerificationDigit("9868503")  // "0"
chilerut.VerificationDigit("12667869") // "K"
chilerut.VerificationDigit("16647869") // "3"
```

Pass the RUT body without its verification digit. To build a compact RUT with
its digit:

```go
body := "12667869"
rut := body + chilerut.VerificationDigit(body) // "12667869K"
```

### Format RUT

```go
chilerut.Format("12.667.869.k")        // "12667869-K"
chilerut.Format("12-667-869-k")        // "12667869-K"
chilerut.Format("12*667*869*K")        // "12667869-K"
chilerut.Format("12 667 869 k")        // "12667869-K"
chilerut.Format("12667869k")           // "12667869-K"
chilerut.Format("   000012667869k   ") // "12667869-K"
```

### Compact and display formats

```go
chilerut.Compact("12.667.869-k")      // "12667869K"
chilerut.FormatWithDots("12667869K")  // "12.667.869-K"
```

`Compact` always includes the final significant character as the verification
digit. All-zero input normalizes to an empty string.

### Compare RUT

```go
chilerut.Compare("12667869k", "12.667.869-K")    // true
chilerut.Compare("12667869-k", "12.667.869-K")   // true
chilerut.Compare("12.667.861-K", "12.667.869-K") // false
```

## Development

Run the test suite:

```sh
go test ./...
```

Measure public API performance and allocations:

```sh
go test -run='^$' -bench=. -benchmem -count=5 .
```

### Baseline

Measured on an Apple M4 Pro with Go 1.26.5 on 2026-07-29. These numbers are
reference measurements, not performance guarantees; compare changes on the
same machine and Go version.

| API | Time | Allocations |
| --- | ---: | ---: |
| `VerificationDigit` | ~4 ns/op | 0 |
| `Valid` | ~35 ns/op | 1 |
| `ValidStrict` (compact) | ~124 ns/op | 1 |
| `ValidStrict` (dotted) | ~179 ns/op | 1 |
| `Format` | ~48 ns/op | 2 |
| `Compact` | ~46 ns/op | 2 |
| `FormatWithDots` | ~61 ns/op | 3 |
| `Compare` | ~83 ns/op | 2 |
