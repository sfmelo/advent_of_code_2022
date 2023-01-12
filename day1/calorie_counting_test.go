package main

import (
	"testing"
)

func TestCalorieCounting(t *testing.T) {
	type test struct {
		input string
		top   int
		want  int
	}

	example := `1000
2000
3000
	
4000
	
5000
6000
	
7000
8000
9000
	
10000`

	tests := []test{
		{input: example, top: 1, want: 24000},
		{input: example, top: 3, want: 45000},
	}

	for _, tc := range tests {
		got := CalorieCounting(tc.input, tc.top)
		if got != tc.want {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
