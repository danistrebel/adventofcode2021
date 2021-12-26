package day6

import (
	"adventofcode/utils"
	"fmt"
)

func SimulateLatnernFishPopulation(initialPopulation []int, daysSimulation int) int {
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
	return len(simulatedPopulation)
}

func CompactSimulateLatnernFishPopulation(initialPopulation []int, daysSimulation int) uint64 {
	simulatedPopulation := make([]uint64, 9)
	for _, p := range initialPopulation {
		simulatedPopulation[p] = simulatedPopulation[p] + 1
	}
	for day := 1; day <= daysSimulation; day++ {
		newFish := simulatedPopulation[0]
		for i := 0; i < 8; i++ {
			simulatedPopulation[i] = simulatedPopulation[i+1]
		}
		simulatedPopulation[8] = newFish
		simulatedPopulation[6] = simulatedPopulation[6] + newFish
	}

	populationSize := uint64(0)
	for _, size := range simulatedPopulation {
		populationSize += size
	}
	return populationSize
}

func PrintSolution() {
	initalPopulation := utils.ParseIntArrays(utils.ParseLines("./inputs/day6.txt"), ",")
	simulatedPopulation := SimulateLatnernFishPopulation(initalPopulation[0], 80)
	fmt.Println("Simulated Population Size after 80 days(Part 1)", simulatedPopulation)
	initalPopulation = utils.ParseIntArrays(utils.ParseLines("./inputs/day6.txt"), ",")
	largeSimulatedPopulation := CompactSimulateLatnernFishPopulation(initalPopulation[0], 256)
	fmt.Println("Simulated Population Size after 256 (Part 2)", largeSimulatedPopulation)
}
