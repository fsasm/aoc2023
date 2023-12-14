package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func race(totalTime int, pressTime int) int {
	if pressTime > totalTime {
		log.Fatal("Press time should be less or equal the race time")
	}

	remainingTime := totalTime - pressTime
	speed := pressTime

	return remainingTime * speed
}

func tryRaceVariants(totalTime int, recordDistance int) int {
	count := 0
	for pressTime := 0; pressTime < totalTime; pressTime++ {
		distance := race(totalTime, pressTime)
		if distance > recordDistance {
			count++
		}
	}
	return count
}

func getIntegerList(line string) []int {
	str := strings.Split(line, ":")
	result := make([]int, 0)
	for _, s := range strings.Split(str[1], " ") {
		number, err := strconv.Atoi(s)
		if err != nil {
			continue
		}

		result = append(result, number)
	}

	return result
}

func getIntegerWithoutSpace(line string) int {
	str := strings.Split(line, ":")
	strNumber := strings.ReplaceAll(str[1], " ", "")

	number, err := strconv.Atoi(strNumber)
	if err != nil {
		return 0
	}

	return number
}

func part1() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	s := bufio.NewScanner(file)

	// FIXME there has to be a better way to read these lines, when we know
	// that only two lines are needed
	s.Scan()
	times := getIntegerList(s.Text())

	s.Scan()
	distances := getIntegerList(s.Text())

	product := 1
	for i := range times {
		product *= tryRaceVariants(times[i], distances[i])
	}

	fmt.Printf("part 1 product: %v\n", product)

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

	s.Scan()
	times := getIntegerWithoutSpace(s.Text())

	s.Scan()
	distances := getIntegerWithoutSpace(s.Text())

	numRaces := tryRaceVariants(times, distances)

	fmt.Printf("part 2 result: %v\n", numRaces)

	if err = s.Err(); err != nil {
		log.Fatal(err)
		return
	}
}

func main() {
	part1()
	part2()
}
