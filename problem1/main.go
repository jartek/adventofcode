package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func TotalFloors(input string) int {
	ctr := 0
	for i := 0; i < len(input); i++ {
		if input[i] == '(' {
			ctr = ctr + 1
		} else {
			ctr = ctr - 1
		}
	}
	return ctr
}

func FirstBasementIndex(input string) int {
	ctr := 0
	for i := 0; i < len(input); i++ {
		if input[i] == '(' {
			ctr = ctr + 1
		} else {
			ctr = ctr - 1
		}

		if ctr < 0 {
			return i + 1
		}
	}
	return -1
}

func main() {
	path, _ := filepath.Abs("input")
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		fmt.Println("Total Floors: ", TotalFloors(scanner.Text()))
		fmt.Println("First Basement Index: ", FirstBasementIndex(scanner.Text()))
	}
}
