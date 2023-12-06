package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const FilePath = "input2.txt"

func main() {

	data, err := os.ReadFile(FilePath)
	if err != nil {
		// Handle the error here
		fmt.Println("Error reading file:", err)
		return
	}

	lines := strings.Split(string(data), "\n")

	fmt.Println(lines[0])
	fmt.Println(lines[1])

	regex := regexp.MustCompile(`\d+`)

	timesStr := regex.FindAllString(lines[0], -1)
	distancesStr := regex.FindAllString(lines[1], -1)

	result := 1

	for i := 0; i < len(timesStr); i++ {
		time, _ := strconv.Atoi(timesStr[i])
		distance, _ := strconv.Atoi(distancesStr[i])
		result *= determineFasterMethodsCount(time, distance)
	}

	fmt.Printf("Result: %d\n", result)
}

func determineFasterMethodsCount(time int, distance int) int {
	count := 0
	for i := 1; i <= time; i++ {
		if distanceForTimeHeld(time, i) > distance {
			count++
		}
	}
	return count
}

func distanceForTimeHeld(time int, timeHeld int) int {
	return timeHeld * (time - timeHeld)
}
