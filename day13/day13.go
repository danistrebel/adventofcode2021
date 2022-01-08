package day13

import (
	"adventofcode/utils"
	"fmt"
	"strconv"
	"strings"
)

type FoldInstruction struct {
	axis  string
	coord int
}

func parseInput(lines []string) ([][]bool, []FoldInstruction) {
	folds := []FoldInstruction{}
	maxWidth := 0
	maxHeight := 0

	points := [][]int{}

	for _, l := range lines {
		if strings.HasPrefix(l, "fold along") {
			axis := l[11:12]
			coord, _ := strconv.Atoi(l[13:])
			folds = append(folds, FoldInstruction{axis: axis, coord: coord})
		} else if len(l) > 0 {
			split := strings.Split(l, ",")
			x, _ := strconv.Atoi(split[0])
			y, _ := strconv.Atoi(split[1])
			if x > maxWidth {
				maxWidth = x
			}
			if y > maxHeight {
				maxHeight = y
			}
			points = append(points, []int{x, y})

		}
	}

	inputPoints := make([][]bool, maxWidth+2)
	for i := 0; i < maxWidth+2; i++ {
		inputPoints[i] = make([]bool, maxHeight+2)
	}

	for _, point := range points {
		inputPoints[point[0]][point[1]] = true
	}

	return inputPoints, folds
}

func printGrid(grid [][]bool) {
	for c := range grid[0] {
		for r := range grid {
			if grid[r][c] {
				fmt.Print("# ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func foldGrid(grid [][]bool, f FoldInstruction) [][]bool {

	var newWidh int
	var newHeight int
	var oldWidh = len(grid)
	var oldHeight = len(grid[0])

	if f.axis == "x" {
		newWidh = f.coord
	} else {
		newWidh = oldWidh
	}

	if f.axis == "y" {
		newHeight = f.coord
	} else {
		newHeight = oldHeight
	}

	newGrid := make([][]bool, newWidh)
	for i := 0; i < newWidh; i++ {
		newGrid[i] = make([]bool, newHeight)
	}

	for x := 0; x < newWidh; x++ {
		for y := 0; y < newHeight; y++ {
			if f.axis == "x" {
				newGrid[x][y] = grid[x][y] || grid[2*f.coord-x][y]
			} else if f.axis == "y" {
				newGrid[x][y] = grid[x][y] || grid[x][2*f.coord-y]
			}
		}
	}
	return newGrid
}

func countDots(grid [][]bool) int {
	count := 0
	for c := range grid[0] {
		for r := range grid {
			if grid[r][c] {
				count++
			}
		}
	}
	return count
}

func PrintSolution() {
	lines := utils.ParseLines("./inputs/day13.txt")
	grid, folds := parseInput(lines)

	gridFold := foldGrid(grid, folds[0])

	fmt.Println("Dots visible (Part 1)", countDots(gridFold))
}
