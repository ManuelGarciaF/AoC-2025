package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ManuelGarciaF/AoC-2025/commons"
)

func main() {
	banks := parseInput(os.Args[1])
	fmt.Println(banks)
	fmt.Println("Part 1: ", part1(banks))
	fmt.Println("Part 2: ", part2(banks))
}

func part1(banks [][]int) int {
	sum := 0
	for _, bank := range banks {
		first := maxIndex(bank, 0, len(bank)-1)
		second := maxIndex(bank, first+1, len(bank))
		value := 10*bank[first] + bank[second]
		// fmt.Println(bank, bank[:len(bank)-1], bank[first+1:], value)
		sum += value
	}
	return sum
}

func part2(banks [][]int) int {
	sum := 0
	for _, bank := range banks {
		indices := make([]int, 12)
		for i := 0; i < 12; i++ {
			if i == 0 {
				indices[0] = maxIndex(bank, 0, len(bank)-12)
			} else {
				indices[i] = maxIndex(bank, indices[i-1]+1, len(bank)-11+i)
			}
			
		}
		value := 0
		for _, i := range indices {
			value *= 10
			value += bank[i]
		}
		sum += value
	}
	return sum
}

func maxIndex(xs []int, s, e int) int {
	maxIndex := 0
	subs := xs[s:e]
	for i, x := range subs {
		if x > subs[maxIndex] {
			maxIndex = i
		}
	}
	return maxIndex + s
}

func parseInput(path string) [][]int {
	f := commons.Must(os.Open(path))
	defer f.Close()

	s := bufio.NewScanner(f)

	banks := make([][]int, 0)

	for s.Scan() {
		bank := make([]int, 0, len(s.Text()))
		for _, c := range s.Text() {
			bank = append(bank, int(c)-48)
		}
		banks = append(banks, bank)
	}

	return banks
}
