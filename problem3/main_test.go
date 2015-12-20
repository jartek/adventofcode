package main

import "testing"

func TestTotalHousesVistedBySanta(t *testing.T) {
	type testpair struct {
		input  string
		output int
	}

	var tests = []testpair{
		{">", 2},
		{"^>v<", 4},
		{"^v^v^v^v^v", 2},
	}

	for _, pair := range tests {
		num := TotalHousesVisitedBySanta(pair.input)
		if num != pair.output {
			t.Error(
				"For", pair.input, "Expected", pair.output, "Got", num,
			)
		}
	}
}

func TestTotalHousesVistedBySantaAndRobot(t *testing.T) {
	type testpair struct {
		input  string
		output int
	}

	var tests = []testpair{
		{"^v", 3},
		{"^>v<", 3},
		{"^v^v^v^v^v", 11},
	}

	for _, pair := range tests {
		num := TotalHousesVisitedBySantaAndRobot(pair.input)
		if num != pair.output {
			t.Error(
				"For", pair.input, "Expected", pair.output, "Got", num,
			)
		}
	}
}
