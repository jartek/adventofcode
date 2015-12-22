package main

import (
	"reflect"
	"testing"
)

func TestIlluminate(t *testing.T) {
	type testpair struct {
		input  string
		output [][]Light
	}

	var lights = make([][]Light, 2)

	for i := 0; i < len(lights); i++ {
		lights[i] = make([]Light, len(lights)+1)

		for j := 0; j < len(lights); j++ {
			lights[i][j].Lit = false
			lights[i][j].Brightness = 0
		}
	}

	var tests = []testpair{
		{"turn on 0,0 through 1,2", [][]Light{{{true, 1}, {true, 1}, {true, 1}}, {{true, 1}, {true, 1}, {true, 1}}}},
		// This command will work on the previous output
		{"turn off 1,0 through 1,1", [][]Light{{{true, 1}, {true, 1}, {true, 1}}, {{false, 0}, {false, 0}, {true, 1}}}},
		// Turn off again just to ensure brightness doesn't slip beyond 0
		{"turn off 1,0 through 1,1", [][]Light{{{true, 1}, {true, 1}, {true, 1}}, {{false, 0}, {false, 0}, {true, 1}}}},
		// So will this
		{"toggle 0,0 through 1,1", [][]Light{{{false, 3}, {false, 3}, {true, 1}}, {{true, 2}, {true, 2}, {true, 1}}}},
	}

	for _, pair := range tests {
		num := Illuminate(lights, pair.input)
		if !reflect.DeepEqual(num, pair.output) {
			t.Error(
				"For", pair.input, "Expected", pair.output, "Got", num,
			)
		}
	}
}
