package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

type Light struct {
	Lit        bool
	Brightness int
}

func turnOn(lights [][]Light, startX int, startY int, finishX int, finishY int) [][]Light {
	for i := startX; i <= finishX; i++ {
		for j := startY; j <= finishY; j++ {
			lights[i][j].Lit = true
			lights[i][j].Brightness = lights[i][j].Brightness + 1
		}
	}

	return lights
}

func turnOff(lights [][]Light, startX int, startY int, finishX int, finishY int) [][]Light {
	for i := startX; i <= finishX; i++ {
		for j := startY; j <= finishY; j++ {
			lights[i][j].Lit = false
			if lights[i][j].Brightness > 0 {
				lights[i][j].Brightness = lights[i][j].Brightness - 1
			}
		}
	}

	return lights
}

func toggle(lights [][]Light, startX int, startY int, finishX int, finishY int) [][]Light {
	for i := startX; i <= finishX; i++ {
		for j := startY; j <= finishY; j++ {
			lights[i][j].Lit = !lights[i][j].Lit
			lights[i][j].Brightness = lights[i][j].Brightness + 2
		}
	}

	return lights
}

func Illuminate(lights [][]Light, input string) [][]Light {
	re := regexp.MustCompile(`(?P<instruction>toggle|turn off|turn on)\s*(?P<startx>\d*),(?P<starty>\d*)\s*through\s*(?P<finishx>\d+),(?P<finishy>\d+)`)

	parsedString := re.FindAllStringSubmatch(input, -1)[0]
	instruction := parsedString[1]
	startX, _ := strconv.Atoi(parsedString[2])
	startY, _ := strconv.Atoi(parsedString[3])
	finishX, _ := strconv.Atoi(parsedString[4])
	finishY, _ := strconv.Atoi(parsedString[5])

	switch instruction {
	case "toggle":
		lights = toggle(lights, startX, startY, finishX, finishY)
		break

	case "turn on":
		lights = turnOn(lights, startX, startY, finishX, finishY)
		break

	case "turn off":
		lights = turnOff(lights, startX, startY, finishX, finishY)
		break
	}

	return lights
}

func main() {
	path, _ := filepath.Abs("input")
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	lights := make([][]Light, 1000)

	for i := 0; i < len(lights); i++ {
		lights[i] = make([]Light, len(lights))

		for j := 0; j < len(lights); j++ {
			lights[i][j].Lit = false
			lights[i][j].Brightness = 0
		}
	}

	for scanner.Scan() {
		lights = Illuminate(lights, scanner.Text())
	}

	totalLightsOn := 0
	totalBrightness := 0

	for i := 0; i < len(lights); i++ {
		for j := 0; j < len(lights); j++ {
			if lights[i][j].Lit {
				totalLightsOn = totalLightsOn + 1
			}
			totalBrightness += lights[i][j].Brightness
		}
	}

	fmt.Println("Number of lights on", totalLightsOn)
	fmt.Println("Total Brightness", totalBrightness)
}
