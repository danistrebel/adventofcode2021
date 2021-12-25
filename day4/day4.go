package day4

import (
	"adventofcode/utils"
	"fmt"
)

type BingoBoardEntry = struct {
	x      int
	marked bool
}

type BingoBoard = [][]BingoBoardEntry

func InitBingoBoards(inputs [][]int) []BingoBoard {
	boards := []BingoBoard{}

	currentBoard := BingoBoard{}
	for _, l := range inputs {
		if len(l) == 0 {
			boards = append(boards, currentBoard)
			currentBoard = BingoBoard{}
		} else {
			row := []struct {
				x      int
				marked bool
			}{}
			for _, x := range l {
				row = append(row, BingoBoardEntry{x: x, marked: false})
			}
			currentBoard = append(currentBoard, row)
		}
	}
	return boards
}

func printBoard(board BingoBoard) {
	for _, row := range board {
		for _, e := range row {
			if e.marked {
				fmt.Print("*")
			}
			fmt.Print(e.x, "  ")
		}
		fmt.Println()
	}
}

func isCompleteBoard(board BingoBoard) bool {
	for _, row := range board {
		completeRow := true
		for _, entry := range row {
			completeRow = completeRow && entry.marked
			if !completeRow {
				break
			}
		}
		if completeRow {
			return true
		}
	}

	for col := 0; col < len(board[0]); col++ {
		completeCol := true
		for row := 0; row < len(board); row++ {
			completeCol = completeCol && board[row][col].marked
			if !completeCol {
				break
			}
		}
		if completeCol {
			return true
		}
	}
	return false
}

func computeBoardScore(board BingoBoard, currentNumber int) int {
	sumUnmarked := 0
	for _, row := range board {
		for _, entry := range row {
			if !entry.marked {
				sumUnmarked += entry.x
			}
		}
	}
	return currentNumber * sumUnmarked
}

func PlayBingo(boards []BingoBoard, seq []int) int {
	for _, n := range seq {
		for _, board := range boards {
			for _, row := range board {
				for entryIndex, entry := range row {
					if entry.x == n {
						row[entryIndex].marked = true
						break
					}
				}
			}

			if isCompleteBoard(board) {
				return computeBoardScore(board, n)
			}
		}
	}
	return -1
}

func PlayLoserBingo(boards []BingoBoard, seq []int) int {
	for _, n := range seq {
		for boardIndex := len(boards) - 1; boardIndex >= 0; boardIndex-- {
			board := boards[boardIndex]
			for _, row := range board {
				for entryIndex, entry := range row {
					if entry.x == n {
						row[entryIndex].marked = true
						break
					}
				}
				if isCompleteBoard(board) {
					if len(boards) == 1 {
						return computeBoardScore(board, n)
					} else {
						boards = append(boards[:boardIndex], boards[boardIndex+1:]...)
						break
					}
				}
			}

		}
	}
	return -1
}

func PrintSolution() {
	input := utils.ParseLines("./inputs/day4.txt")
	boardInput := utils.ParseIntArraysFromFields(append(input[2:], ""))
	boards := InitBingoBoards(boardInput)
	score := PlayBingo(boards, utils.ParseIntArrays(input[0:1], ",")[0])
	fmt.Println("Bingo score (Part 1)", score)

	input = utils.ParseLines("./inputs/day4.txt")
	boardInput = utils.ParseIntArraysFromFields(append(input[2:], ""))
	boards = InitBingoBoards(boardInput)
	score = PlayLoserBingo(boards, utils.ParseIntArrays(input[0:1], ",")[0])
	fmt.Println("Loser Bingo score (Part 2)", score)
}
