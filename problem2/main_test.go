package main

import (
	"reflect"
	"testing"
)

func TestArea(t *testing.T) {
	type testpair struct {
		input  string
		output int
	}

	var tests = []testpair{
		{"2x3x4", 52},
		{"1x1x10", 42},
	}

	for _, pair := range tests {
		num := Area(pair.input)
		if num != pair.output {
			t.Error(
				"For", pair.input, "Expected", pair.output, "Got", num,
			)
		}
	}
}

func TestVolume(t *testing.T) {
	type testpair struct {
		input  string
		output int
	}

	var tests = []testpair{
		{"2x3x4", 24},
		{"1x1x10", 10},
	}

	for _, pair := range tests {
		num := Volume(pair.input)
		if num != pair.output {
			t.Error(
				"For", pair.input, "Expected", pair.output, "Got", num,
			)
		}
	}
}

func TestMinPerimeter(t *testing.T) {
	type testpair struct {
		input  string
		output int
	}

	var tests = []testpair{
		{"2x3x1", 6},
		{"1x1x10", 4},
	}

	for _, pair := range tests {
		num := MinPerimeter(pair.input)
		if num != pair.output {
			t.Error(
				"For", pair.input, "Expected", pair.output, "Got", num,
			)
		}
	}
}

func TestMinArea(t *testing.T) {
	type testpair struct {
		input  string
		output int
	}

	var tests = []testpair{
		{"2x3x4", 6},
		{"1x1x10", 1},
	}

	for _, pair := range tests {
		num := MinArea(pair.input)
		if num != pair.output {
			t.Error(
				"For", pair.input, "Expected", pair.output, "Got", num,
			)
		}
	}
}

func TestSides(t *testing.T) {
	type testpair struct {
		input  string
		output []int
	}

	var tests = []testpair{
		{"2x3x4", []int{2, 3, 4}},
		{"1x1x10", []int{1, 1, 10}},
	}

	for _, pair := range tests {
		num := getSides(pair.input)
		if !reflect.DeepEqual(num, pair.output) {
			t.Error(
				"For", pair.input, "Expected", pair.output, "Got", num,
			)
		}
	}
}

func TestAreaPerSide(t *testing.T) {
	type testpair struct {
		side1  int
		side2  int
		output int
	}

	var tests = []testpair{
		{2, 3, 6},
		{1, 10, 10},
	}

	for _, pair := range tests {
		num := AreaPerSide(pair.side1, pair.side2)
		if num != pair.output {
			t.Error(
				"For", pair.side1, pair.side2, "Expected", pair.output, "Got", num,
			)
		}
	}
}

func TestPerimeter(t *testing.T) {
	type testpair struct {
		side1  int
		side2  int
		output int
	}

	var tests = []testpair{
		{2, 3, 10},
		{1, 10, 22},
	}

	for _, pair := range tests {
		num := Perimeter(pair.side1, pair.side2)
		if num != pair.output {
			t.Error(
				"For", pair.side1, pair.side2, "Expected", pair.output, "Got", num,
			)
		}
	}
}
