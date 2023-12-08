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

const FilePath = "input_test2.txt"

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
		//fmt.Println(hand)
		hand.Bid = bid
		hands = append(hands, hand)

		//fmt.Println(cards, "  ===>  ", bid, "  ===>  ", hand)
	}

	sort.Slice(hands, func(i, j int) bool {
		return hands[i].Type > hands[j].Type
	})

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
	//fmt.Println(hands)

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

	//sort.Slice(bytes, func(i, j int) bool {
	//	return bytes[i] < bytes[j]
	//})
	//slices.Reverse(bytes)

	//fmt.Println(bytes)

	cards := bytes2.Runes(bytes)

	//sort.Slice(cards, func(i, j int) bool {
	//	return cards[i] > cards[j]
	//})

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

	valueMap := map[byte]int{
		'A': 13, 'K': 12, 'Q': 11, 'T': 10,
		'9': 9, '8': 8, '7': 7, '6': 6, '5': 5, '4': 4,
		'3': 3, '2': 2, 'J': 1,
	}

	bytes := []byte(hand)
	bytes = bytes2.Map(func(r rune) rune {
		return rune(valueMap[byte(r)])
	}, bytes)

	//sort.Slice(bytes, func(i, j int) bool {
	//	return bytes[i] < bytes[j]
	//})
	//slices.Reverse(bytes)

	//fmt.Println(bytes)

	cards := bytes2.Runes(bytes)

	//sort.Slice(cards, func(i, j int) bool {
	//	return cards[i] > cards[j]
	//})

	joker := false

	totalChar := make(map[rune]int)
	for _, char := range hand {
		totalChar[char] += 1
		if char == 'J' {
			joker = true
			continue
		}
	}

	maxChar := ' '
	maxVal := 0
	count := 0

	if joker {
		first := true
		for k, v := range totalChar {
			count += v
			if first || v > maxVal || (v == maxVal && k > maxChar) {
				maxChar = k
				maxVal = v
				first = false
			}
		}
		if maxVal > 2 {
			if count == 4 {

			}
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
