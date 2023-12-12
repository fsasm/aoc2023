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

func parseNumberList(str string) []int {
	result := make([]int, 0)

	for _, token := range strings.Split(str, " ") {
		number, err := strconv.Atoi(token)
		if err != nil {
			continue
		}

		result = append(result, number)
	}

	return result
}
 
func processLine1(line string) int {
	numMatches := processLine2(line)

	if numMatches == 0 {
		return 0
	}

	return 1 << (numMatches - 1)
}
 
func processLine2(line string) int {
	str := strings.Split(line, ":")

	str = strings.Split(str[1], "|")
	winningNum := parseNumberList(str[0])
	haveNum := parseNumberList(str[1])

	numMatches := 0
	for _, w := range winningNum {
		if slices.Contains(haveNum, w) {
			numMatches++
			continue
		}
	}

	return numMatches
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
		sum += processLine1(line)
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

	matchTable := make([]int, 0)
	for s.Scan() {
		line := s.Text()
		matchTable = append(matchTable, processLine2(line))
	}


	copiesTable := make([]int, len(matchTable))
	for i, w := range matchTable {
		numCopies := copiesTable[i] + 1
		for j := 0; j < w; j++ {
			copiesTable[i+1+j] += numCopies
		}
	}
	sum := 0
	for _, m := range copiesTable {
		sum += m + 1 // +1 because we also have to count the originals and not only the copies
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
