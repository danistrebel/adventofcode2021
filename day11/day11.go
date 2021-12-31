package day11

import (
	"adventofcode/utils"
	"fmt"
)

func copyMap(input [][]int) [10][10]int {
	copy := [10][10]int{}
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			copy[i][j] = input[i][j]
		}
	}
	return copy
}

func SimulateOctopusFlashes(input [][]int, days int) int {
	simulation := copyMap(input)

	flashes := 0

	for day := 0; day < days; day++ {
		octopusToIncrement := [][]int{}
		// init all octopusToIncrement
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				octopusToIncrement = append(octopusToIncrement, []int{i, j})
			}
		}

		for len(octopusToIncrement) > 0 {

			nextOctopus := octopusToIncrement[0]
			octopusToIncrement = octopusToIncrement[1:]

			valueBefore := simulation[nextOctopus[0]][nextOctopus[1]]
			if valueBefore >= 0 {
				if valueBefore == 9 {
					simulation[nextOctopus[0]][nextOctopus[1]] = -1

					for r := nextOctopus[0] - 1; r <= nextOctopus[0]+1; r++ {
						for c := nextOctopus[1] - 1; c <= nextOctopus[1]+1; c++ {
							if r >= 0 && r < 10 && c >= 0 && c < 10 {
								if !(r == nextOctopus[0] && c == nextOctopus[1]) {
									octopusToIncrement = append(octopusToIncrement, []int{r, c})
								}
							}
						}
					}
				} else {
					simulation[nextOctopus[0]][nextOctopus[1]] = valueBefore + 1
				}
			}
		}

		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				if simulation[i][j] < 0 {
					flashes++
					simulation[i][j] = 0
				}
			}
		}
	}

	return flashes
}

func PrintSolution() {
	input := utils.ParseIntArrays(utils.ParseLines("./inputs/day11.txt"), "")
	flashes := SimulateOctopusFlashes(input, 100)
	fmt.Println("Flashes (Part 1)", flashes)
}
