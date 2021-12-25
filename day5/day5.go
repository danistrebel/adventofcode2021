package day5

import (
	"adventofcode/utils"
	"fmt"
	"strconv"
	"strings"
)

func orderedPath(a int, b int) (int, int) {
	if a < b {
		return a, b
	} else {
		return b, a
	}
}

func ParseWindMatrix(winds []string, parseDiagonal bool) [][]uint8 {
	matrixDim := 1000
	matrix := make([][]uint8, matrixDim)
	for i := range matrix {
		matrix[i] = make([]uint8, matrixDim)
	}

	for _, wind := range winds {
		windSplits := strings.Split(wind, " -> ")

		startSplits := strings.Split(windSplits[0], ",")
		endSplits := strings.Split(windSplits[1], ",")

		startX, _ := strconv.Atoi(startSplits[0])
		startY, _ := strconv.Atoi(startSplits[1])
		endX, _ := strconv.Atoi(endSplits[0])
		endY, _ := strconv.Atoi(endSplits[1])

		// vertical
		if startX == endX {
			minY, maxY := orderedPath(startY, endY)

			for y := minY; y <= maxY; y++ {
				matrix[startX][y] = matrix[startX][y] + 1
			}
		}

		// horizontal
		if startY == endY {
			minX, maxX := orderedPath(startX, endX)

			for x := minX; x <= maxX; x++ {
				matrix[x][startY] = matrix[x][startY] + 1
			}
		}

		// diagonal
		if parseDiagonal {
			minY, maxY := orderedPath(startY, endY)
			minX, maxX := orderedPath(startX, endX)

			if maxY-minY == maxX-minX {

				for d := 0; d <= maxX-minX; d++ {
					if startX < endX && startY < endY {
						matrix[startX+d][startY+d] = matrix[startX+d][startY+d] + 1
					} else if startX < endX && startY > endY {
						matrix[startX+d][startY-d] = matrix[startX+d][startY-d] + 1
					} else if startX > endX && startY < endY {
						matrix[startX-d][startY+d] = matrix[startX-d][startY+d] + 1
					} else {
						matrix[startX-d][startY-d] = matrix[startX-d][startY-d] + 1
					}
				}
			}
		}
	}

	return matrix
}

func CountTresholdValues(matrix [][]uint8, treshold uint8) int {
	thresholdCounter := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] >= treshold {
				thresholdCounter++
			}
		}
	}
	return thresholdCounter
}

func PrintSolution() {
	wind := utils.ParseLines("./inputs/day5.txt")
	windMatrix := ParseWindMatrix(wind, false)
	counter := CountTresholdValues(windMatrix, 2)
	fmt.Println("Wind Counter (Part 1)", counter)

	windMatrixDiagonal := ParseWindMatrix(wind, true)
	counterDiagonal := CountTresholdValues(windMatrixDiagonal, 2)
	fmt.Println("Wind Counter with Diagonal (Part 2)", counterDiagonal)
}
