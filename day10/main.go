package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"

	"github.com/ManuelGarciaF/AoC-2025/commons"

	"github.com/draffensperger/golp"
)

type Machine struct {
	Lights   commons.BitArray16
	Buttons  [][]int
	Joltages []int
}

func main() {
	machines := parseInput(os.Args[1])
	// for _, m := range machines {
	// 	fmt.Printf("m: %v\n", m)
	// }
	fmt.Println("Part1: ", commons.Sum(commons.Map(machines, minMovesLights)))
	fmt.Println("Part2: ", part2(machines))
}

type LightState struct {
	Lights commons.BitArray16
	Moves  int
}

func minMovesLights(m Machine) int {
	// Bfs
	q := make(commons.Queue[LightState], 0)
	q.Push(LightState{
		Lights: 0,
		Moves:  0,
	})

	for !q.IsEmpty() {
		state := q.Pop()

		// Queue all possible button presses
		for _, b := range m.Buttons {
			newState := state
			newState.Moves++
			// Toggle all lights
			for _, light := range b {
				newState.Lights = newState.Lights.Toggle(light)
			}
			if newState.Lights == m.Lights {
				return newState.Moves
			}
			q.Push(newState)
		}
	}

	panic("Unreachable")
}

// Need to return an answer as a float to avoid wrong rounding
func part2(ms []Machine) float64 {
	sum := float64(0)
	for _, m := range ms {
		sum += minMovesJoltages(m)
	}
	return sum
}

func minMovesJoltages(machine Machine) float64 {
	// Integer linear programming problem, solve using lp_solve
	m := len(machine.Buttons)
	lp := golp.NewLP(0, m)

	// Minimize the number of presses, all coeficients are 1
	lp.SetObjFn(repeat(m, 1))

	for j := range m {
		lp.SetInt(j, true)
	}

	// Constraints
	for i, joltage := range machine.Joltages {
		row := make([]float64, m)
		for j, b := range machine.Buttons {
			if slices.Contains(b, i) {
				row[j] = 1
			}
		}
		lp.AddConstraint(row, golp.EQ, float64(joltage))
	}
	lp.Solve()

	// fmt.Printf("lp.Variables(): %v (%v)\n", lp.Variables(), lp.Objective())
	return lp.Objective()
}

func repeat(n int, v float64) []float64 {
	ns := make([]float64, n)
	for i := range n {
		ns[i] = v
	}
	return ns
}

/*
 * Working solution, but will not finish until after the heat death of the universe.
 */

// type JoltageState struct {
// 	Joltages []int
// 	Moves    int
// }

// func minMovesJoltages(m Machine) int {
// 	// A* kinda
// 	q := commons.NewPriorityQueue[JoltageState]()
// 	initialState := JoltageState{
// 		Joltages: make([]int, len(m.Joltages)),
// 		Moves:    0,
// 	}
// 	q.PushItem(initialState, distance(m.Joltages, initialState.Joltages))

// 	visited := make(map[string]int, 0)

// 	maxDepth := 0

// 	for !q.IsEmpty() {
// 		state, _ := q.PopItem()

// 		// Queue all benefitial button presses
// 		for _, b := range benefitialButtons(m, state.Joltages) {
// 			newState := JoltageState{
// 				Joltages: make([]int, len(m.Joltages)),
// 				Moves:    state.Moves + 1,
// 			}
// 			copy(newState.Joltages, state.Joltages)

// 			for _, counter := range b {
// 				newState.Joltages[counter]++
// 			}

// 			// Discard already visited configurations
// 			currSer := serialize(newState.Joltages)
// 			currCost := newState.Moves+distance(m.Joltages, newState.Joltages)
// 			if cost, ok := visited[currSer]; ok && cost <= currCost {
// 				continue
// 			}
// 			visited[currSer] = currCost

// 			// Found a match
// 			if slices.Equal(newState.Joltages, m.Joltages) {
// 				fmt.Println()
// 				fmt.Println("Done:", newState.Moves, "moves")
// 				return newState.Moves
// 			}

// 			if maxDepth < newState.Moves {
// 				maxDepth = newState.Moves
// 				fmt.Printf("\rDepth: %v    ", maxDepth)
// 			}


// 			q.PushItem(newState, currCost)
// 		}
// 	}

// 	panic("No valid solution")
// }

// func benefitialButtons(m Machine, currJoltages []int) [][]int {
// 	bs := make([][]int, 0)
// outer:
// 	for _, b := range m.Buttons {
// 		for _, counterIdx := range b {
// 			if currJoltages[counterIdx] >= m.Joltages[counterIdx] {
// 				continue outer
// 			}
// 		}
// 		bs = append(bs, b)
// 	}
// 	return bs
// }

// func distance(target, curr []int) int {
// 	return commons.Sum(target) - commons.Sum(curr)
// }

// func serialize(vs []int) string {
// 	var sb strings.Builder
// 	for _, v := range vs {
// 		sb.WriteString(strconv.Itoa(v))
// 		sb.WriteByte(',')
// 	}
// 	return sb.String()
// }

func parseInput(path string) []Machine {
	f := commons.Must(os.Open(path))
	defer f.Close()

	s := bufio.NewScanner(f)

	reLights := regexp.MustCompile(`\[([.#]+)\]`)
	reButtons := regexp.MustCompile(`\(([\d,?]+)\)`)
	reJoltages := regexp.MustCompile(`{([\d,?]+)}`)

	ms := make([]Machine, 0)

	for s.Scan() {
		var lights commons.BitArray16
		lightsStr := reLights.FindStringSubmatch(s.Text())[1]
		if len(lightsStr) > 16 {
			panic("Too many lights")
		}
		for i, c := range lightsStr {
			if c == '#' {
				lights = lights.Set(i)
			}
		}

		buttonsRes := reButtons.FindAllStringSubmatch(s.Text(), -1)
		buttons := make([][]int, 0, len(buttonsRes))
		for _, match := range buttonsRes {
			buttons = append(buttons, commons.AtoiMap(strings.Split(match[1], ",")))
		}

		joltages := commons.AtoiMap(strings.Split(reJoltages.FindStringSubmatch(s.Text())[1], ","))

		ms = append(ms, Machine{
			Lights:   lights,
			Buttons:  buttons,
			Joltages: joltages,
		})
	}

	return ms
}
