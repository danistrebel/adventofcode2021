package day6

import (
	"adventofcode/utils"
	"fmt"
)

func SimulateLatnernFishPopulation(initialPopulation []int, daysSimulation int) []int {
	simulatedPopulation := initialPopulation
	for day := 1; day <= daysSimulation; day++ {
		newFish := []int{}
		for i, counter := range simulatedPopulation {
			if counter == 0 {
				simulatedPopulation[i] = 6
				newFish = append(newFish, 8)
			} else {
				simulatedPopulation[i] = counter - 1
			}
		}
		simulatedPopulation = append(simulatedPopulation, newFish...)
	}
	return simulatedPopulation
}

func PrintSolution() {
	initalPopulation := utils.ParseIntArrays(utils.ParseLines("./inputs/day6.txt"), ",")
	simulatedPopulation := SimulateLatnernFishPopulation(initalPopulation[0], 80)
	fmt.Println("Simulated Population Size (Part 1)", len(simulatedPopulation))
}
