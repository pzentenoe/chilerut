package chilerut

import (
	"testing"
)

func TestFormat(t *testing.T) {
	type args struct {
		rut string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"12.667.869-k",
			args{
				rut: "12.667.869.k",
			},
			"12667869-K",
		},
		{
			"12-667-869-k",
			args{
				rut: "12-667-869-k",
			},
			"12667869-K",
		},
		{
			"12*667*869*K",
			args{
				rut: "12*667*869*K",
			},
			"12667869-K",
		},
		{
			"12 667 869 k",
			args{
				rut: "12 667 869 k",
			},
			"12667869-K",
		},
		{
			"   000012667869k   ",
			args{
				rut: "   000012667869k   ",
			},
			"12667869-K",
		},
		{
			"98685030",
			args{
				rut: "98685030",
			},
			"9868503-0",
		},
		{
			"empty",
			args{
				rut: "",
			},
			"",
		},
		{
			"only zeros",
			args{
				rut: "0000",
			},
			"",
		},
		{
			"tabs and newlines",
			args{
				rut: "\t98685030\n",
			},
			"9868503-0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Format(tt.args.rut); got != tt.want {
				t.Errorf("Format() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVerificationDigit(t *testing.T) {
	type args struct {
		rut string
	}
	tests := []struct {
		name   string
		args   args
		wantDv string
	}{
		{
			"9868503(0)",
			args{
				rut: "9868503",
			},
			"0",
		},
		{
			"12667869(K)",
			args{
				rut: "12667869",
			},
			"K",
		},
		{
			"16.647.869(3)",
			args{
				rut: "16.647.869",
			},
			"3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDv := VerificationDigit(tt.args.rut); gotDv != tt.wantDv {
				t.Errorf("VerificationDigit() = %v, want %v", gotDv, tt.wantDv)
			}
		})
	}
}

func TestValid(t *testing.T) {
	type args struct {
		rut string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"12345678-9",
			args{
				rut: "12345678-9",
			},
			false,
		},
		{
			"6265837-1",
			args{
				rut: "6265837-1",
			},
			true,
		},
		{
			"98685030",
			args{
				rut: "98685030",
			},
			true,
		},
		{
			"9868503-0",
			args{
				rut: "9868503-0",
			},
			true,
		},
		{
			"9.868.503-0",
			args{
				rut: "9.868.503-0",
			},
			true,
		},
		{
			"12.667.869-K",
			args{
				rut: "12.667.869-K",
			},
			true,
		},
		{
			"12.667.869-k",
			args{
				rut: "12.667.869-k",
			},
			true,
		},
		{
			"11.111.111-1",
			args{
				rut: "11.111.111-1",
			},
			true,
		},
		{
			"9999999-3",
			args{
				rut: "9999999-3",
			},
			true,
		},
		{
			"wrong dv",
			args{
				rut: "12667869-0",
			},
			false,
		},
		{
			"trailing garbage",
			args{
				rut: "123X",
			},
			false,
		},
		{
			"empty",
			args{
				rut: "",
			},
			false,
		},
		{
			"only dv",
			args{
				rut: "K",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Valid(tt.args.rut); got != tt.want {
				t.Errorf("Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompare(t *testing.T) {
	type args struct {
		rut1 string
		rut2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"12.667.869-K == 12.667.869-k",
			args{
				rut1: "12.667.869-K",
				rut2: "12.667.869-k",
			},
			true,
		},
		{
			"12.667.869-K != 12.667.861-K",
			args{
				rut1: "12.667.869-K",
				rut2: "12.667.861-K",
			},
			false,
		},
		{
			"12667869K == 12.667.869-k",
			args{
				rut1: "12667869K",
				rut2: "12.667.869-k",
			},
			true,
		},
		{
			"12.667.869-K == 12667869k",
			args{
				rut1: "12.667.869-K",
				rut2: "12667869k",
			},
			true,
		},
		{
			"98685030 == 9.868.503-0",
			args{
				rut1: "98685030",
				rut2: "9.868.503-0",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Compare(tt.args.rut1, tt.args.rut2); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
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
		out := Format(rut)
		for i := 0; i < len(out); i++ {
			if c := out[i]; !isNumeric(c) && c != '-' && c != 'K' {
				t.Fatalf("Format(%q) contains %q", rut, c)
			}
		}
	})
}
