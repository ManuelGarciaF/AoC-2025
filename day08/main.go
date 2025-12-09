package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/ManuelGarciaF/AoC-2025/commons"
)

type Coord struct {
	X, Y, Z int
}

func main() {
	coords := parseInput(os.Args[1])
	fmt.Println("Part 1: ", part1(coords, commons.MustAtoi(os.Args[2])))
	fmt.Println("Part 2: ", part2(coords))
}

type DistPair struct {
	idxA, idxB int
	distSquare int
}

func part1(coords []Coord, steps int) int {
	// Put each box in its own circuit.
	circuits := commons.Map(coords, func(c Coord) commons.Set[Coord] {
		return commons.NewSet[Coord]().Add(c)
	})

	dists := calculateDistances(coords)

	for i := 0; i < steps; i++ {
		// Merge only if they are on the same circuit
		cA := findCircuitIdx(coords[dists[i].idxA], circuits)
		cB := findCircuitIdx(coords[dists[i].idxB], circuits)
		if cA != cB {
			// fmt.Println("Merging", circuits[cA], "with", circuits[cB])
			// Merge circuits
			circuits[cA].Union(circuits[cB])
			circuits = slices.Delete(circuits, cB, cB+1)
		}
	}

	slices.SortFunc(circuits, func(a, b commons.Set[Coord]) int {
		return b.Size() - a.Size()
	})

	for i, s := range circuits {
		fmt.Println(i, s)
	}
	return circuits[0].Size() * circuits[1].Size() * circuits[2].Size()
}

func part2(coords []Coord) int {
	// Put each box in its own circuit.
	circuits := commons.Map(coords, func(c Coord) commons.Set[Coord] {
		return commons.NewSet[Coord]().Add(c)
	})

	dists := calculateDistances(coords)

	for _, d := range dists {
		boxA := coords[d.idxA]
		boxB := coords[d.idxB]

		// Merge only if they are on the same circuit
		cA := findCircuitIdx(boxA, circuits)
		cB := findCircuitIdx(boxB, circuits)
		if cA != cB {
			// Merge circuits
			circuits[cA].Union(circuits[cB])
			circuits = slices.Delete(circuits, cB, cB+1)

			if len(circuits) == 1 { // We just merged the last two circuits.
				return boxA.X * boxB.X
			}
		}
	}

	panic("Unreachable")
}

func findCircuitIdx(box Coord, circuits []commons.Set[Coord]) int {
	for i, circuit := range circuits {
		if circuit.Contains(box) {
			return i
		}
	}
	panic("Box isn't in any circuit")
}

// Get all distances and sort them
func calculateDistances(coords []Coord) []DistPair {
	dists := make([]DistPair, 0, len(coords)*len(coords))

	for i := 0; i < len(coords); i++ {
		for j := i + 1; j < len(coords); j++ {
			dists = append(dists, DistPair{
				idxA:       i,
				idxB:       j,
				distSquare: distSquare(coords[i], coords[j]),
			})
		}
	}

	// Sort by closest
	slices.SortFunc(dists, func(a, b DistPair) int {
		return a.distSquare - b.distSquare
	})
	return dists
}

// We don't actually need the real distance, d1 < d2 <=> d1^2 < d2^2.
// This way we avoid doing floating point calculations.
func distSquare(a, b Coord) int {
	xd := a.X - b.X
	yd := a.Y - b.Y
	zd := a.Z - b.Z
	return xd*xd + yd*yd + zd*zd
}

func parseInput(path string) []Coord {
	f := commons.Must(os.Open(path))
	defer f.Close()

	s := bufio.NewScanner(f)

	cs := make([]Coord, 0)

	for s.Scan() {
		nums := strings.Split(s.Text(), ",")
		cs = append(cs, Coord{
			X: commons.MustAtoi(nums[0]),
			Y: commons.MustAtoi(nums[1]),
			Z: commons.MustAtoi(nums[2]),
		})
	}

	return cs
}
