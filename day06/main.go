package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ManuelGarciaF/AoC-2025/commons"
)

type Operation int

const (
	Add Operation = iota
	Mult
)

type Problem struct {
	Ns []int
	O  Operation
}

func main() {
	fmt.Println("Part 1: ", part1())
	fmt.Println("Part 2: ", part2())
}

func part1() int {
	operations := parsePart1(os.Args[1])
	sum := 0
	for _, o := range operations {
		term := 0
		if o.O == Mult {
			term = 1
		}
		for _, i := range o.Ns {
			if o.O == Add {
				term += i
			} else {
				term *= i
			}
		}
		sum += term
	}
	return sum
}

func parsePart1(path string) []Problem {
	f := commons.Must(os.Open(path))
	defer f.Close()

	s := bufio.NewScanner(f)

	ps := make([]Problem, 0)

	parsedNums := make([][]int, 0)
	var ops []Operation
	for s.Scan() {
		strs := strings.Fields(s.Text())

		if strs[0] == "*" || strs[0] == "+" {
			ops = commons.Map(strs, func(s string) Operation {
				switch s {
				case "+":
					return Add
				case "*":
					return Mult
				default:
					panic("Invalid Operation")
				}
			})
			break
		}

		parsedNums = append(parsedNums, commons.Map(strs, commons.MustAtoi))
	}

	// Order now that we have them all parsed
	for i := 0; i < len(ops); i++ {
		ns := make([]int, 0, len(parsedNums))
		for j := 0; j < len(parsedNums); j++ {
			ns = append(ns, parsedNums[j][i])
		}

		ps = append(ps, Problem{
			Ns: ns,
			O:  ops[i],
		})
	}

	return ps
}

// The easiest way to solve this is by iterating through the raw text, fun!
// Prepare for spaghetti
func part2() int {
	lines := parsePart2(os.Args[1])

	// Iterate column by column, starting from the rightmost one.
	sum := 0
	problem := make([]int, 0, 4)
	var op byte
	for col := len(lines[0]) - 1; col >= 0; col-- {
		emptyCol := true
		num := 0
		for row := 0; row < len(lines); row++ {
			c := lines[row][col]
			if c >= '0' && c <= '9' {
				num *= 10
				num += int(lines[row][col] - '0')
				emptyCol = false
			}
			if c == '*' || c == '+' { // We found the operation
				op = c
				emptyCol = false
			}
		}
		if !emptyCol {
			problem = append(problem, num)
		}
		if emptyCol || col == 0 { // We reached a separator, we know the whole problem now, time to calculate the result
			switch op {
			case '*':
				sum += commons.Foldl(1, problem, func(a, b int) int { return a * b })
			case '+':
				sum += commons.Sum(problem)
			default:
				panic("Invalid Operation")
			}
			// Reset things for next problem
			problem = make([]int, 0, 4)
		}
	}

	return sum
}

func parsePart2(path string) []string {
	text := commons.Must(os.ReadFile(path))
	lines := strings.Split(string(text), "\n")

	return lines[:len(lines)-1] // One empty line is added
}
