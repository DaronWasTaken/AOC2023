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

type Game struct {
	id                 int
	winningNumbers     []string
	scratchcardNumbers []string
}

func main() {
	file, err := os.Open(FilePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0

	var games []Game

	gameCount := 0

	for scanner.Scan() {

		line := scanner.Text()
		numbersPart := strings.Split(line, ":")[1]
		numbers := strings.Split(numbersPart, "|")

		winningNumbers := regexp.MustCompile(`\d+`).FindAllString(numbers[0], -1)
		scratchcardNumbers := regexp.MustCompile(`\d+`).FindAllString(numbers[1], -1)

		games = append(games, Game{id: gameCount, winningNumbers: winningNumbers, scratchcardNumbers: scratchcardNumbers})

		gameCount++
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

	for i := 0; i < len(games); i++ {
		count := 0
		game := games[i]
		for _, winningNumber := range game.winningNumbers {
			if slices.Contains(game.scratchcardNumbers, winningNumber) {
				count++
			}
		}

		if count > 0 {
			for n := 1; n <= count; n++ {
				if n > len(games) {
					break
				}
				games = append(games, games[game.id+n])
			}
		}
	}

	fmt.Printf("Sum of points: %d\n", sum)
	fmt.Printf("Number of scratchcards: %d\n", len(games))
}
