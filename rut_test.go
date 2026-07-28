package chilerut

import "testing"

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
