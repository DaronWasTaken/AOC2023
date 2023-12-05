package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

const FilePath = "input.txt"

type Map struct {
	name     string
	mappings []Mapping
}

type Mapping struct {
	destination int
	source      int
	length      int
}

func main() {
	file, err := os.Open(FilePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	var seeds []int
	var maps []Map

	scanner := bufio.NewScanner(file)

	willBeMap := false
	var theMap *Map

	for scanner.Scan() {

		line := scanner.Text()

		if strings.Contains(line, "seeds:") {
			parts := strings.Split(line, ": ")
			seedsString := strings.Split(parts[1], " ")
			for _, str := range seedsString {
				seed, err := strconv.Atoi(str)
				if err != nil {
					fmt.Println(err)
				}
				seeds = append(seeds, seed)
			}
		}

		if line == "" && willBeMap {
			willBeMap = false
			maps = append(maps, *theMap)
			continue
		}

		if strings.Contains(line, "map:") {
			willBeMap = true
			theMap = &Map{name: line, mappings: make([]Mapping, 0)}
			continue
		}

		if willBeMap {
			mappingString := strings.Split(line, " ")
			rangeStart, _ := strconv.Atoi(mappingString[0])
			sourceRangeStart, _ := strconv.Atoi(mappingString[1])
			rangeLength, _ := strconv.Atoi(mappingString[2])
			theMap.mappings = append(theMap.mappings, Mapping{
				rangeStart,
				sourceRangeStart,
				rangeLength},
			)
		}

	}

	if willBeMap {
		maps = append(maps, *theMap)
	}

	partOne(&seeds, &maps)
	partTwo(&seeds, &maps)
}

func partOne(seeds *[]int, maps *[]Map) {
	var results []int
	for _, seed := range *seeds {
		number := seed
		for _, currentMap := range *maps {
			for _, mapping := range currentMap.mappings {
				minN := mapping.source
				maxN := mapping.source + mapping.length
				if number >= minN && number < maxN {
					number = number - minN + mapping.destination
					break
				}
			}
		}
		results = append(results, number)
	}

	minN := slices.Min(results)
	fmt.Println(minN)
}

func partTwo(seeds *[]int, maps *[]Map) {
	var results []int

	for n := 0; n < len(*seeds); n += 2 {
		start := (*seeds)[n]
		end := (*seeds)[n] + (*seeds)[n+1]
		for i := start; i < end; i++ {
			number := i
			for _, currentMap := range *maps {
				for _, mapping := range currentMap.mappings {
					minN := mapping.source
					maxN := mapping.source + mapping.length
					if number >= minN && number < maxN {
						number = number - minN + mapping.destination
						break
					}
				}
			}
			results = append(results, number)
		}
	}

	minN := slices.Min(results)
	fmt.Println(minN)
}
