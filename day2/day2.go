package day2

import (
	"adventofcode/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func calculateCoordinates(course []string) (int, int) {
	horizontal := 0
	depth := 0
	for _, instr := range course {
		instrSplit := strings.Split(instr, " ")

		distance, err := strconv.Atoi(instrSplit[1])
		if err != nil {
			fmt.Printf("Error Parsing Number: %s \n", err.Error())
		}

		switch instrSplit[0] {
		case "forward":
			horizontal += distance
		case "down":
			depth += distance
		case "up":
			depth -= distance
		default:
			log.Fatal("unknown instruction:", instrSplit[0])
		}
	}
	return horizontal, depth
}

func calculateAimCoordinates(course []string) (int, int) {
	horizontal := 0
	depth := 0
	aim := 0
	for _, instr := range course {
		instrSplit := strings.Split(instr, " ")

		distance, err := strconv.Atoi(instrSplit[1])
		if err != nil {
			fmt.Printf("Error Parsing Number: %s \n", err.Error())
		}

		switch instrSplit[0] {
		case "forward":
			horizontal += distance
			depth += aim * distance
		case "down":
			aim += distance
		case "up":
			aim -= distance
		default:
			log.Fatal("unknown instruction:", instrSplit[0])
		}
	}
	return horizontal, depth
}

func PrintSolution() {
	course := utils.ParseLines("./inputs/day2.txt")
	x, d := calculateCoordinates(course)
	fmt.Println("Result Position (Part 1)", x*d)
	xa, da := calculateAimCoordinates(course)
	fmt.Println("Result Position (Part 2)", xa*da)
}
