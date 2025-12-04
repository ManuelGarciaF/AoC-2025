package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ManuelGarciaF/AoC-2025/commons"
)

func main() {
	positions := parseInput(os.Args[1])
	fmt.Println("Part 1: ", part1(positions))
	fmt.Println("Part 2: ", part2(positions))
}

func part1(positions [][]bool) int {
	count := 0
	for y, row := range positions {
		for x, v := range row {
			if v && countNeighbors(positions, commons.Coord{X: x, Y: y}) < 4 {
				count++
			}
		}
	}
	return count
}

func part2(positions [][]bool) int {
	count := 0

	for {
		removed := false
		for y, row := range positions {
			for x := range row {
				if positions[y][x] && countNeighbors(positions, commons.Coord{X: x, Y: y}) < 4 {
					removed = true
					positions[y][x] = false
					count++
				}
			}
		}
		if !removed {
			break
		}
	}
	return count
}

func countNeighbors(positions [][]bool, c commons.Coord) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			// Out of bounds
			pos := commons.Coord{X: c.X+j, Y:c.Y+i}
			if !pos.Inbounds(len(positions[0])-1, len(positions)-1) || (i==0 && j==0) {
				continue
			}
			if positions[pos.Y][pos.X] {
				count++
			}
		}
	}
	return count
}

func parseInput(path string) [][]bool {
	f := commons.Must(os.Open(path))
	defer f.Close()

	s := bufio.NewScanner(f)

	positions := make([][]bool, 0)

	for s.Scan() {
		row := make([]bool, len(s.Text()))
		for i, c := range s.Text() {
			row[i] = c == '@'
		}
		positions = append(positions, row)
	}

	return positions
}
