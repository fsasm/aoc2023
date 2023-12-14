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

type Hand struct {
	cards    [5]int
	bid      int
	handType HandType
}

type HandType int

const (
	// from lowest to highest
	HighCard HandType = iota + 1
	OnePair
	TwoPair
	ThreeKind
	FullHouse
	FourKind
	FiveKind
)

const cardOrder = "23456789TJQKA"

func CardToOrder(card rune) int {
	return strings.IndexRune(cardOrder, card)
}

const cardOrderJoker = "J23456789TQKA"

func CardToOrderJoker(card rune) int {
	return strings.IndexRune(cardOrderJoker, card)
}

func cardHistogram(cards [5]int) (result map[int]int) {
	result = make(map[int]int)
	for _, c := range cards {
		result[c]++
	}
	return
}

func classifyHand(cards [5]int) HandType {
	histogram := cardHistogram(cards)

	values := make([]int, 0)
	for _, v := range histogram {
		values = append(values, v)
	}
	slices.Sort(values)

	if len(values) == 1 {
		return FiveKind
	}
	if len(values) == 2 {
		if values[0] == 1 && values[1] == 4 {
			return FourKind
		}
		if values[0] == 2 && values[1] == 3 {
			return FullHouse
		}

		log.Fatal("Invalid combination of cards")
	}
	if len(values) == 3 {
		if values[0] == 1 && values[1] == 1 && values[2] == 3 {
			return ThreeKind
		}
		if values[0] == 1 && values[1] == 2 && values[2] == 2 {
			return TwoPair
		}

		log.Fatal("Invalid combination of cards")
	}
	if len(values) == 4 {
		if values[0] == 1 && values[1] == 1 && values[2] == 1 && values[3] == 2 {
			return OnePair
		}

		log.Fatal("Invalid combination of cards")
	}

	return HighCard
}

func classifyHandWithJoker(cards [5]int) HandType {
	baseType := classifyHand(cards)
	if !slices.Contains(cards[:], 0) {
		return baseType
	}

	for i := 1; i < 13; i++ {
		cardCopy := slices.Clone(cards[:])

		// replace joker with i
		for j := range cardCopy {
			if cardCopy[j] == 0 {
				cardCopy[j] = i
			}
		}

		newType := classifyHand([5]int(cardCopy))
		baseType = max(newType, baseType)
	}

	return baseType
}

func cmpHands(a, b Hand) int {
	if a.handType < b.handType {
		return -1
	}
	if a.handType > b.handType {
		return 1
	}

	for i := 0; i < 5; i++ {
		if a.cards[i] < b.cards[i] {
			return -1
		}
		if a.cards[i] > b.cards[i] {
			return 1
		}
	}

	return 0
}

func parseHand(line string) Hand {
	str := strings.Split(line, " ")

	if len(str[0]) != 5 {
		log.Fatal("A hand must have 5 cards")
	}

	var result Hand

	for i, s := range str[0] {
		card := CardToOrder(s)
		if card == -1 {
			log.Fatal("Unrecognized card")
		}

		result.cards[i] = card
	}
	result.handType = classifyHand(result.cards)

	bid, err := strconv.Atoi(str[1])
	if err != nil {
		log.Fatal(err)
	}
	result.bid = bid

	return result
}

func parseHandJoker(line string) Hand {
	str := strings.Split(line, " ")

	if len(str[0]) != 5 {
		log.Fatal("A hand must have 5 cards")
	}

	var result Hand

	for i, s := range str[0] {
		card := CardToOrderJoker(s)
		if card == -1 {
			log.Fatal("Unrecognized card")
		}

		result.cards[i] = card
	}
	result.handType = classifyHandWithJoker(result.cards)

	bid, err := strconv.Atoi(str[1])
	if err != nil {
		log.Fatal(err)
	}
	result.bid = bid

	return result
}

func part1() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	s := bufio.NewScanner(file)

	hands := []Hand{}
	for s.Scan() {
		line := s.Text()

		hand := parseHand(line)
		hands = append(hands, hand)
	}

	slices.SortFunc(hands, cmpHands)

	sum := 0
	for i, h := range hands {
		sum += (i + 1) * h.bid
	}
	fmt.Printf("Part 1 sum: %v\n", sum)

	if err = s.Err(); err != nil {
		log.Fatal(err)
		return
	}
}

func part2() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	s := bufio.NewScanner(file)

	hands := []Hand{}
	for s.Scan() {
		line := s.Text()

		hand := parseHandJoker(line)
		hands = append(hands, hand)
	}

	slices.SortFunc(hands, cmpHands)

	sum := 0
	for i, h := range hands {
		sum += (i + 1) * h.bid
	}
	fmt.Printf("Part 2 sum: %v\n", sum)

	if err = s.Err(); err != nil {
		log.Fatal(err)
		return
	}
}

func main() {
	part1()
	part2()
}
