package day9

import (
	"adventofcode/utils"
	"fmt"
)

func isLowPoint(inputs [][]int, r int, c int) bool {
	v := inputs[r][c]
	smaller := true
	if r > 0 && v >= inputs[r-1][c] {
		smaller = false
	}
	if smaller && c > 0 && v >= inputs[r][c-1] {
		smaller = false
	}
	if smaller && r < len(inputs)-1 && v >= inputs[r+1][c] {
		smaller = false
	}
	if smaller && c < len(inputs[r])-1 && v >= inputs[r][c+1] {
		smaller = false
	}
	return smaller
}

func totalRiskLevel(inputs [][]int) int {
	totalRisk := 0
	for r := range inputs {
		for c, v := range inputs[r] {

			if isLowPoint(inputs, r, c) {
				totalRisk += v + 1
			}
		}
	}
	return totalRisk
}

func expandBaisin(inputs [][]int, lowPoint []int) int {
	fringe := [][]int{lowPoint}
	baisinSize := 0

	for len(fringe) > 0 {
		r := fringe[0][0]
		c := fringe[0][1]
		fringe = fringe[1:]
		v := inputs[r][c]
		inputs[r][c] = 10 // Mark as visited

		if v >= 9 {
			continue
		}
		baisinSize++

		if r > 0 && v < inputs[r-1][c] {
			fringe = append(fringe, []int{r - 1, c})
		}
		if c > 0 && v < inputs[r][c-1] {
			fringe = append(fringe, []int{r, c - 1})
		}
		if r < len(inputs)-1 && v < inputs[r+1][c] {
			fringe = append(fringe, []int{r + 1, c})
		}
		if c < len(inputs[r])-1 && v < inputs[r][c+1] {
			fringe = append(fringe, []int{r, c + 1})
		}
	}
	return baisinSize
}

func largestBaisins(inputs [][]int) int {
	lowPoints := [][]int{}
	for r := range inputs {
		for c := range inputs[r] {
			if isLowPoint(inputs, r, c) {
				lowPoints = append(lowPoints, []int{r, c})
			}
		}
	}

	top3Baisins := [3]int{}
	for _, lowPoint := range lowPoints {
		size := expandBaisin(inputs, lowPoint)
		if size > top3Baisins[0] {
			top3Baisins[2] = top3Baisins[1]
			top3Baisins[1] = top3Baisins[0]
			top3Baisins[0] = size
		} else if size > top3Baisins[1] {
			top3Baisins[2] = top3Baisins[1]
			top3Baisins[1] = size
		} else if size > top3Baisins[2] {
			top3Baisins[2] = size
		}
	}

	return top3Baisins[0] * top3Baisins[1] * top3Baisins[2]
}

func PrintSolution() {
	lines := utils.ParseLines("./inputs/day9.txt")
	input := utils.ParseIntArrays(lines, "")
	riskLevel := totalRiskLevel(input)
	fmt.Println("Risk Level (Day 1)", riskLevel)
	biggestBaisins := largestBaisins(input)
	fmt.Println("Largest Baisins (Day 2)", biggestBaisins)
}
