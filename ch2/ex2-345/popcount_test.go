package popcount

import "testing"

var cases = []struct {
	input    uint64
	expected int
}{
	{0, 0},
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 1},
	{5, 2},
	{6, 2},
	{7, 3},
	{^uint64(0), 64},
}

// -- Functional Tests --

func TestPopCount(t *testing.T) {
	for _, c := range cases {
		if got := PopCount(c.input); got != c.expected {
			t.Errorf("PopCount(%v) != %v", c.input, got)
		}
	}
}

func TestPopCountByLoop(t *testing.T) {
	for _, c := range cases {
		if got := PopCountByShift(c.input); got != c.expected {
			t.Errorf("PopCountByShift(%v) != %v", c.input, got)
		}
	}
}

func TestPopCountByShift(t *testing.T) {
	for _, c := range cases {
		if got := PopCountByShift(c.input); got != c.expected {
			t.Errorf("PopCountByShift(%v) != %v", c.input, got)
		}
	}
}

func TestPopCountByClearing(t *testing.T) {
	for _, c := range cases {
		if got := PopCountByClearing(c.input); got != c.expected {
			t.Errorf("PopCountByClearing(%v) != %v", c.input, got)
		}
	}
}

// -- Benchmarks --

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByLoop(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByShift(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByClearing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByClearing(0x1234567890ABCDEF)
	}
}

// -- Results --

// 2.60GHz Intel Core i5

// $ go test -bench=.

// BenchmarkPopCount-4                     2000000000               0.33 ns/op
// BenchmarkPopCountByLoop-4               100000000               24.7 ns/op
// BenchmarkPopCountByShift-4              20000000                79.2 ns/op
// BenchmarkPopCountByClearing-4           50000000                29.3 ns/op
// PASS
// ok      gopl-ex/ch2/ex2-345     6.475s
