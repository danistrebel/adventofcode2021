package day14

import (
	"adventofcode/utils"
	"fmt"
	"strings"
)

func simulateSteps(template string, rules []string, steps int) map[string]int {
	instructions := make(map[string]string)

	for _, r := range rules {
		split := strings.Split(r, " -> ")
		instructions[split[0]] = split[1]
	}

	state := make(map[string]int)
	for i := 0; i < len(template)-1; i++ {
		pattern := template[i : i+2]
		s := state[pattern]
		s++
		state[pattern] = s
	}

	for step := 1; step <= steps; step++ {
		newState := make(map[string]int)
		for pattern, count := range state {
			replacement := instructions[pattern]
			newPairA := pattern[:1] + replacement
			newPairB := replacement + pattern[1:]

			aCount := newState[newPairA]
			aCount += count
			newState[newPairA] = aCount

			bCount := newState[newPairB]
			bCount += count
			newState[newPairB] = bCount
		}

		state = newState
	}

	return state
}

func countPolymerCountsDiff(state map[string]int, last string) int {
	counts := make(map[string]int)
	for e, count := range state {
		splits := strings.Split(e, "")
		counts[splits[0]] = counts[splits[0]] + count
	}
	counts[last] = counts[last] + 1

	minCount := 999999999999999999
	maxCount := 0

	for _, count := range counts {
		if count > maxCount {
			maxCount = count
		}
		if count < minCount {
			minCount = count
		}
	}

	return maxCount - minCount
}

func polymerSimulation(inputs []string, steps int) int {
	template := inputs[0]
	state := simulateSteps(template, inputs[2:], steps)
	return countPolymerCountsDiff(state, template[len(template)-1:])
}

func PrintSolution() {
	inputs := utils.ParseLines("./inputs/day14.txt")
	simulation := polymerSimulation(inputs, 10)
	fmt.Println("Poymer Simulation (Part 1)", simulation)
}
