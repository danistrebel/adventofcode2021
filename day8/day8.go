package day8

import (
	"adventofcode/utils"
	"fmt"
	"log"
	"sort"
	"strings"
)

func CountUniqueOutputDigits(data []string) int {

	uniqueDigitsMap := map[int]int{
		2: 1,
		3: 7,
		4: 4,
		7: 8,
	}
	counter := 0

	for _, d := range data {

		_, outDigits := parseLine(d)

		for _, outDigit := range outDigits {
			if _, ok := uniqueDigitsMap[len(outDigit)]; ok {
				counter++
			}
		}
	}
	return counter
}

func parseLine(line string) ([]string, []string) {
	lineSplit := strings.Split(line, " | ")
	patterns := strings.Split(lineSplit[0], " ")
	outDigits := strings.Split(lineSplit[1], " ")
	return patterns, outDigits

}

func SumDecodedOuts(inputs []string) int {
	sum := 0
	for _, input := range inputs {
		sum += DecodeOutputDigit(input)
	}
	return sum
}

func DecodeOutputDigit(input string) int {
	patterns, outDigits := parseLine(input)

	uniqueDigitsMap := map[int]int{
		2: 1,
		3: 7,
		4: 4,
		7: 8,
	}

	numberPatternMap := make(map[int]string)
	patternNumberMap := make(map[string]int)
	for _, p := range patterns {
		if d, ok := uniqueDigitsMap[len(p)]; ok {
			sortedP := sortChars(p)
			numberPatternMap[d] = sortedP
			patternNumberMap[sortedP] = d
		}
	}

	for _, p := range patterns {
		sortedP := sortChars(p)
		if _, ok := patternNumberMap[sortedP]; !ok {
			commonOne := commonChars(sortedP, numberPatternMap[1])
			commonFour := commonChars(sortedP, numberPatternMap[4])
			if len(sortedP) == 5 {
				if commonOne == 1 && commonFour == 2 {
					patternNumberMap[sortedP] = 2
				} else if commonOne == 2 && commonFour == 3 {
					patternNumberMap[sortedP] = 3
				} else if commonOne == 1 && commonFour == 3 {
					patternNumberMap[sortedP] = 5
				} else {
					log.Fatalln("unknown pattern", p, "1:", numberPatternMap[1], "4:", numberPatternMap[4])
					return -1
				}
			} else if len(sortedP) == 6 {
				if commonOne == 2 && commonFour == 3 {
					patternNumberMap[sortedP] = 0
				} else if commonOne == 1 && commonFour == 3 {
					patternNumberMap[sortedP] = 6
				} else if commonOne == 2 && commonFour == 4 {
					patternNumberMap[sortedP] = 9
				} else {
					log.Fatalln("unknown pattern", p, "1:", numberPatternMap[1], "4:", numberPatternMap[4])
					return -1
				}
			}
		}
	}

	outValue := 0
	positionMultiplier := 1
	for i := len(outDigits) - 1; i >= 0; i-- {
		sortedOut := sortChars(outDigits[i])
		outValue += positionMultiplier * patternNumberMap[sortedOut]
		positionMultiplier *= 10
	}
	return outValue
}

func commonChars(a string, b string) int {
	aMap := make(map[string]bool)
	for _, ac := range strings.Split(a, "") {
		aMap[ac] = true
	}
	count := 0
	for _, bc := range strings.Split(b, "") {
		_, ok := aMap[bc]
		if ok {
			count++
		}
	}
	return count
}

func sortChars(p string) string {
	chars := strings.Split(p, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

func PrintSolution() {
	data := utils.ParseLines("./inputs/day8.txt")
	count := CountUniqueOutputDigits(data)
	fmt.Println("Unique digits count (Part 1)", count)

	sumOfOuts := SumDecodedOuts(data)
	fmt.Println("Sum of Output Values (Part 2)", sumOfOuts)

}
