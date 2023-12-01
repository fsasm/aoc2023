package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Pair struct {
	n string
	d int
}

var digitMap = map[rune]int{
	'0': 0, '1': 1, '2': 2, '3': 3, '4': 4,
	'5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
}

var digitNames = []Pair{
	{"one", 1}, {"two", 2}, {"three", 3}, {"four", 4},
	{"five", 5}, {"six", 6}, {"seven", 7}, {"eight", 8},
	{"nine", 9},
}

func toDigitWithNames(line string) (int, bool) {

	if digit, ok := digitMap[rune(line[0])]; ok {
		return digit, true
	}

	for _, pair := range digitNames {
		if len(pair.n) <= len(line) {
			if line[0:len(pair.n)] == pair.n {
				return pair.d, true
			}
		}
	}

	return 0, false
}

func processLine1(line string) (a, b int) {
	first := true
	for _, c := range line {
		digit, ok := digitMap[c]
		if !ok {
			continue
		}

		if first {
			a = digit
			b = digit
			first = false
		} else {
			b = digit
		}
	}

	return
}

func processLine2(line string) (a, b int) {

	for i := 0; i < len(line); i++ {
		if digit, ok := toDigitWithNames(line[i:]); ok {
			a = digit
			b = digit
			break
		}
	}

	// go backwards
	for i := len(line) - 1; i >= 0; i-- {
		if digit, ok := toDigitWithNames(line[i:]); ok {
			b = digit
			break
		}
	}

	return
}

func part1() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	s := bufio.NewScanner(file)

	sum := 0
	for s.Scan() {
		line := s.Text()

		a, b := processLine1(line)

		sum += a*10 + b
	}

	fmt.Printf("part 1 sum: %v\n", sum)

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

	sum := 0
	for s.Scan() {
		line := s.Text()
		a, b := processLine2(line)

		sum += a*10 + b
	}

	fmt.Printf("part 2 sum: %v\n", sum)

	if err = s.Err(); err != nil {
		log.Fatal(err)
		return
	}
}

func main() {
	part1()
	part2()
}
