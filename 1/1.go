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
	file, err := os.Open(FilePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Map for number words
	//numberWords := map[string]string{
	//	"zero": "0", "one": "1", "two": "2", "three": "3", "four": "4",
	//	"five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9",
	//}

	numberWords := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}

	// Regular expression to match single digits and number words
	digitRegex := regexp.MustCompile(`\d`)

	var sum int

	for scanner.Scan() {
		line := scanner.Text()

		for number, word := range numberWords {

			line = strings.ReplaceAll(line, word, word+strconv.Itoa(number)+word)
		}

		// Extract all number words and digits
		allMatches := digitRegex.FindAllStringSubmatch(line, -1)

		// Convert matches to a slice of strings
		matches := make([]string, len(allMatches))
		for i, match := range allMatches {
			matches[i] = match[0]
		}

		if len(matches) > 0 {
			firstItem := matches[0]
			lastItem := matches[len(matches)-1]

			// Combine values
			calibrationValue, _ := strconv.Atoi(firstItem + lastItem)
			sum += calibrationValue
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading from file:", err)
	}

	fmt.Println("Total Sum:", sum)
}
