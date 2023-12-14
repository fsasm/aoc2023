package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

type Pair struct {
	left  string
	right string
}

func parseNode(line string) (string, Pair) {
	regex := regexp.MustCompile(`^([A-Z0-9]+) = \(([A-Z0-9]+), ([A-Z0-9]+)\)`)

	parts := regex.FindStringSubmatch(line)
	if len(parts) != 4 {
		log.Fatal("Something wrong with parsing")
	}

	return parts[1], Pair{parts[2], parts[3]}
}

func iteratePath(startNode string, nodes map[string]Pair, sequence string) int {
	steps := 0
	curNode := startNode

	for true {
		for _, s := range sequence {
			if s == 'L' {
				curNode = nodes[curNode].left
			} else if s == 'R' {
				curNode = nodes[curNode].right
			} else {
				log.Fatal("Invalid sequence code")
			}

			steps++
			if curNode[len(curNode)-1] == 'Z' {
				return steps
			}
		}
	}
	return 0
}

func gcd(a, b int64) int64 {
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}

	return a
}

func lcm(a, b int64) int64 {
	return (a * b) / gcd(a, b)
}

func lcmSlice(s []int) int64 {
	t := lcm(int64(s[0]), int64(s[1]))

	for i := 2; i < len(s); i++ {
		t = lcm(t, int64(s[i]))
	}
	return t
}

func part1() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	s := bufio.NewScanner(file)

	// L-R sequence
	s.Scan()
	sequence := s.Text()

	s.Scan()
	s.Text()

	nodes := make(map[string]Pair)
	for s.Scan() {
		line := s.Text()

		name, node := parseNode(line)
		nodes[name] = node
	}

	steps := 0
	curNode := "AAA"
outerLoop:
	for true {
		for _, s := range sequence {
			if s == 'L' {
				curNode = nodes[curNode].left
			} else if s == 'R' {
				curNode = nodes[curNode].right
			} else {
				log.Fatal("Invalid sequence code")
			}

			steps++
			if curNode == "ZZZ" {
				break outerLoop
			}
		}
	}

	fmt.Printf("Part 1 steps: %v\n", steps)

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

	// L-R sequence
	s.Scan()
	sequence := s.Text()

	s.Scan()
	s.Text()

	nodes := make(map[string]Pair)
	var curNodes []string
	for s.Scan() {
		line := s.Text()

		name, node := parseNode(line)
		nodes[name] = node

		if name[len(name)-1] == 'A' {
			curNodes = append(curNodes, name)
		}
	}

	steps := make([]int, len(curNodes))
	for i, n := range curNodes {
		steps[i] = iteratePath(n, nodes, sequence)
	}

	totalSteps := lcmSlice(steps)

	fmt.Printf("Part 2 steps: %v\n", totalSteps)

	if err = s.Err(); err != nil {
		log.Fatal(err)
		return
	}
}

func main() {
	part1()
	part2()
}
