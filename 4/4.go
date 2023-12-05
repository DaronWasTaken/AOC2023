package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strings"
)

const FilePath = "input.txt"

func main() {
	file, err := os.Open(FilePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {

		line := scanner.Text()
		numbersPart := strings.Split(line, ":")[1]
		numbers := strings.Split(numbersPart, "|")

		winningNumbers := regexp.MustCompile(`\d+`).FindAllString(numbers[0], -1)
		scratchcardNumbers := regexp.MustCompile(`\d+`).FindAllString(numbers[1], -1)

		//fmt.Println(winningNumbers)
		//fmt.Println(scratchcardNumbers)

		count := 0
		points := 0

		for _, winningNumber := range winningNumbers {
			if slices.Contains(scratchcardNumbers, winningNumber) {
				count++
			}
		}

		if count > 0 {
			points = 1
			for i := 1; i < count; i++ {
				points *= 2
			}
		}

		sum += points
	}

	fmt.Printf("Sum of points: %d\n", sum)
}
