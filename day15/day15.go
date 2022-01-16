package day15

import (
	"adventofcode/utils"
	"container/heap"
	"fmt"
)

type Item struct {
	row   int
	col   int
	g     int
	f     int
	index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].f <= pq[j].f
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func estimateH(row int, col int, rowCount int, colCount int) int {
	averageStepCost := 1
	return ((rowCount - row - 1) + (colCount - col - 1)) * averageStepCost
}

func calculateMinRisk(input [][]int) int {
	rowCount := len(input)
	colCount := len(input[0])

	minGMatrix := make([][]int, rowCount)

	for r := 0; r < rowCount; r++ {
		colMatrix := make([]int, colCount)
		for c := 0; c < colCount; c++ {
			colMatrix[c] = 999999
		}
		minGMatrix[r] = colMatrix
	}

	q := make(PriorityQueue, 0)
	heap.Init(&q)
	heap.Push(&q, &Item{0, 0, 0, estimateH(0, 0, rowCount, colCount), 0})

	for len(q) > 0 {
		cand := heap.Pop(&q).(*Item)
		if cand.row == rowCount-1 && cand.col == colCount-1 {
			return cand.g
		}

		neighborCandidates := [][]int{}

		if cand.row-1 >= 0 {
			neighborCandidates = append(neighborCandidates, []int{cand.row - 1, cand.col})
		}
		if cand.row+1 < rowCount {
			neighborCandidates = append(neighborCandidates, []int{cand.row + 1, cand.col})
		}
		if cand.col-1 >= 0 {
			neighborCandidates = append(neighborCandidates, []int{cand.row, cand.col - 1})
		}
		if cand.col+1 < colCount {
			neighborCandidates = append(neighborCandidates, []int{cand.row, cand.col + 1})
		}

		for _, nc := range neighborCandidates {
			newRow := nc[0]
			newCol := nc[1]
			newG := cand.g + input[newRow][newCol]
			if minGMatrix[newRow][newCol] > newG {
				minGMatrix[newRow][newCol] = newG
				newF := newG + estimateH(newRow, newCol, rowCount, colCount)
				heap.Push(&q, &Item{newRow, newCol, newG, newF, len(q)})
			}
		}
	}

	return 9999
}

func PrintSolution() {
	lines := utils.ParseLines("inputs/day15.txt")
	input := utils.ParseIntArrays(lines, "")
	minRisk := calculateMinRisk(input)
	fmt.Println("Min Risk (Part 1)", minRisk)
}
