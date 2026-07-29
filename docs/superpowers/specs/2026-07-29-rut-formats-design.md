# RUT Formats Design

## Goal

Add explicit compact, display, and strict-validation APIs while preserving the
current permissive behavior of `Valid` and `Format`.

## Public API

- `Compact(rut string) string` returns the normalized RUT without separators,
  including its presumed verification digit: `"12667869K"`.
- `FormatWithDots(rut string) string` returns the conventional display form:
  `"12.667.869-K"`.
- `ValidStrict(rut string) bool` validates both the modulo 11 verification
  digit and accepted input syntax.

## Accepted Strict Syntax

`ValidStrict` trims outer Unicode whitespace and accepts only these forms
(where `D` is `[0-9]` and `V` is `[0-9Kk]`):

- Compact: `D+V`, such as `12667869K`
- Compact with a hyphen: `D+-V`, such as `12667869-K`
- Conventional dots and hyphen: `D{1,3}(\.D{3})+-V`, such as
  `12.667.869-K`

It rejects arbitrary separators, misplaced dots or hyphens, and non-RUT
characters. Lowercase `k` is accepted and normalized. It does not impose a
maximum body-length policy, preserving support for valid historical or test
RUT lengths; it requires at least one non-zero body digit and one verification
digit. An all-zero body is invalid.

## Compatibility

`Valid` and `Format` remain permissive and continue normalizing existing
inputs such as `"12*667*869*k"`. The new APIs reuse existing normalization
and modulo 11 logic; no new type or parsing-error API is introduced.

`Compact` and `FormatWithDots` normalize only; they do not verify the check
digit. As with the existing `Format`, the final significant character is
treated as the verification digit. They return an empty string when no
significant character remains after normalization, including all-zero input.
For a lone verification digit, such as `"K"` or `"-K"`, `Compact` returns
`"K"` and `FormatWithDots` returns `"-K"`; `ValidStrict` rejects it because
there is no body digit.

## Testing

- Table tests cover all accepted strict forms, malformed separators, compact
  output, and dot grouping.
- Existing fuzz tests continue to enforce no panics and safe output.
- A strict-validation fuzz target uses an independent test regular expression
  for the accepted grammar and an independent modulo 11 calculation, then
  verifies arbitrary text is never accepted unless it matches that grammar and
  has a valid digit.
