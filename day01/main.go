package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ManuelGarciaF/AoC-2025/commons"
)

func main() {
	rs := parseInput(os.Args[1])
	fmt.Println("Part 1: ", part1(rs))
	fmt.Println("Part 2 (wrong): ", part2(rs))
	fmt.Println("Part 2 brute force: ", part2Brute(rs))
}

func part1(rs []int) int {
	count := 0
	position := 50
	for _, r := range rs {
		position = (position + r + 100) % 100
		if position == 0 {
			count++
		}
	}
	return count
}

// Really can't figure out why it doesn't work
func part2(rs []int) int {
	count := 0
	pos := 50
	for _, r := range rs {
		// fmt.Println("next rot: ", r)
		if pos+r > 99 {
			count += ((pos + r) / 100)
			// fmt.Println("Turned ", (pos+r)/100, "times right")
		} else if pos+r < 0 && pos != 0 {
			count += -((pos + r) / 100) + 1
			// fmt.Println("Turned ", -((pos+r)/100)+1, "times left")
		} else if pos+r == 0 {
			count++
			// fmt.Println("Landed on 0 +1")
		}

		pos = ((pos+r)%100 + 100) % 100
		if pos == 0 {

		}
		// fmt.Println("Final Position: ", pos)
	}

	return count
}

func part2Brute(rs []int) int {
	count := 0
	pos := 50
	for _, r := range rs {
		if r > 0 {
			for i := 0; i < r; i++ {
				pos = (pos + 1) % 100
				if pos == 0 {
					count++
				}
			}
		} else {
			for i := 0; i < -r; i++ {
				pos = (pos + 99) % 100
				if pos == 0 {
					count++
				}
			}
		}
	}

	return count
}

func parseInput(path string) []int {
	file := commons.Must(os.Open(path))
	defer file.Close()

	s := bufio.NewScanner(file)

	rs := make([]int, 0)

	for s.Scan() {
		l := s.Text()
		switch l[0] {
		case 'L':
			rs = append(rs, -commons.MustAtoi(l[1:]))
		case 'R':
			rs = append(rs, commons.MustAtoi(l[1:]))
		}
	}

	return rs
}
