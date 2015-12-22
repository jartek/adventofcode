package main

import "testing"

func TestMining(t *testing.T) {
	type testpair struct {
		input  string
		output int
	}

	var tests = []testpair{
		{"abcdef", 609043},
		{"pqrstuv", 1048970},
	}

	for _, pair := range tests {
		num := Mine(pair.input, "^00000")
		if num != pair.output {
			t.Error(
				"For", pair.input, "Expected", pair.output, "Got", num,
			)
		}
	}
}
