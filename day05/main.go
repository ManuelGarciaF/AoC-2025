package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/ManuelGarciaF/AoC-2025/commons"
)

type Range struct {
	S, E int
}

func (r Range) Contains(n int) bool {
	return n >= r.S && n <= r.E
}

func main() {
	ranges, ingredients := parseInput(os.Args[1])
	// Sort all ranges and ingredients to make search faster
	slices.Sort(ingredients)
	slices.SortFunc(ranges, func(a, b Range) int {
		return a.S - b.S
	})

	fmt.Println("Part 1: ", part1(ranges, ingredients))
	fmt.Println("Part 2: ", part2(ranges))
}

func part1(ranges []Range, ingredients []int) int {
	count := 0

	for _, ing := range ingredients {
		for _, r := range ranges {
			if r.Contains(ing) {
				count++
				break
			}
		}
	}
	return count
}

func part2(ranges []Range) int {
	count := 0
	maximum := slices.MaxFunc(ranges, func(a, b Range) int {
		return a.E - b.E
	}).E

	for pos, rangeIdx := 0, 0; pos <= maximum && rangeIdx < len(ranges); { // Sweep through all values.
		r := ranges[rangeIdx]
		if r.Contains(pos) { // We advance to the end, counting all values
			oldI := pos
			pos = r.E + 1
			count += pos - oldI
		} else if r.S > pos { // We still have to wait for the next range to start.
			pos = r.S
		} else { // We aren't contained nor are we before the start, we continue to the next range.
			rangeIdx++
		}
	}

	return count
}

func parseInput(path string) ([]Range, []int) {
	f := commons.Must(os.Open(path))
	defer f.Close()

	s := bufio.NewScanner(f)

	ranges := make([]Range, 0)
	ingredients := make([]int, 0)

	inRanges := true
	for s.Scan() {
		if s.Text() == "" {
			inRanges = false
			continue
		}
		if inRanges {
			numStrs := strings.Split(s.Text(), "-")
			if len(numStrs) != 2 {
				panic("Invalid input")
			}
			ranges = append(ranges, Range{
				S: commons.MustAtoi(numStrs[0]),
				E: commons.MustAtoi(numStrs[1]),
			})
		} else {
			ingredients = append(ingredients, commons.MustAtoi(s.Text()))
		}
	}

	return ranges, ingredients
}
