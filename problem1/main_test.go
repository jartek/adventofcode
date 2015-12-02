package main

import "testing"

func TestTotalFloors(t *testing.T) {
	type testpair struct {
		input  string
		output int
	}

	var tests = []testpair{
		{"()", 0},
		{"(()))", -1},
		{"))(((((", 3},
	}

	for _, pair := range tests {
		num := TotalFloors(pair.input)
		if num != pair.output {
			t.Error(
				"For", pair.input, "Expected", pair.output, "Got", num,
			)
		}
	}
}

func TestFirstBasementIndex(t *testing.T) {
	type testpair struct {
		input  string
		output int
	}

	var tests = []testpair{
		{"()())", 5},
		{")", 1},
	}

	for _, pair := range tests {
		num := FirstBasementIndex(pair.input)
		if num != pair.output {
			t.Error(
				"For", pair.input, "Expected", pair.output, "Got", num,
			)
		}
	}
}
