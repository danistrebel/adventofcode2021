package day17

import "fmt"

type coord struct {
	x int
	y int
}

type speed struct {
	vx int
	vy int
}

type target struct {
	xMin int
	xMax int
	yMin int
	yMax int
}

func (c coord) onTarget(t target) bool {
	return t.xMin <= c.x && c.x <= t.xMax && t.yMin <= c.y && c.y <= t.yMax
}

func (s speed) applyDrag() speed {
	if s.vx > 0 {
		return speed{vx: s.vx - 1, vy: s.vy - 1}
	} else if s.vx < 0 {
		return speed{vx: s.vx + 1, vy: s.vy - 1}
	} else {
		return speed{vx: 0, vy: s.vy - 1}
	}
}

func simulateNext(c coord, s speed) (coord, speed) {
	return coord{x: c.x + s.vx, y: c.y + s.vy}, s.applyDrag()
}

func testAngle(s speed, t target) bool {
	position := coord{0, 0}
	maxY := 0
	for position.y > t.yMin && position.x < t.xMax {
		position, s = simulateNext(position, s)
		if position.y > maxY {
			maxY = position.y
		}
		if position.onTarget(t) {
			return true
		}
	}
	return false
}

func findHighestY(t target) int {
	return (t.yMin * (t.yMin + 1)) / 2
}

func findAllSettings(t target) int {
	maxVX := t.xMax
	minVX := 0
	for (minVX*(minVX+1))/2 < t.xMin {
		minVX++
	}

	maxVY := -t.yMin - 1
	minVY := t.yMin

	hitCount := 0

	for vx := minVX; vx <= maxVX; vx++ {
		for vy := minVY; vy <= maxVY; vy++ {
			if testAngle(speed{vx, vy}, t) {
				hitCount++
			}
		}
	}
	return hitCount
}

func PrintSolution() {
	//target area: x=257..286, y=-101..-57
	t := target{257, 286, -101, -57}
	maxY := findHighestY(t)
	fmt.Println("Highest Y (Part 1)", maxY)
	validSettings := findAllSettings(t)
	fmt.Println("Valid Settings (Part 2)", validSettings)
}
