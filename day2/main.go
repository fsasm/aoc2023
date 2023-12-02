package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const MAX_RED = 12
const MAX_GREEN = 13
const MAX_BLUE = 14

func getGameID(s string) (int, bool) {
	gameRegex := regexp.MustCompile(`^Game ([0-9]+)`)

	parts := gameRegex.FindStringSubmatch(s)
	id, err := strconv.Atoi(string(parts[1]))

	if err != nil {
		fmt.Fprintf(os.Stderr, "getGameID: %v\n", err)
		return 0, false
	}
	return id, true
}

func getColors(set string) (int, int, int) {
	redRegex := regexp.MustCompile(`([0-9]+) red`)
	greenRegex := regexp.MustCompile(`([0-9]+) green`)
	blueRegex := regexp.MustCompile(`([0-9]+) blue`)

	redSum := 0
	greenSum := 0
	blueSum := 0

	for _, color := range strings.Split(set, ",") {
		if redRegex.MatchString(color) {
			parts := redRegex.FindStringSubmatch(color)
			r, _ := strconv.Atoi(parts[1])
			redSum += r
		}
		if greenRegex.MatchString(color) {
			parts := greenRegex.FindStringSubmatch(color)
			g, _ := strconv.Atoi(parts[1])
			greenSum += g
		}
		if blueRegex.MatchString(color) {
			parts := blueRegex.FindStringSubmatch(color)
			b, _ := strconv.Atoi(parts[1])
			blueSum += b
		}
	}
	return redSum, greenSum, blueSum
}

func processLine1(line string) (int, bool) {
	// 1. split : to get ID
	// 2. split ; to get sets
	// 3. split , to get colors

	str := strings.Split(line, ":")
	id, _ := getGameID(str[0])

	sets := strings.Split(str[1], ";")
	for _, set := range sets {
		r, g, b := getColors(set)

		if r > MAX_RED || g > MAX_GREEN || b > MAX_BLUE {
			return id, false
		}
	}

	return id, true
}

func processLine2(line string) (int, bool) {
	// same idea with split as in processLine1

	str := strings.Split(line, ":")

	sets := strings.Split(str[1], ";")
	minRed := 0
	minGreen := 0
	minBlue := 0
	for _, set := range sets {
		r, g, b := getColors(set)

		minRed = max(r, minRed)
		minGreen = max(g, minGreen)
		minBlue = max(b, minBlue)
	}

	return minRed * minGreen * minBlue, true
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

		if id, ok := processLine1(line); ok {
			sum += id
		}
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
		if id, ok := processLine2(line); ok {
			sum += id
		}
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
