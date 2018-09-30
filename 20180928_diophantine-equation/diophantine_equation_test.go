package diophantine_equation

import "testing"

var n = 90009

func BenchmarkSolequaV1(b *testing.B) {
	b.ResetTimer()
	for i:=0; i < b.N; i++{
		SolequaV1(n)
	}
}

func BenchmarkSolequaV2(b *testing.B) {
	for i:=0; i < b.N; i++{
		SolequaV2(n)
	}
}

func BenchmarkSolequaV3(b *testing.B) {
	for i:=0; i < b.N; i++{
		SolequaV3(n)
	}
}

func BenchmarkSolequaV4(b *testing.B) {
	for i:=0; i < b.N; i++{
		SolequaV4(n)
	}
}
