package echo

import "testing"

// -- Benchmarks --

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo1()
	}
}

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo2()
	}
}

func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo3()
	}
}

/*
BenchmarkEcho1
   20000             58603 ns/op
PASS
ok      _/D_/gowork/gopl-ex/ch1/ex1_3   2.222s


BenchmarkEcho2
   20000             59653 ns/op
PASS
ok      _/D_/gowork/gopl-ex/ch1/ex1_3   2.269s

BenchmarkEcho3
   20000             58603 ns/op
PASS
ok      _/D_/gowork/gopl-ex/ch1/ex1_3   2.251s

*/
