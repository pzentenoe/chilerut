package chilerut

import (
	"regexp"
	"strings"
	"testing"
)

var strictRUTSyntax = regexp.MustCompile(`^(?:[0-9]+[0-9Kk]|[0-9]+-[0-9Kk]|[0-9]{1,3}(?:\.[0-9]{3})+-[0-9Kk])$`)

func TestFormat(t *testing.T) {
	tests := []struct{ name, in, want string }{
		{"dotted lowercase k", "12.667.869.k", "12667869-K"},
		{"dashed", "12-667-869-k", "12667869-K"},
		{"asterisks", "12*667*869*K", "12667869-K"},
		{"spaces", "12 667 869 k", "12667869-K"},
		{"padded with zeros", "   000012667869k   ", "12667869-K"},
		{"numeric check digit", "98685030", "9868503-0"},
		{"tabs and newlines", "\t98685030\n", "9868503-0"},
		{"empty", "", ""},
		{"only zeros", "0000", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Format(tt.in); got != tt.want {
				t.Errorf("Format(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestCompact(t *testing.T) {
	tests := []struct{ name, in, want string }{
		{"dotted lowercase k", "12.667.869-k", "12667869K"},
		{"already compact", "12667869K", "12667869K"},
		{"permissive separator", "12*667*869*K", "12667869K"},
		{"leading zeros", "000012667869K", "12667869K"},
		{"lone check digit", "K", "K"},
		{"lone check digit with separator", "-K", "K"},
		{"only zeros", "0000", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Compact(tt.in); got != tt.want {
				t.Errorf("Compact(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestFormatWithDots(t *testing.T) {
	tests := []struct{ name, in, want string }{
		{"compact", "12667869K", "12.667.869-K"},
		{"hyphenated", "12667869-K", "12.667.869-K"},
		{"seven digit body", "98685030", "9.868.503-0"},
		{"body grouped in threes", "1234567", "123.456-7"},
		{"leading zeros", "000012667869K", "12.667.869-K"},
		{"lone check digit", "K", "-K"},
		{"only zeros", "0000", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatWithDots(tt.in); got != tt.want {
				t.Errorf("FormatWithDots(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestVerificationDigit(t *testing.T) {
	tests := []struct{ name, in, want string }{
		{"dv zero", "9868503", "0"},
		{"dv K", "12667869", "K"},
		{"dv numeric", "16.647.869", "3"},
		{"all ones", "11111111", "1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VerificationDigit(tt.in); got != tt.want {
				t.Errorf("VerificationDigit(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestValid(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want bool
	}{
		{"wrong check digit", "12345678-9", false},
		{"plain", "6265837-1", true},
		{"no separator", "98685030", true},
		{"dash", "9868503-0", true},
		{"dotted", "9.868.503-0", true},
		{"uppercase K", "12.667.869-K", true},
		{"lowercase k", "12.667.869-k", true},
		{"all ones", "11.111.111-1", true},
		{"all nines", "9999999-3", true},
		{"valid body wrong dv", "12667869-0", false},
		{"trailing garbage", "123X", false},
		{"empty", "", false},
		{"only check digit", "K", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Valid(tt.in); got != tt.want {
				t.Errorf("Valid(%q) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestValidStrict(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want bool
	}{
		{"compact", "12667869K", true},
		{"hyphenated", "12667869-K", true},
		{"dotted", "12.667.869-K", true},
		{"lowercase k", "12.667.869-k", true},
		{"outer unicode whitespace", "\u00a0 12.667.869-k \u00a0", true},
		{"leading zeros", "000001-9", true},
		{"permissive separator", "12*667*869*K", false},
		{"embedded spaces", "12 667 869-K", false},
		{"misgrouped dots", "1.266.786.9-K", false},
		{"dotted without hyphen", "12.667.869K", false},
		{"double hyphen", "12.667.869--K", false},
		{"trailing garbage", "123X", false},
		{"lone check digit", "K", false},
		{"lone check digit with separator", "-K", false},
		{"empty", "", false},
		{"all zeros", "0000", false},
		{"all zero body", "000000-0", false},
		{"wrong check digit", "12667869-0", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidStrict(tt.in); got != tt.want {
				t.Errorf("ValidStrict(%q) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestCompare(t *testing.T) {
	tests := []struct {
		name      string
		rut1      string
		rut2      string
		wantEqual bool
	}{
		{"same dv different case", "12.667.869-K", "12.667.869-k", true},
		{"different body", "12.667.869-K", "12.667.861-K", false},
		{"plain vs dotted", "12667869K", "12.667.869-k", true},
		{"dotted vs plain", "12.667.869-K", "12667869k", true},
		{"no separator vs dotted", "98685030", "9.868.503-0", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Compare(tt.rut1, tt.rut2); got != tt.wantEqual {
				t.Errorf("Compare(%q, %q) = %v, want %v", tt.rut1, tt.rut2, got, tt.wantEqual)
			}
		})
	}
}

// Valid must never panic, regardless of input.
func FuzzValid(f *testing.F) {
	f.Add("12.667.869-K")
	f.Add("98685030")
	f.Add("123X")
	f.Add("")
	f.Fuzz(func(t *testing.T, rut string) {
		Valid(rut)
	})
}

// Format must only output digits, '-' and 'K'.
func FuzzFormat(f *testing.F) {
	f.Add("12.667.869-k")
	f.Add("   000012667869k   ")
	f.Fuzz(func(t *testing.T, rut string) {
		for _, c := range Format(rut) {
			if c != '-' && c != 'K' && (c < '0' || c > '9') {
				t.Fatalf("Format(%q) produced %q", rut, c)
			}
		}
	})
}

func FuzzValidStrict(f *testing.F) {
	f.Add("12667869K")
	f.Add("12.667.869-k")
	f.Add("12*667*869*K")
	f.Add("000000-0")
	f.Fuzz(func(t *testing.T, rut string) {
		trimmed := strings.TrimSpace(rut)
		want := false
		if strictRUTSyntax.MatchString(trimmed) {
			body, dv := strictParts(trimmed)
			want = body != "" && strictVerificationDigit(body) == dv
		}
		if got := ValidStrict(rut); got != want {
			t.Errorf("ValidStrict(%q) = %v, want %v", rut, got, want)
		}
	})
}

func strictParts(rut string) (body, dv string) {
	rut = strings.NewReplacer(".", "", "-", "").Replace(strings.ToUpper(rut))
	return strings.TrimLeft(rut[:len(rut)-1], "0"), rut[len(rut)-1:]
}

func strictVerificationDigit(body string) string {
	total := 0
	for offset := 0; offset < len(body); offset++ {
		total += int(body[len(body)-1-offset]-'0') * (2 + offset%6)
	}
	switch digit := 11 - total%11; digit {
	case 11:
		return "0"
	case 10:
		return "K"
	default:
		return string(rune('0' + digit))
	}
}
