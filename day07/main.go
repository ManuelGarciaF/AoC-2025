package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"

	"github.com/ManuelGarciaF/AoC-2025/commons"
)

type BoardElem int

const (
	Empty BoardElem = iota
	Splitter
	Start
)

func main() {
	board := parseInput(os.Args[1])
	fmt.Println("Part 1: ", part1(board))
	fmt.Println("Part 2: ", part2(board))
}

func part1(board [][]BoardElem) int {
	beams := make(commons.Queue[commons.Coord], 0)

	beams.Push(commons.Coord{
		X: slices.Index(board[0], Start),
		Y: 0,
	})

	splittersReached := make(commons.Set[commons.Coord])

	for !beams.IsEmpty() {
		beam := beams.Pop()
		below := beam.Move(commons.DOWN)

		// End of board, beam is done.
		if below.Y >= len(board) {
			continue
		}

		// Check on what's below
		switch commons.IndexMap(board, below) {
		case Empty:
			beams.Push(below)
		case Splitter:
			if splittersReached.Contains(below) { // Already done
				continue
			}

			// New splitter
			splittersReached.Add(below)
			// Add both sides
			beams.Push(
				below.Move(commons.LEFT),
				below.Move(commons.RIGHT),
			)
		}
	}

	return splittersReached.Size()
}

func part2(board [][]BoardElem) int {
	var timelinesFrom func(commons.Coord) int
	timelinesFrom = commons.Memoize(func(c commons.Coord) int {
		below := c.Move(commons.DOWN)
		if below.Y >= len(board) {
			return 1
		}
		switch commons.IndexMap(board, below) {
		case Empty:
			return timelinesFrom(below)
		case Splitter:
			return timelinesFrom(below.Move(commons.LEFT)) + timelinesFrom(below.Move(commons.RIGHT))
		}
		panic("Unreachable")
	})

	start := commons.Coord{
		X: slices.Index(board[0], Start),
		Y: 0,
	}

	return timelinesFrom(start)
}

func parseInput(path string) [][]BoardElem {
	f := commons.Must(os.Open(path))
	defer f.Close()

	s := bufio.NewScanner(f)

	board := make([][]BoardElem, 0)

	for s.Scan() {
		board = append(board, commons.Map(s.Bytes(), func(s byte) BoardElem {
			switch s {
			case '.':
				return Empty
			case '^':
				return Splitter
			case 'S':
				return Start
			default:
				panic("Invalid Input")
			}
		}))
	}

	return board
}
