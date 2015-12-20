package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func TotalHousesVisitedBySanta(input string) int {
	locations := [][]int{{0, 0}}
	currentLocation := locations[0]

	for i := 0; i < len(input); i++ {
		currentLocation = findNewPosition(currentLocation, string(input[i]))
		locations = appendIfMissing(locations, currentLocation)
	}

	return len(locations)
}

func TotalHousesVisitedBySantaAndRobot(input string) int {
	locations := [][]int{{0, 0}}
	currentLocationOfSanta := locations[0]
	currentLocationOfRobot := locations[0]

	for i := 0; i < len(input); i++ {
		if i%2 == 0 {
			currentLocationOfSanta = findNewPosition(currentLocationOfSanta, string(input[i]))
			locations = appendIfMissing(locations, currentLocationOfSanta)
		} else {
			currentLocationOfRobot = findNewPosition(currentLocationOfRobot, string(input[i]))
			locations = appendIfMissing(locations, currentLocationOfRobot)
		}
	}

	return len(locations)
}

func main() {
	path, _ := filepath.Abs("input")
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		fmt.Println("The total houses visted in the first year is", TotalHousesVisitedBySanta(scanner.Text()))

		fmt.Println("The total houses visted in the second year is", TotalHousesVisitedBySantaAndRobot(scanner.Text()))
	}
}

func findNewPosition(currentPosition []int, direction string) []int {
	var newPosition []int
	if direction == "^" {
		newPosition = []int{currentPosition[0] + 1, currentPosition[1]}
	} else if direction == ">" {
		newPosition = []int{currentPosition[0], currentPosition[1] + 1}
	} else if direction == "v" {
		newPosition = []int{currentPosition[0] - 1, currentPosition[1]}
	} else if direction == "<" {
		newPosition = []int{currentPosition[0], currentPosition[1] - 1}
	}
	return newPosition
}

func appendIfMissing(slice [][]int, pos []int) [][]int {
	for _, loc := range slice {
		if loc[0] == pos[0] && loc[1] == pos[1] {
			return slice
		}
	}
	return append(slice, pos)
}
