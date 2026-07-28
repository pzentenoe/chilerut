# chilerut

[![CI](https://github.com/pzentenoe/chilerut/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/pzentenoe/chilerut/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/pzentenoe/chilerut/graph/badge.svg?branch=main)](https://codecov.io/gh/pzentenoe/chilerut)
[![Go Report Card](https://goreportcard.com/badge/github.com/pzentenoe/chilerut)](https://goreportcard.com/report/github.com/pzentenoe/chilerut)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/jonathanhecl/chilerut.svg)](https://pkg.go.dev/github.com/jonathanhecl/chilerut)
[![Go 1.18+](https://img.shields.io/badge/go-%3E%3D1.18-blue)](go.mod)

Go package to validate, format and compare Chilean RUTs. Zero dependencies, 100% test coverage, fuzz-tested.

## Installation

```sh
go get github.com/jonathanhecl/chilerut
```

Requires Go 1.18+.

## Usage

```go
import "github.com/jonathanhecl/chilerut"
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

### Get verification digit

```go
chilerut.VerificationDigit("9868503")  // "0"
chilerut.VerificationDigit("12667869") // "K"
chilerut.VerificationDigit("16647869") // "3"
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

### Compare RUT

```go
chilerut.Compare("12667869k", "12.667.869-K")    // true
chilerut.Compare("12667869-k", "12.667.869-K")   // true
chilerut.Compare("12.667.861-K", "12.667.869-K") // false
```
