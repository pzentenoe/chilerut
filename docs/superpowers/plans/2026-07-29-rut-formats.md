# RUT Formats Implementation Plan

> **For agentic workers:** REQUIRED: Use subagent-driven development when subagents are available. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Add compact formatting, dotted display formatting, and strict RUT validation without changing the permissive API.

**Architecture:** Keep the existing `digits` and `split` normalization path as the compatibility layer. Add `Compact` and `FormatWithDots` on top of it. Add `ValidStrict` behind an explicit grammar check, then reuse modulo 11 verification only after syntax is known to be valid.

**Tech Stack:** Go 1.18, standard library (`regexp`, `strings`, `testing`), native Go fuzzing.

---

## Chunk 1: Public APIs and Tests

### Task 1: Specify failing behavior

**Files:**
- Modify: `rut_test.go`

- [ ] Add table cases for `Compact`, `FormatWithDots`, and `ValidStrict`.
- [ ] Include compact, hyphenated, dotted, lowercase `k`, outer Unicode whitespace,
  leading-zero, malformed-separator, lone-DV (`"K"` and `"-K"`), empty, and
  all-zero cases.
- [ ] Run `go test ./...` and confirm compilation fails because the new APIs do not exist.

### Task 2: Implement minimal API

**Files:**
- Modify: `rut.go`

- [ ] Add `Compact` using existing normalized body/DV splitting.
- [ ] Add `FormatWithDots` with right-aligned three-digit grouping.
- [ ] Add `ValidStrict` using the documented strict syntax and existing modulo 11 verification.
- [ ] Keep `Valid` and `Format` unchanged in behavior.
- [ ] Run `go test ./...` and confirm all table tests pass.

## Chunk 2: Strict Fuzzing and Verification

### Task 3: Add an independent strict-validation oracle

**Files:**
- Modify: `rut_test.go`

- [ ] Add a strict syntax regular expression owned by the test.
- [ ] Add an independent modulo 11 helper owned by the test.
- [ ] Add `FuzzValidStrict`; model Unicode trimming independently, then assert
  acceptance iff the test grammar, a non-zero normalized body, and the test
  modulo 11 calculation all pass.
- [ ] Run `go test -fuzz=FuzzValidStrict -fuzztime=10s .`.

### Task 4: Full validation and release hygiene

**Files:**
- Modify: `rut.go`, `rut_test.go`

- [ ] Run `go vet ./...`, `go test -race -coverprofile=coverage.out ./...`, and `golangci-lint run`.
- [ ] Run existing `FuzzValid` and `FuzzFormat` for 10 seconds each to preserve
  their panic-free and safe-output contracts.
- [ ] Confirm statement coverage remains 100%.
- [ ] Commit the API and test changes separately from the design documents.
