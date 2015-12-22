package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ContainsThreeVowels(input string) bool {
	vowelCount := 0

	for _, value := range input {
		switch value {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			vowelCount++
		}
	}

	return vowelCount >= 3
}

func ContainsSameLetterTwice(input string) bool {
	for i := 0; i < len(input)-1; i++ {
		if input[i] == input[i+1] {
			return true
		}
	}

	return false
}

func ContainsInvalidSequences(input string) bool {
	invalidStrings := []string{"ab", "cd", "pq", "xy"}
	for i := 0; i < len(invalidStrings); i++ {
		if strings.Index(input, invalidStrings[i]) != -1 {
			return true
		}
	}
	return false
}

func Nice(input string) bool {
	return !ContainsInvalidSequences(input) && ContainsThreeVowels(input) && ContainsSameLetterTwice(input)
}

func ContainsRepeatedCharacterWithOneCharacterBetween(input string) bool {
	for i := 0; i < len(input)-2; i++ {
		if input[i] == input[i+2] {
			return true
		}
	}

	return false
}

func ContainsPairWithoutOverlap(input string) bool {
	for i := 0; i < len(input)-1; i++ {
		substring := input[i : i+2]
		firstIndex := strings.Index(input, substring)
		lastIndex := strings.LastIndex(input, substring)
		if lastIndex != firstIndex && lastIndex != firstIndex+1 {
			return true
		}
	}

	return false
}

func UpdatedNice(input string) bool {
	return ContainsPairWithoutOverlap(input) && ContainsRepeatedCharacterWithOneCharacterBetween(input)
}

func main() {
	path, _ := filepath.Abs("input")
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	totalNiceStrings := 0
	totalUpdatedNiceStrings := 0

	for scanner.Scan() {
		if Nice(scanner.Text()) {
			totalNiceStrings += 1
		}

		if UpdatedNice(scanner.Text()) {
			totalUpdatedNiceStrings += 1
		}
	}

	fmt.Println("The total number of nice strings are", totalNiceStrings)
	fmt.Println("The total number of updated nice strings are", totalUpdatedNiceStrings)
}
