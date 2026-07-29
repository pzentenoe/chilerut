package chilerut

import "testing"

const (
	benchmarkCompactRUT   = "12667869K"
	benchmarkFormattedRUT = "12.667.869-K"
	benchmarkRUTBody      = "12667869"
)

var (
	benchmarkBoolResult   bool
	benchmarkStringResult string
)

func BenchmarkValid(b *testing.B) {
	for _, tt := range []struct {
		name string
		rut  string
	}{
		{"compact", benchmarkCompactRUT},
		{"formatted", benchmarkFormattedRUT},
	} {
		b.Run(tt.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				benchmarkBoolResult = Valid(tt.rut)
			}
		})
	}
}

func BenchmarkValidStrict(b *testing.B) {
	for _, tt := range []struct {
		name string
		rut  string
	}{
		{"compact", benchmarkCompactRUT},
		{"formatted", benchmarkFormattedRUT},
	} {
		b.Run(tt.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				benchmarkBoolResult = ValidStrict(tt.rut)
			}
		})
	}
}

func BenchmarkFormat(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchmarkStringResult = Format(benchmarkFormattedRUT)
	}
}

func BenchmarkCompact(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchmarkStringResult = Compact(benchmarkFormattedRUT)
	}
}

func BenchmarkFormatWithDots(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchmarkStringResult = FormatWithDots(benchmarkCompactRUT)
	}
}

func BenchmarkVerificationDigit(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchmarkStringResult = VerificationDigit(benchmarkRUTBody)
	}
}

func BenchmarkCompare(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchmarkBoolResult = Compare(benchmarkCompactRUT, benchmarkFormattedRUT)
	}
}
