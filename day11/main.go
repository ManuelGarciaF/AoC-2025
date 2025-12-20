package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ManuelGarciaF/AoC-2025/commons"
)

func main() {
	devices := parseInput(os.Args[1])
	fmt.Println("Part1: ", solve(devices, "you", "out"))
	fmt.Println("Part2: ", part2(devices))
}

// Topological sort and count paths to each node
func solve(devices map[string][]string, start, end string) int {
	// Construct topological order
	inDegree := make(map[string]int, len(devices))
	for _, outs := range devices {
		for _, out := range outs {
			inDegree[out]++
		}
	}
	// Kahn's algorithm, we add nodes with no more neighbors to the ordering.
	q := make(commons.Queue[string], 0)
	for d := range devices {
		if inDegree[d] == 0 {
			q.Push(d)
		}
	}
	topologicalOrder := make([]string, 0, len(devices))
	for !q.IsEmpty() {
		d := q.Pop()
		topologicalOrder = append(topologicalOrder, d)

		for _, neighbor := range devices[d] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				q.Push(neighbor)
			}
		}
	}

	paths := make(map[string]int, len(devices))
	paths[start] = 1
	for _, node := range topologicalOrder {
		for _, neighbor := range devices[node] {
			paths[neighbor] += paths[node]
		}
	}

	return paths[end]
}

func part2(devices map[string][]string) int {
	order1 := solve(devices, "svr", "fft") * solve(devices, "fft", "dac") * solve(devices, "dac", "out")
	order2 := solve(devices, "svr", "dac") * solve(devices, "dac", "fft") * solve(devices, "fft", "out")
	return order1 + order2
}

func parseInput(path string) map[string][]string {
	f := commons.Must(os.Open(path))
	defer f.Close()

	s := bufio.NewScanner(f)

	ds := make(map[string][]string)

	for s.Scan() {
		t := s.Text()
		name := t[:3]
		outputs := strings.Split(t[5:], " ")
		ds[name] = outputs
	}

	return ds
}
