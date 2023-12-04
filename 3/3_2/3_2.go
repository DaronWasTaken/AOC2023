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

type Number struct {
	x     int
	y     int
	value int
}

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

	var numbers []Number
	var sum int
	for i := 0; i < len(charArray); i++ {
		number := ""
		for j := 0; j < len(charArray[i]); j++ {
			if unicode.IsDigit(charArray[i][j]) {
				number += string(charArray[i][j])
			} else {
				x1, y1 := checkAdjacent(&charArray, i, j-len(number), j-1)

				if x1 > 0 {
					value, err := strconv.Atoi(number)
					if err != nil {
						fmt.Printf("Number to str fail: %s\n", number)
					}
					numbers = append(numbers, Number{x1, y1, value})
				}
				number = ""
			}
		}

		// Check for a number at the end of the line
		if number != "" {
			x1, y1 := checkAdjacent(&charArray, i, len(charArray[i])-len(number), len(charArray[i])-1)
			if x1 > 0 {
				value, err := strconv.Atoi(number)
				if err != nil {
					fmt.Printf("Number to str fail: %s\n", number)
				}
				numbers = append(numbers, Number{x1, y1, value})
			}
			number = ""
		}
	}

	var filteredNumbers []Number

	for _, number := range numbers {
		if number.value != 0 {
			filteredNumbers = append(filteredNumbers, number)
		}
	}

	//for _, filteredNumber := range filteredNumbers {
	//	fmt.Println(filteredNumber)
	//}

	groups := make(map[string][]Number)
	for _, s := range filteredNumbers {
		key := fmt.Sprintf("%d,%d", s.x, s.y) // Create a unique key for each coordinate pair
		groups[key] = append(groups[key], s)
	}

	for _, numberList := range groups {
		if len(numberList) == 2 {
			fmt.Println(numberList)
			sum += numberList[0].value * numberList[1].value
		}
	}

	fmt.Println("Sum:", sum)
}

func checkAdjacent(array *[][]rune, i, start, end int) (int, int) {

	x, y := checkLeft(array, i, start)
	if x > 0 {
		return x, y
	}

	x, y = checkRight(array, i, end)
	if x > 0 {
		return x, y
	}

	x, y = checkBottom(array, i, start, end)
	if x > 0 {
		return x, y
	}

	x, y = checkTop(array, i, start, end)
	if x > 0 {
		return x, y
	}

	x, y = checkDiagonalLeft(array, i, start)
	if x > 0 {
		return x, y
	}

	x, y = checkDiagonalRight(array, i, end)
	if x > 0 {
		return x, y
	}

	return -1, -1
}

func checkTop(array *[][]rune, i, start, end int) (int, int) {
	if i > 0 {
		for k := start; k <= end; k++ {
			if (*array)[i-1][k] == '*' {
				return i - 1, k
			}
		}
	}
	return -1, -1
}

func checkBottom(array *[][]rune, i, start, end int) (int, int) {
	if i < len(*array)-1 {
		for k := start; k <= end; k++ {
			if (*array)[i+1][k] == '*' {
				return i + 1, k
			}
		}
	}
	return -1, -1
}

func checkLeft(array *[][]rune, i, start int) (int, int) {
	if start > 0 && (*array)[i][start-1] == '*' {
		return i, start - 1
	}
	return -1, -1
}

func checkRight(array *[][]rune, i, end int) (int, int) {
	if end < len((*array)[i])-1 && (*array)[i][end+1] == '*' {
		return i, end + 1
	}
	return -1, -1
}

func checkDiagonalLeft(array *[][]rune, i, start int) (int, int) {
	if i > 0 && start > 0 && (*array)[i-1][start-1] == '*' {
		return i - 1, start - 1
	}
	if i < len(*array)-1 && start > 0 && (*array)[i+1][start-1] == '*' {
		return i + 1, start - 1
	}
	return -1, -1
}

func checkDiagonalRight(array *[][]rune, i, end int) (int, int) {
	if i > 0 && end < len((*array)[i])-1 && (*array)[i-1][end+1] == '*' {
		return i - 1, end + 1
	}
	if i < len(*array)-1 && end < len((*array)[i])-1 && (*array)[i+1][end+1] == '*' {
		return i + 1, end + 1
	}
	return -1, -1
}
