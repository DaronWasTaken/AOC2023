package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

const FilePath = "../input.txt"

func main() {
	// Open the file
	file, err := os.Open(FilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a 2D slice (2D char array)
	var charArray [][]rune

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var lineSlice []rune
		for _, ch := range scanner.Text() {
			lineSlice = append(lineSlice, ch)
		}
		charArray = append(charArray, lineSlice)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var sum int
	for i := 0; i < len(charArray); i++ {
		number := ""
		for j := 0; j < len(charArray[i]); j++ {
			if unicode.IsDigit(charArray[i][j]) {
				number += string(charArray[i][j])
			} else {
				if number != "" && checkAdjacent(&charArray, i, j-len(number), j-1) {
					num, _ := strconv.Atoi(number)
					sum += num
				}
				number = ""
			}
		}
		// Check for a number at the end of the line
		if number != "" && checkAdjacent(&charArray, i, len(charArray[i])-len(number), len(charArray[i])-1) {
			num, _ := strconv.Atoi(number)
			sum += num
		}
	}

	fmt.Println("Sum:", sum)
}

func checkAdjacent(array *[][]rune, i, start, end int) bool {
	return checkLeft(array, i, start) || checkRight(array, i, end) ||
		checkTop(array, i, start, end) || checkBottom(array, i, start, end) ||
		checkDiagonalLeft(array, i, start) || checkDiagonalRight(array, i, end)
}

func checkTop(array *[][]rune, i, start, end int) bool {
	if i > 0 {
		for k := start; k <= end; k++ {
			if (*array)[i-1][k] != '.' {
				return true
			}
		}
	}
	return false
}

func checkBottom(array *[][]rune, i, start, end int) bool {
	if i < len(*array)-1 {
		for k := start; k <= end; k++ {
			if (*array)[i+1][k] != '.' {
				return true
			}
		}
	}
	return false
}

func checkLeft(array *[][]rune, i, start int) bool {
	if start > 0 && (*array)[i][start-1] != '.' {
		return true
	}
	return false
}

func checkRight(array *[][]rune, i, end int) bool {
	if end < len((*array)[i])-1 && (*array)[i][end+1] != '.' {
		return true
	}
	return false
}

func checkDiagonalLeft(array *[][]rune, i, start int) bool {
	if i > 0 && start > 0 && (*array)[i-1][start-1] != '.' {
		return true
	}
	if i < len(*array)-1 && start > 0 && (*array)[i+1][start-1] != '.' {
		return true
	}
	return false
}

func checkDiagonalRight(array *[][]rune, i, end int) bool {
	if i > 0 && end < len((*array)[i])-1 && (*array)[i-1][end+1] != '.' {
		return true
	}
	if i < len(*array)-1 && end < len((*array)[i])-1 && (*array)[i+1][end+1] != '.' {
		return true
	}
	return false
}
