package day12

import (
	"adventofcode/utils"
	"fmt"
	"strings"
	"unicode"
)

type Cave struct {
	id    string
	links []string
}

func isBigCave(caveId string) bool {
	return unicode.IsUpper([]rune(caveId)[0])
}

func parseCaveLayout(links []string) map[string]Cave {
	allCaves := make(map[string]Cave)
	for _, link := range links {
		linkSplit := strings.Split(link, "-")

		idLeft := linkSplit[0]
		if _, ok := allCaves[idLeft]; !ok {
			lcave := Cave{idLeft, []string{}}
			allCaves[idLeft] = lcave
		}

		idRight := linkSplit[1]
		if _, ok := allCaves[idRight]; !ok {
			rcave := Cave{idRight, []string{}}
			allCaves[idRight] = rcave
		}

		lCave := allCaves[idLeft]
		rCave := allCaves[idRight]

		lCave.links = append(lCave.links, idRight)
		allCaves[idLeft] = lCave

		rCave.links = append(rCave.links, idLeft)
		allCaves[idRight] = rCave

	}

	return allCaves
}

func safeAppend(list []string, elem string) []string {
	new := make([]string, len(list)+1)
	copy(new, list)
	new[len(list)] = elem
	return new
}

func findNumberOfPaths(links []string) int {
	layout := parseCaveLayout(links)

	pathCandiates := [][]string{{"start"}}
	completePaths := [][]string{}

	for len(pathCandiates) > 0 {
		pc := pathCandiates[0]
		pathCandiates = pathCandiates[1:]

		pcLast := pc[len(pc)-1]

		nextLinks := layout[pcLast].links
		for _, link := range nextLinks {
			nextPathCandidate := safeAppend(pc, link) //Side effect!!
			if link == "end" {
				completePaths = append(completePaths, nextPathCandidate)
			} else if !containsSmallCycle(nextPathCandidate) {
				pathCandiates = append(pathCandiates, nextPathCandidate)
			}
		}
	}

	return len(completePaths)
}

func containMaxOneSmallCycle(nextPath []string) bool {
	lastElement := nextPath[len(nextPath)-1]
	if isBigCave(lastElement) {
		return false
	}

	loopCounter := 0
	count := make(map[string]int)
	for _, p := range nextPath {
		if !isBigCave(p) {
			if prevCount, ok := count[p]; ok {
				loopCounter++
				count[p] = prevCount + 1
			} else {
				count[p] = 1
			}
		}
	}
	return count["start"] > 1 || loopCounter > 1
}

func containsSmallCycle(nextPath []string) bool {
	lastElement := nextPath[len(nextPath)-1]
	if isBigCave(lastElement) {
		return false
	}

	for i := 0; i < len(nextPath)-2; i++ {
		if nextPath[i] == lastElement {
			return true
		}
	}
	return false
}

func findNumberOfPathsWithSingleSmallLoop(links []string) int {
	layout := parseCaveLayout(links)

	pathCandiates := [][]string{{"start"}}
	completePaths := [][]string{}

	for len(pathCandiates) > 0 {
		pc := pathCandiates[0]
		pathCandiates = pathCandiates[1:]
		pcLast := pc[len(pc)-1]

		nextLinks := layout[pcLast].links
		for _, link := range nextLinks {
			nextPathCandidate := safeAppend(pc, link) //Side effect!!
			if link == "end" {
				completePaths = append(completePaths, nextPathCandidate)
			} else if !containMaxOneSmallCycle(nextPathCandidate) {
				pathCandiates = append(pathCandiates, nextPathCandidate)
			}
		}
	}

	return len(completePaths)
}

func PrintSolution() {
	lines := utils.ParseLines("./inputs/day12.txt")
	paths := findNumberOfPaths(lines)
	fmt.Println("Different paths (Part 1)", paths)
	pathsWithMaxOneLoop := findNumberOfPathsWithSingleSmallLoop(lines)
	fmt.Println("Different paths with a single loop (Part 2)", pathsWithMaxOneLoop)
}
