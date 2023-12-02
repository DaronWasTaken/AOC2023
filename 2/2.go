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
			setColors := strings.Split(set, ",")
			colorsCubeCounts := [3]int{0, 0, 0} // red, green, blue

			for _, colorStr := range setColors {
				switch {
				case strings.Contains(colorStr, "red"):
					extractCubeCountOfColor(colorStr, &colorsCubeCounts[0]) //pass red
					if maxRed < colorsCubeCounts[0] {
						maxRed = colorsCubeCounts[0]
					}
				case strings.Contains(colorStr, "green"):
					extractCubeCountOfColor(colorStr, &colorsCubeCounts[1])
					if maxGreen < colorsCubeCounts[1] {
						maxGreen = colorsCubeCounts[1]
					}
				case strings.Contains(colorStr, "blue"):
					extractCubeCountOfColor(colorStr, &colorsCubeCounts[2])
					if maxBlue < colorsCubeCounts[2] {
						maxBlue = colorsCubeCounts[2]
					}
				}
			}

			isSetPossible := true
			for i := 0; i < len(colorsCubeCounts); i++ {
				if colorsCubeCounts[i] > maxCubeCount[i] {
					isSetPossible = false
					break
				}
			}

			if isSetPossible {
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
func extractCubeCountOfColor(colorStr string, cubeColorCount *int) {
	*cubeColorCount, _ = strconv.Atoi(regexp.MustCompile("[0-9]+").FindString(colorStr))
}
