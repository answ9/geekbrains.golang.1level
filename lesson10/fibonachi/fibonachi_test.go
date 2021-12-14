package fibonachi_test

// не стал, к сожалению, изобретать велосипед, повторил что было на лекции

import (
	"fmt"
	"lesson10/fibonachi"
	"testing"
)

var ExpectedList = []uint64{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144}

func TestFibZero(t *testing.T) {
	if res := fibonachi.Fib(0, false); res != 0 {
		t.Errorf("Fib for 0 should be 0, but result: %d", res)
	}
}

func TestFibOne(t *testing.T) {
	if res := fibonachi.Fib(1, false); res != 1 {
		t.Errorf("Fib for 1 should be 1, but result: %d", res)
	}
}

func TestFibMult(t *testing.T) {
	for i, exp := range ExpectedList {
		if res := fibonachi.Fib(uint64(i), false); res != exp {
			t.Errorf("Fib for %d should be %d, but result %d", i, exp, res)
		}
	}
}

func TestFibMultWithCache(t *testing.T) {
	for i, exp := range ExpectedList {
		if res := fibonachi.Fib(uint64(i), false); res != exp {
			t.Errorf("Fib for %d should be %d, but result %d", i, exp, res)
		}
	}
}

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fibonachi.Fib(45, false)
	}
}

func BenchmarkFibWithCache(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fibonachi.Fib(45, true)
	}
}

func ExampleFib() {
	fmt.Println(fibonachi.Fib(20, true))
}
