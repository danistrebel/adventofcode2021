package day1

import (
	"adventofcode/utils"
	"fmt"
)

func countIncreases(sonarLog []int) int {
	increaseCounter := 0
	for i := 1; i < len(sonarLog); i++ {
		if sonarLog[i] > sonarLog[i-1] {
			increaseCounter++
		}
	}
	return increaseCounter
}

func windowSum(sonarLog []int, windowSize int, lastIndex int) int {
	sum := 0
	for _, l := range sonarLog[lastIndex-windowSize+1 : lastIndex+1] {
		sum += l
	}
	return sum
}

func countWindowIncreases(sonarLog []int, windowSize int) int {
	increaseCounter := 0

	for i := windowSize; i < len(sonarLog); i++ {
		if windowSum(sonarLog, windowSize, i) > windowSum(sonarLog, windowSize, i-1) {
			increaseCounter++
		}
	}
	return increaseCounter
}

func PrintSolution() {
	sonarLog := utils.ParseInts("./inputs/day1.txt")
	increases := countIncreases(sonarLog)
	fmt.Println("increases (Part 1)", increases)
	windowIncreases := countWindowIncreases(sonarLog, 3)
	fmt.Println("increases (Part 2)", windowIncreases)
}
