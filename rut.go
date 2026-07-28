package chilerut

import (
	"strconv"
	"strings"
)

func isNumeric(s byte) bool {
	return s >= '0' && s <= '9'
}

// Format returns the RUT as "body-dv" (e.g. "12667869-K"), stripping any
// separators, leading zeros and lowercasing.
func Format(rut string) string {
	rut = strings.TrimSpace(rut)
	rut = strings.TrimLeft(rut, "0")
	rut = strings.ToUpper(rut)
	rutFormated := ""
	for i := 0; i < len(rut); i++ {
		if i == len(rut)-1 {
			if isNumeric(rut[i]) {
				rutFormated += "-" + string(rut[i])
			}
			if rut[i] == 'K' {
				rutFormated += "-K"
			}
		} else if isNumeric(rut[i]) {
			rutFormated += string(rut[i])
		}
	}
	return rutFormated
}

// Valid reports whether rut is a valid Chilean RUT by checking its
// verification digit. Any common separator format is accepted.
func Valid(rut string) bool {
	rut = Format(rut)
	if len(rut) < 3 {
		return false
	}
	t := strings.Split(rut, "-")
	if len(t) != 2 {
		return false
	}
	return VerificationDigit(t[0]) == t[1]
}

// VerificationDigit returns the expected verification digit ("0"-"9" or "K")
// for the given RUT body.
func VerificationDigit(rut string) (dv string) {
	rut = strings.TrimSpace(rut)
	rut = strings.TrimLeft(rut, "0")
	f := 0
	sum := 0
	for i := len(rut) - 1; i >= 0; i-- {
		if isNumeric(rut[i]) {
			sum += int(rut[i]-'0') * (f + 2)
			f = (f + 1) % 6
		}
	}
	num := 11 - (sum % 11)
	if num < 10 {
		dv = strconv.Itoa(num)
	} else if num == 10 {
		dv = "K"
	} else {
		dv = "0"
	}
	return dv
}

// Compare reports whether two RUTs are equal, ignoring format differences.
func Compare(rut1, rut2 string) bool {
	return Format(rut1) == Format(rut2)
}
