package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseIntegerList(line string) []int {
	str := strings.Split(line, " ")

	result := make([]int, 0)
	for _, s := range str {
		number, err := strconv.Atoi(s)
		if err != nil {
			continue
		}

		result = append(result, number)
	}

	return result
}

func extrapolate(seq []int) int {
	diff, zero := diffSlice(seq)
	lastNumber := seq[len(seq)-1]

	if zero {
		return lastNumber
	}
	return extrapolate(diff) + lastNumber
}

func extrapolateBackward(seq []int) int {
	diff, zero := diffSlice(seq)
	firstNumber := seq[0]

	if zero {
		return firstNumber
	}
	return firstNumber - extrapolateBackward(diff)
}

func diffSlice(seq []int) ([]int, bool) {
	result := make([]int, 0)
	allZero := true

	for i := 1; i < len(seq); i++ {
		diff := seq[i] - seq[i-1]
		result = append(result, diff)

		if diff != 0 {
			allZero = false
		}
	}

	return result, allZero
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

		sum += extrapolate(parseIntegerList(line))
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

	sum := 0
	for s.Scan() {
		line := s.Text()

		sum += extrapolateBackward(parseIntegerList(line))
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
