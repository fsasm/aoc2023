package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Schematic []string

func readSchematic(file *bufio.Scanner) Schematic {
	result := make(Schematic, 0)
	size := 0

	for file.Scan() {
		line := file.Text()
		if size == 0 {
			size = len(line)
		} else {
			if size != len(line) {
				log.Fatal("readSchematic: Line size mismatch")
			}
		}
		result = append(result, line)
	}

	if err := file.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

func (s Schematic) Width() int {
	return len(s[0])
}

func (s Schematic) Height() int {
	return len(s)
}

func (s Schematic) IsInside(pos Pos) bool {
	return pos.x >= 0 && pos.y >= 0 && pos.x < s.Width() && pos.y < s.Height()
}

func (s Schematic) Get(pos Pos) rune {
	if !s.IsInside(pos) {
		return 0
	}
	return rune(s[pos.y][pos.x])
}

func (s Schematic) IsSymbol(pos Pos) bool {
	if !s.IsInside(pos) {
		return false
	}
	r := rune(s[pos.y][pos.x])

	return runeToCellType(r) == SYMBOL
}

type Pos struct {
	x, y int
}

func (p *Pos) Next(skipX int, s Schematic) {
	p.x += skipX

	if p.x >= s.Width() {
		p.x = 0
		p.y++
	}
}

var digitMap = map[rune]int{
	'0': 0, '1': 1, '2': 2, '3': 3, '4': 4,
	'5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
}

type CellType int

const (
	NUMBER CellType = iota + 1
	SYMBOL
	NONE
)

func runeToCellType(c rune) CellType {
	if c == '.' {
		return NONE
	}

	if _, ok := digitMap[c]; ok {
		return NUMBER
	}

	return SYMBOL
}

func runeToDigit(r rune) int {
	if d, ok := digitMap[r]; ok {
		return d
	}
	return -1
}

func getNumberGroup(schematic Schematic, pos Pos) (int, int) {
	number := 0
	length := 0
	for true {
		if !schematic.IsInside(pos) {
			break
		}
		r := schematic.Get(pos)
		cell := runeToCellType(r)

		if cell != NUMBER {
			break
		}

		number = number*10 + runeToDigit(r)
		length++

		pos.x++
	}

	return number, length
}

func countPartNumbers(schematic Schematic) int {
	iter := Pos{0, 0}
	sum := 0
	sumAll := 0
	maxPartNo := 0

	for schematic.IsInside(iter) {
		r := schematic.Get(iter)
		cell := runeToCellType(r)

		if cell == NONE || cell == SYMBOL {
			iter.Next(1, schematic)
			continue
		}

		// cell can only be NUMBER
		partNumber, length := getNumberGroup(schematic, iter)
		sumAll += partNumber
		maxPartNo = max(maxPartNo, partNumber)

		isValid := false

		fromX := iter.x - 1
		toX := iter.x + length

		// left and right
		isValid = isValid || schematic.IsSymbol(Pos{fromX, iter.y})
		isValid = isValid || schematic.IsSymbol(Pos{toX, iter.y})

		// top and bottom row
		for x := fromX; x <= toX; x++ {
			isValid = isValid || schematic.IsSymbol(Pos{x, iter.y - 1})
			isValid = isValid || schematic.IsSymbol(Pos{x, iter.y + 1})
		}

		if isValid {
			sum += partNumber
		}

		iter.Next(length, schematic)
	}
	return sum
}

func searchNumberGroup(schematic Schematic, pos Pos) (number, length int) {
	if !schematic.IsInside(pos) {
		return
	}

	r := schematic.Get(pos)

	if runeToCellType(r) != NUMBER {
		return
	}

	// go left
	leftPos := pos
	leftPos.x--
	for true {
		if !schematic.IsInside(leftPos) {
			break
		}

		r = schematic.Get(leftPos)

		if runeToCellType(r) != NUMBER {
			break
		}
		leftPos.x--
	}
	// correct it because we previously overstepped
	leftPos.x++

	// go right
	rightPos := pos
	rightPos.x++
	for true {
		if !schematic.IsInside(rightPos) {
			break
		}

		r = schematic.Get(rightPos)

		if runeToCellType(r) != NUMBER {
			break
		}
		rightPos.x++
	}
	rightPos.x--

	length = rightPos.x - leftPos.x + 1

	for x := 0; x < length; x++ {
		pos.x = x + leftPos.x
		r = schematic.Get(pos)
		number = number*10 + runeToDigit(r)
	}

	return
}

func countGearRatios(schematic Schematic) int {
	iter := Pos{0, 0}
	sum := 0

	for schematic.IsInside(iter) {
		r := schematic.Get(iter)

		if r != '*' {
			iter.Next(1, schematic)
			continue
		}

		// check for number groups in all 8 positions
		product := 1
		cnt := 0

		// top row
		prevNum := 0
		if num, len := searchNumberGroup(schematic, Pos{iter.x - 1, iter.y - 1}); len > 0 {
			product *= num
			cnt++
			prevNum = num
		}
		if num, len := searchNumberGroup(schematic, Pos{iter.x, iter.y - 1}); len > 0 && prevNum != num {
			product *= num
			cnt++
			prevNum = num
		}
		if num, len := searchNumberGroup(schematic, Pos{iter.x + 1, iter.y - 1}); len > 0 && prevNum != num {
			product *= num
			cnt++
			prevNum = num
		}

		// bottom row
		if num, len := searchNumberGroup(schematic, Pos{iter.x - 1, iter.y + 1}); len > 0 {
			product *= num
			cnt++
			prevNum = num
		}
		if num, len := searchNumberGroup(schematic, Pos{iter.x, iter.y + 1}); len > 0 && prevNum != num {
			product *= num
			cnt++
			prevNum = num
		}
		if num, len := searchNumberGroup(schematic, Pos{iter.x + 1, iter.y + 1}); len > 0 && prevNum != num {
			product *= num
			cnt++
			prevNum = num
		}

		// left
		if num, len := searchNumberGroup(schematic, Pos{iter.x - 1, iter.y}); len > 0 {
			product *= num
			cnt++
		}

		// right
		if num, len := searchNumberGroup(schematic, Pos{iter.x + 1, iter.y}); len > 0 {
			product *= num
			cnt++
		}

		if cnt == 2 {
			sum += product
		}
		iter.Next(1, schematic)
	}

	return sum
}

func part1() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	s := bufio.NewScanner(file)

	schematic := readSchematic(s)
	sum := countPartNumbers(schematic)

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

	schematic := readSchematic(s)
	sum := countGearRatios(schematic)

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
