package main

import "testing"

func TestContainsThreeVowels(t *testing.T) {
	type testpair struct {
		input  string
		output bool
	}

	var tests = []testpair{
		{"aei", true},
		{"xazegov", true},
		{"aeiouaeiouaeiou", true},
		{"abcdef", false},
	}

	for _, pair := range tests {
		num := ContainsThreeVowels(pair.input)
		if num != pair.output {
			t.Error(
				"For", pair.input, "Expected", pair.output, "Got", num,
			)
		}
	}
}

func TestContainsSameLetterTwice(t *testing.T) {
	type testpair struct {
		input  string
		output bool
	}

	var tests = []testpair{
		{"xx", true},
		{"abcdde", true},
		{"aabbccdd", true},
		{"abcdef", false},
	}

	for _, pair := range tests {
		num := ContainsSameLetterTwice(pair.input)
		if num != pair.output {
			t.Error(
				"For", pair.input, "Expected", pair.output, "Got", num,
			)
		}
	}
}

func TestContainsInvalidSequences(t *testing.T) {
	type testpair struct {
		input  string
		output bool
	}

	var tests = []testpair{
		{"abcdef", true},
		{"mnop", false},
		{"mxyp", true},
		{"cdaab", true},
	}

	for _, pair := range tests {
		num := ContainsInvalidSequences(pair.input)
		if num != pair.output {
			t.Error(
				"For", pair.input, "Expected", pair.output, "Got", num,
			)
		}
	}
}

func TestNiceStrings(t *testing.T) {
	type testpair struct {
		input  string
		output bool
	}

	var tests = []testpair{
		{"ugknbfddgicrmopn", true},
		{"aaa", true},
		{"jchzalrnumimnmhp", false},
		{"haegwjzuvuyypxyu", false},
		{"dvszwmarrgswjxmb", false},
	}

	for _, pair := range tests {
		num := Nice(pair.input)
		if num != pair.output {
			t.Error(
				"For", pair.input, "Expected", pair.output, "Got", num,
			)
		}
	}
}
