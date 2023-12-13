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

type Mapping struct {
	SrcStart int
	DstStart int
	Length   int
}

func ParseMapping(line string) (Mapping, error) {
	str := strings.Split(line, " ")

	dst, err := strconv.Atoi(str[0])
	if err != nil {
		return Mapping{}, err
	}

	src, err := strconv.Atoi(str[1])
	if err != nil {
		return Mapping{}, err
	}

	length, err := strconv.Atoi(str[2])
	if err != nil {
		return Mapping{}, err
	}

	return Mapping{
		SrcStart: src,
		DstStart: dst,
		Length:   length,
	}, nil
}

func (m Mapping) Contains(srcIndex int) bool {
	return m.SrcStart <= srcIndex && srcIndex < (m.SrcStart+m.Length)
}

func (m Mapping) Map(srcIndex int) int {
	return srcIndex - m.SrcStart + m.DstStart
}

type Mappings []Mapping

func (m Mappings) Map(index int) int {
	for _, mapping := range m {
		if mapping.Contains(index) {
			return mapping.Map(index)
		}
	}
	return index
}

func parseSeeds(line string) []int {
	// ignore "seeds:"
	str := strings.Split(line, ":")
	str = strings.Split(line, " ")

	result := make([]int, 0)
	for _, s := range str {
		num, err := strconv.Atoi(s)
		if err != nil {
			continue
		}

		result = append(result, num)
	}

	return result
}

func unpackSeedPairs(seedPairs []int) []int {
	if len(seedPairs)%2 != 0 {
		log.Fatal("Number of seeds have to be even")
	}

	/*
	 * Yes, this is a brute force approach. Expanding the seed pairs would create
	 * ~6 GiB of seeds. A better approach would be to still see them as ranges
	 * and during mapping see if ranges have be split and after mapping try to
	 * merge neighbouring ranges.
	 */
	totalSize := 0
	for i := 0; i < len(seedPairs); i += 2 {
		length := seedPairs[i+1]
		totalSize += length
	}

	seeds := make([]int, 0, totalSize)

	for i := 0; i < len(seedPairs); i += 2 {
		start := seedPairs[i]
		length := seedPairs[i+1]

		for j := 0; j < length; j++ {
			seed := start + j
			seeds = append(seeds, seed)
		}
	}

	return seeds
}

func part1() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	s := bufio.NewScanner(file)

	seeds := []int{}
	skip := false
	var mappings Mappings
	for i := 0; s.Scan(); i++ {
		line := s.Text()

		if i == 0 {
			seeds = parseSeeds(line)
			continue
		}

		if len(line) == 0 {
			// delimiter for new mapping
			skip = true // skip the header line

			if len(mappings) == 0 {
				continue
			}

			// map all seeds
			for j := range seeds {
				seeds[j] = mappings.Map(seeds[j])
			}

			// remove all mappings
			mappings = mappings[:0]
			continue
		}

		if skip {
			skip = false
			continue
		}

		mapping, err := ParseMapping(line)
		if err != nil {
			log.Fatal(err)
		}
		mappings = append(mappings, mapping)
	}

	// map the seeds again
	for j := range seeds {
		seeds[j] = mappings.Map(seeds[j])
	}

	fmt.Printf("part 1 seeds: %v\n", slices.Min(seeds))

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

	seeds := []int{}
	skip := false
	var mappings Mappings
	for i := 0; s.Scan(); i++ {
		line := s.Text()

		if i == 0 {
			seeds = unpackSeedPairs(parseSeeds(line))
			continue
		}

		if len(line) == 0 {
			// delimiter for new mapping
			skip = true // skip the header line

			if len(mappings) == 0 {
				continue
			}

			// map all seeds
			for j := range seeds {
				seeds[j] = mappings.Map(seeds[j])
			}

			// remove all mappings
			mappings = mappings[:0]
			continue
		}

		if skip {
			skip = false
			continue
		}

		mapping, err := ParseMapping(line)
		if err != nil {
			log.Fatal(err)
		}
		mappings = append(mappings, mapping)
	}

	// map the seeds again
	for j := range seeds {
		seeds[j] = mappings.Map(seeds[j])
	}

	fmt.Printf("part 2 seeds: %v\n", slices.Min(seeds))

	if err = s.Err(); err != nil {
		log.Fatal(err)
		return
	}
}

func main() {
	part1()
	part2()
}
