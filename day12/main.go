package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/ManuelGarciaF/AoC-2025/commons"
)

type Shape [][]bool

type Grid struct {
	xSize, ySize int
	reqShapes []int
}

func main() {
	shapes, grids := parseInput(os.Args[1])
	// fmt.Printf("shapes: %v\n", shapes)
	// fmt.Printf("grids: %v\n", grids)
	fmt.Println("Part1: ", part1(shapes, grids))
	// fmt.Println("Part2: ", part2(devices))
}

func part1(shapes []Shape, grids []Grid) any {
	count := 0
	for _, grid := range grids {
		if shapesFit(grid, shapes) {
			count++
		}
	}

	return count
}

func shapesFit(grid Grid, shapes []Shape) bool {
	// Discard grid if total area is smaller than shape areas
	ga := grid.xSize * grid.ySize
	sa := 0
	for i, count := range grid.reqShapes {
		sa += shapeArea(shapes[i])*count
	}
	if ga < sa {
		return false
	}

	// Maybe that check is enough?
	
	return true
}

// Could be precalculated if slow
func shapeArea(s Shape) int {
	area := 0
	for _, row := range s {
		for _, v := range row {
			if v {
				area++
			}
		}
	}
	return area
}

func parseInput(path string) ([]Shape, []Grid) {
	text := string(commons.Must(os.ReadFile(path)))

	shapeRe := regexp.MustCompile(`\d:\n(([.#]*\n){3})`)
	gridRe := regexp.MustCompile(`(\d+x\d+):(( \d+)+)`)

	ss := commons.Map(shapeRe.FindAllStringSubmatch(text, -1), func(groups []string) Shape {
		shape := make(Shape, 0)
		shapeStr := groups[1][:len(groups[1])-1] // Remove trailing \n
		for line := range strings.SplitSeq(shapeStr, "\n") {
			row := make([]bool, len(line))
			for i, c := range line {
				row[i] = c == '#'
			}
			shape = append(shape, row)
		}
		return shape
	})
	gs := commons.Map(gridRe.FindAllStringSubmatch(text, -1), func(groups []string) Grid {
		dims := commons.AtoiMap(strings.Split(groups[1], "x"))
		return Grid{
			xSize:     dims[0],
			ySize:     dims[1],
			reqShapes: commons.AtoiMap(strings.Split(groups[2][1:], " ")),
		}
	})

	return ss, gs
}
