package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ManuelGarciaF/AoC-2025/commons"
)

type Range struct{ Start, End int }

func main() {
	rs := parseInput(os.Args[1])
	fmt.Println(rs)
	fmt.Println("Part 1: ", part1(rs))
	fmt.Println("Part 2: ", part2(rs))
}

func part1(rs []Range) int {
	total := 0
	for _, r := range rs {
		for curr := r.Start; curr <= r.End; curr++ {
			cs := strconv.Itoa(curr)

			if len(cs) % 2 != 0 {
				continue
			}

			if cs[:len(cs)/2] == cs[len(cs)/2:] {
				total += curr
			}
		}
	}
	return total
}

// Spaghetti code
func part2(rs []Range) int {
	total := 0
	for _, r := range rs {
		for curr := r.Start; curr <= r.End; curr++ {
			cs := strconv.Itoa(curr)

			for repeatLen := 1; repeatLen <= len(cs)/2; repeatLen++ {
				// Check the length divides the str
				if len(cs) % repeatLen != 0 {
					continue
				}

				pattern := cs[:repeatLen]

				valid := true
				for i := repeatLen; i < len(cs); i += repeatLen {
					if cs[i:i+repeatLen] != pattern {
						valid = false
						break
					}
				}
				if valid {
					fmt.Println(r, curr)
					total += curr
					break
				}
			}
		}
	}
	return total
}

func parseInput(path string) []Range {
	file := commons.Must(os.ReadFile(path))
	rangeStrs := strings.Split(strings.TrimRight(string(file), "\n"), ",")

	rs := make([]Range, 0)

	for _, rStr := range rangeStrs {
		parts := strings.Split(rStr, "-")
		rs = append(rs, Range{
			Start: commons.MustAtoi(parts[0]),
			End: commons.MustAtoi(parts[1]),
		})
	}

	return rs
}
