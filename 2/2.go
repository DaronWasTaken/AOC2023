package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	FilePath = "input.txt"
)

func main() {

	maxCubeCount := [3]int{12, 13, 14}
	file, err := os.Open(FilePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	gameId := 0
	sum := 0
	powerOfSets := 0

	for scanner.Scan() {
		gameId++
		line := scanner.Text()
		gameParts := strings.Split(line, ":")
		sets := strings.Split(gameParts[1], ";")

		possibleSetCount := 0

		maxRed := 0
		maxGreen := 0
		maxBlue := 0

		for _, set := range sets {
			colors := strings.Split(set, ",")
			cubeCountPerColor := [3]int{0, 0, 0} // red, green, blue

			for _, color := range colors {
				switch {
				case strings.Contains(color, "red"):
					extractCubeCountOfColor(color, &cubeCountPerColor, 0)
					if maxRed < cubeCountPerColor[0] {
						maxRed = cubeCountPerColor[0]
					}
				case strings.Contains(color, "green"):
					extractCubeCountOfColor(color, &cubeCountPerColor, 1)
					if maxGreen < cubeCountPerColor[1] {
						maxGreen = cubeCountPerColor[1]
					}
				case strings.Contains(color, "blue"):
					extractCubeCountOfColor(color, &cubeCountPerColor, 2)
					if maxBlue < cubeCountPerColor[2] {
						maxBlue = cubeCountPerColor[2]
					}
				}
			}

			isBelowMax := true
			for i := 0; i < len(cubeCountPerColor); i++ {
				if cubeCountPerColor[i] > maxCubeCount[i] {
					isBelowMax = false
					break
				}
			}

			if isBelowMax {
				possibleSetCount++
			}

		}

		powerOfSets += maxRed * maxGreen * maxBlue

		if possibleSetCount == len(sets) {
			sum += gameId
		}

	}

	fmt.Printf("Possible game id sum: %d\n", sum)
	fmt.Printf("Power of sets sum: %d\n", powerOfSets)
}

// 0 - red, 1 - green, 2 - blue
func extractCubeCountOfColor(colorStr string, cubeCountPerColor *[3]int, colorId int) {
	cubeCountPerColor[colorId], _ = strconv.Atoi(regexp.MustCompile("[0-9]+").FindString(colorStr))
}
