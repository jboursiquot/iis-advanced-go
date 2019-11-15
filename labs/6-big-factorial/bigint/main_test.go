package main

import (
	"math/big"
	"testing"
)

func TestFactorial(t *testing.T) {
	tests := map[string]struct {
		num  int64
		want *big.Int
	}{
		"0":  {num: 0, want: big.NewInt(1)},
		"1":  {num: 1, want: big.NewInt(1)},
		"2":  {num: 2, want: big.NewInt(2)},
		"3":  {num: 3, want: big.NewInt(6)},
		"10": {num: 10, want: big.NewInt(3628800)},
		"20": {num: 20, want: big.NewInt(2432902008176640000)},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if got := factorial(tc.num); tc.want.Cmp(got) != 0 {
				t.Errorf("wanted %v got %v", tc.want, got)
			}
		})
	}
}

func BenchmarkFactorial(b *testing.B) {
	for n := 0; n < b.N; n++ {
		factorial(int64(10))
	}
}
