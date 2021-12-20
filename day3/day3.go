package day3

import (
	"adventofcode/utils"
	"fmt"
	"strconv"
)

func powerConsumptionRates(diagReport []string) (int, int) {

	binaryGamma := ""
	binaryEpsilon := ""
	for i := range diagReport[0] {
		oneCounter := 0
		zeroCounter := 0
		for _, r := range diagReport {
			if r[i] == '1' {
				oneCounter++
			} else {
				zeroCounter++
			}
		}
		if oneCounter > zeroCounter {
			binaryGamma += "1"
			binaryEpsilon += "0"
		} else {
			binaryGamma += "0"
			binaryEpsilon += "1"
		}
	}

	gamma, _ := strconv.ParseInt(binaryGamma, 2, 32)
	epsilon, _ := strconv.ParseInt(binaryEpsilon, 2, 32)
	return int(gamma), int(epsilon)
}

func oxygenGenerator(diagReport []string) int {
	return bitCriteria(diagReport, 0, 1)
}

func co2ScrubberRating(diagReport []string) int {
	return bitCriteria(diagReport, 0, -1)
}

func bitCriteria(diagReport []string, position int, inverter int) int {
	if len(diagReport) == 1 {
		result, _ := strconv.ParseInt(diagReport[0], 2, 32)
		return int(result)
	}

	zeros := []string{}
	ones := []string{}

	for _, r := range diagReport {
		if r[position] == '1' {
			ones = append(ones, r)
		} else {
			zeros = append(zeros, r)
		}
	}

	if (len(ones)*inverter > len(zeros)*inverter) || (len(ones) == len(zeros) && inverter == 1) {
		return bitCriteria(ones, position+1, inverter)
	} else {
		return bitCriteria(zeros, position+1, inverter)
	}

}

func PrintSolution() {
	input := utils.ParseLines("./inputs/day3.txt")
	g, e := powerConsumptionRates(input)
	fmt.Println("Power Consumption (Part 1)", g*e)

	o := oxygenGenerator(input)
	c := co2ScrubberRating(input)

	fmt.Println("Life support rating (Part 2)", o*c)
}
