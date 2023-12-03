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
	// Read the file into a 2D slice
	grid := readGrid(FilePath)

	// Find gears and calculate their gear ratios
	var sumGearRatio int
	for y, row := range grid {
		for x, char := range row {
			if char == '*' {
				gearRatio := calculateGearRatio(grid, x, y)
				sumGearRatio += gearRatio
			}
		}
	}

	fmt.Println("Sum of gear ratios:", sumGearRatio)
}

// Reads the file and returns a 2D slice of runes
func readGrid(filePath string) [][]rune {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}
	return grid
}

// Calculates the gear ratio for a gear at a given position
func calculateGearRatio(grid [][]rune, x, y int) int {
	var numbers []int
	directions := [][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} // Left, Right, Up, Down

	for _, dir := range directions {
		nx, ny := x+dir[1], y+dir[0]
		if nx >= 0 && nx < len(grid[0]) && ny >= 0 && ny < len(grid) && unicode.IsDigit(grid[ny][nx]) {
			number := extractNumber(grid, nx, ny, dir)
			numbers = append(numbers, number)
		}
	}

	if len(numbers) == 2 {
		// Calculate product of the first two numbers found
		return numbers[0] * numbers[1]
	}
	return 0
}

func extractNumber(grid [][]rune, x, y int, dir [2]int) int {
	numberStr := ""
	dx, dy := dir[1], dir[0]

	for x >= 0 && x < len(grid[0]) && y >= 0 && y < len(grid) && unicode.IsDigit(grid[y][x]) {
		numberStr += string(grid[y][x])
		x += dx
		y += dy
	}

	number, _ := strconv.Atoi(numberStr)
	return number
}
