// Package chilerut validates, formats and compares Chilean RUT
// (Rol Único Tributario) numbers using the standard módulo 11
// check-digit algorithm.
package chilerut

import (
	"strconv"
	"strings"
)

// Format returns the canonical "body-dv" form of a RUT (e.g. "12667869-K"),
// dropping any separators, leading zeros and lowercasing.
func Format(rut string) string {
	body, dv := split(rut)
	if dv == "" {
		return ""
	}
	return body + "-" + dv
}

// Valid reports whether rut is a valid RUT by recomputing its módulo 11
// check digit. Any common separator format is accepted.
func Valid(rut string) bool {
	body, dv := split(rut)
	if body == "" || dv == "" {
		return false
	}
	return VerificationDigit(body) == dv
}

// VerificationDigit computes the expected módulo 11 check digit ("0"-"9"
// or "K") for the given RUT body.
func VerificationDigit(rut string) string {
	sum, factor := 0, 2
	for i := len(rut) - 1; i >= 0; i-- {
		c := rut[i]
		if c < '0' || c > '9' {
			continue
		}
		sum += int(c-'0') * factor
		if factor == 7 {
			factor = 2
		} else {
			factor++
		}
	}
	switch dv := 11 - sum%11; dv {
	case 11:
		return "0"
	case 10:
		return "K"
	default:
		return strconv.Itoa(dv)
	}
}

// Compare reports whether two RUTs are the same, ignoring format differences.
func Compare(rut1, rut2 string) bool {
	return Format(rut1) == Format(rut2)
}

// split normalizes a RUT into its numeric body and check digit.
func split(rut string) (body, dv string) {
	s := strings.TrimLeft(digits(rut), "0")
	if s == "" {
		return "", ""
	}
	return s[:len(s)-1], s[len(s)-1:]
}

// digits extracts the significant characters of a RUT: decimal digits and
// an optional trailing 'K' check digit.
func digits(rut string) string {
	rut = strings.ToUpper(strings.TrimSpace(rut))
	var b strings.Builder
	b.Grow(len(rut))
	for i := 0; i < len(rut); i++ {
		if c := rut[i]; c >= '0' && c <= '9' || c == 'K' && i == len(rut)-1 {
			b.WriteByte(c)
		}
	}
	return b.String()
}
