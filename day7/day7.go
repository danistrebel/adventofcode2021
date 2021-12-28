package day7

import (
	"adventofcode/utils"
	"fmt"
	"math"
	"sort"
)

func OptimalCrabsPosition(initalPositions []int) int {
	sort.Ints(initalPositions)
	var median int
	if len(initalPositions)%2 == 0 {
		median = (initalPositions[len(initalPositions)/2-1] + initalPositions[len(initalPositions)/2]) / 2
	} else {
		median = initalPositions[len(initalPositions)/2]
	}

	fuel := 0

	for _, p := range initalPositions {
		if median > p {
			fuel += median - p
		} else {
			fuel += p - median
		}
	}
	return fuel
}

func OptimalCrabsPositionTwo(initalPositions []int) int {
	sum := 0
	for _, p := range initalPositions {
		sum += p
	}
	meanFloor := math.Floor(float64(sum) / float64(len(initalPositions)))
	meanCeil := math.Ceil(float64(sum) / float64(len(initalPositions)))

	fuelFloor := calcFuel(initalPositions, meanFloor)
	fuelCeil := calcFuel(initalPositions, meanCeil)

	if fuelFloor < fuelCeil {
		return fuelFloor
	} else {
		return fuelCeil
	}
}

func calcFuel(positions []int, center float64) int {
	fuel := 0

	for _, p := range positions {
		delta := int(math.Abs(float64(p) - center))
		fuelP := (delta * (delta + 1)) / 2
		fuel += fuelP
	}
	return fuel
}

func PrintSolution() {
	initialPositions := utils.ParseIntArrays(utils.ParseLines("./inputs/day7.txt"), ",")
	fuel := OptimalCrabsPosition(initialPositions[0])
	fmt.Println("Fuel required (Part 1)", fuel)
	fuelTwo := OptimalCrabsPositionTwo(initialPositions[0])
	fmt.Println("Fuel required (Part 2)", fuelTwo)
}
