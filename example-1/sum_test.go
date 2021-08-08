package main

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	res := Sum(5, 5)

	if res != 10 {
		t.Errorf("Sum(5, 5) = %d; want 10", res)
	}
}

func TestSumVariadic(t *testing.T) {
	sumTest := SumVariadic(1, 3)
	
	if sumTest != 4 {
		t.Errorf("SumVariadic(%d, %d) = %d; want 4", 1, 3, sumTest)
	}

}

func TestSumTableDriven(t *testing.T) {
	var tests = []struct {
		x, y int
		res  int
	}{
		{1, 2, 3},
		{1, 0, 1},
		{2, 3, 5},
		{2, 2, 4},
		{4, 3, 7},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%d,%d", tt.x, tt.y)
		t.Run(testname, func(t *testing.T) {
			result := Sum(tt.x, tt.y)
			if result != tt.res {
				t.Errorf("result is %d, want %d", result, tt.res)
			}
		})
	}
}
