package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func AreaPerSide(side1 int, side2 int) int {
	return side1 * side2
}

func MinArea(input string) int {
	sides := getSides(input)
	min := -1
	area := 0

	for i := 0; i < len(sides)-1; i++ {
		for j := i + 1; j < len(sides); j++ {
			area = AreaPerSide(sides[i], sides[j])
			if area < min || min < 0 {
				min = area
			}
		}
	}

	return min
}

func getSides(input string) []int {
	sides := make([]int, 3)
	stringSides := strings.Split(input, "x")

	for i := 0; i < 3; i++ {
		sides[i], _ = strconv.Atoi(stringSides[i])
	}

	return sides
}

func Area(input string) int {
	sides := getSides(input)
	return 2*AreaPerSide(sides[0], sides[1]) + 2*AreaPerSide(sides[1], sides[2]) + 2*AreaPerSide(sides[0], sides[2])
}

func MinPerimeter(input string) int {
	sides := getSides(input)
	min := -1
	perimeter := 0

	for i := 0; i < len(sides)-1; i++ {
		for j := i + 1; j < len(sides); j++ {
			perimeter = Perimeter(sides[i], sides[j])
			if perimeter < min || min < 0 {
				min = perimeter
			}
		}
	}

	return min
}

func Perimeter(side1 int, side2 int) int {
	return 2*side1 + 2*side2
}

func Volume(input string) int {
	sides := getSides(input)
	return sides[0] * sides[1] * sides[2]
}

func main() {
	path, _ := filepath.Abs("input")
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	paper := 0
	ribbon := 0

	for scanner.Scan() {
		paper = paper + Area(scanner.Text()) + MinArea(scanner.Text())
		ribbon = ribbon + MinPerimeter(scanner.Text()) + Volume(scanner.Text())
	}

	fmt.Println("Paper needed: ", paper)
	fmt.Println("Ribbon needed: ", ribbon)
}
