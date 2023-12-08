package main

import (
	"bufio"
	bytes2 "bytes"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const FilePath = "input.txt"

type HandType int

const (
	HighCard HandType = iota + 1
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	Poker
)

func main() {
	file, err := os.Open(FilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum1 := 0

	var hands []Hand
	for scanner.Scan() {
		line := scanner.Text()
		splits := strings.Split(line, " ")
		cards := splits[0]
		bid, err := strconv.Atoi(splits[1])
		if err != nil {
			log.Fatal(err)
		}
		hand := determineHand2(cards)
		hand.Bid = bid
		hands = append(hands, hand)
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].Type != hands[j].Type {
			return hands[i].Type < hands[j].Type
		}
		for n := 0; n < len(hands[i].CardValues); n++ {
			if hands[i].CardValues[n] != hands[j].CardValues[n] {
				return hands[i].CardValues[n] < hands[j].CardValues[n]
			}
		}
		return true
	})

	for i, hand := range hands {
		sum1 += hand.Bid * (i + 1)
	}

	fmt.Printf("Part one sum: %d\n", sum1)
}

type Hand struct {
	CardValues []rune
	Type       HandType
	Bid        int
}

func determineHand(hand string) Hand {

	valueMap := map[byte]int{
		'A': 14, 'K': 13, 'Q': 12, 'J': 11, 'T': 10,
		'9': 9, '8': 8, '7': 7, '6': 6, '5': 5, '4': 4,
		'3': 3, '2': 2,
	}

	bytes := []byte(hand)
	bytes = bytes2.Map(func(r rune) rune {
		return rune(valueMap[byte(r)])
	}, bytes)

	cards := bytes2.Runes(bytes)

	totalChar := make(map[rune]int)
	for _, char := range hand {
		totalChar[char] += 1
	}

	if len(totalChar) == 1 {
		return Hand{Type: Poker, CardValues: cards}
	}

	if len(totalChar) == 2 {
		for _, val := range totalChar {
			if val == 2 || val == 3 {
				return Hand{Type: FullHouse, CardValues: cards}
			} else {
				return Hand{Type: FourOfAKind, CardValues: cards}
			}
		}
	}

	if len(totalChar) == 3 {
		for _, v := range totalChar {
			if v == 2 {
				return Hand{Type: TwoPair, CardValues: cards}
			}
		}
		return Hand{Type: ThreeOfAKind, CardValues: cards}
	}

	if len(totalChar) == 4 {
		return Hand{Type: OnePair, CardValues: cards}
	}

	return Hand{Type: HighCard, CardValues: cards}
}

func determineHand2(hand string) Hand {

	valueMap := map[rune]int{
		'A': 13, 'K': 12, 'Q': 11, 'T': 10,
		'9': 9, '8': 8, '7': 7, '6': 6, '5': 5, '4': 4,
		'3': 3, '2': 2, 'J': 1,
	}

	bytes := []byte(hand)
	bytes = bytes2.Map(func(r rune) rune {
		return rune(valueMap[r])
	}, bytes)

	cards := bytes2.Runes(bytes)
	joker := false

	totalChar := make(map[rune]int)
	for _, char := range hand {
		if char == 'J' {
			joker = true
			continue
		}
		totalChar[char] += 1
	}

	maxChar := 'X'
	maxVal := 0
	count := 0

	if joker {
		for k, v := range totalChar {
			count += v
			if v > maxVal {
				maxChar = k
				maxVal = v
			} else if v == maxVal {
				if valueMap[k] > valueMap[maxChar] {
					maxChar = k
				}
			}
		}
		if len(totalChar) < 4 {
			totalChar[maxChar] += len(hand) - count
		} else {
			totalChar[maxChar] += 1
		}
	}

	if len(totalChar) == 1 {
		return Hand{Type: Poker, CardValues: cards}
	}

	if len(totalChar) == 2 {
		for _, val := range totalChar {
			if val == 2 || val == 3 {
				return Hand{Type: FullHouse, CardValues: cards}
			} else {
				return Hand{Type: FourOfAKind, CardValues: cards}
			}
		}
	}

	if len(totalChar) == 3 {
		for _, v := range totalChar {
			if v == 2 {
				return Hand{Type: TwoPair, CardValues: cards}
			}
		}
		return Hand{Type: ThreeOfAKind, CardValues: cards}
	}

	if len(totalChar) == 4 {
		return Hand{Type: OnePair, CardValues: cards}
	}

	return Hand{Type: HighCard, CardValues: cards}
}
