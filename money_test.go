package main

import (
	"testing"
)

// TestFindSolution tests the FindSolution function for all possible combinations of 10-yuan, 5-yuan, and 2-yuan bills (a, b, c from 0 to 99).
// It checks both successful and failing cases:
//   - For success: verifies that the returned combination pays the exact amount.
//   - For failure: verifies that the function returns an error and all returned bill counts are zero.
//
// The test covers not only the exact (a, b, c) combination, but also nearby combinations (a+1, b+1, c+1, etc.) and various insufficient cases.
func TestFindSolution(t *testing.T) {
	checkSuccess := func(a, b, c, amount uint) {
		x, y, z, err := FindSolution(a, b, c, amount)
		if err != nil {
			t.Fatalf("FindSolution(%d,%d,%d,%d) should succeed, but got error: %v; a=%d, b=%d, c=%d, amount=%d", a, b, c, amount, err, a, b, c, amount)
		} else {
			got := x*10 + y*5 + z*2
			if got != amount {
				t.Fatalf("FindSolution(%d,%d,%d,%d) got total %d, want %d; a=%d, b=%d, c=%d, amount=%d", a, b, c, amount, got, amount, a, b, c, amount)
			}
		}
	}

	checkFail := func(a, b, c, amount uint) {
		x, y, z, err := FindSolution(a, b, c, amount)
		if err == nil {
			t.Fatalf("FindSolution(%d,%d,%d,%d) should fail, but got no error; a=%d, b=%d, c=%d, amount=%d", a, b, c, amount, a, b, c, amount)
		}
		if x != 0 || y != 0 || z != 0 {
			t.Fatalf("FindSolution(%d,%d,%d,%d) should fail, but got x=%d, y=%d, z=%d; a=%d, b=%d, c=%d, amount=%d", a, b, c, amount, x, y, z, a, b, c, amount)
		}
	}

	for a := range uint(100) {
		for b := range uint(100) {
			for c := range uint(100) {
				amount := a*10 + b*5 + c*2

				checkSuccess(a, b, c, amount)
				checkSuccess(a, b, c+1, amount)
				checkSuccess(a, b+1, c, amount)
				checkSuccess(a, b+1, c+1, amount)
				checkSuccess(a+1, b, c, amount)
				checkSuccess(a+1, b, c+1, amount)
				checkSuccess(a+1, b+1, c, amount)
				checkSuccess(a+1, b+1, c+1, amount)

				if a > 0 {
					checkFail(a-1, b, c, amount)
					if b > 0 {
						checkFail(a-1, b-1, c, amount)
					}
					if c > 0 {
						checkFail(a-1, b, c-1, amount)
					}
				}
				if b > 0 {
					checkFail(a, b-1, c, amount)
					if a > 0 {
						checkFail(a-1, b-1, c, amount)
					}
					if c > 0 {
						checkFail(a, b-1, c-1, amount)
					}
				}
				if c > 0 {
					checkFail(a, b, c-1, amount)
					if a > 0 {
						checkFail(a-1, b, c-1, amount)
					}
					if b > 0 {
						checkFail(a, b-1, c-1, amount)
					}
				}
			}
		}
	}
}
