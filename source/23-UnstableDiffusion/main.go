package main

import (
	"fmt"
	"strings"

	"github.com/LenaBullens/advent-of-code-2022-go/source/helper"
)

type Point struct {
	x int
	y int
}

func createPoint(x int, y int) Point {
	return Point{x: x, y: y}
}

func main() {
	solve1()
	solve2()
}

func solve1() {
	lines := helper.ReadLines("input-23.txt")

	elves := make(map[Point]bool)

	for row := 0; row < len(lines); row++ {
		splitLine := strings.Split(lines[row], "")
		for column := 0; column < len(splitLine); column++ {
			if splitLine[column] == "#" {
				point := createPoint(column, row)
				elves[point] = true
			}
		}
	}

	directions := [4]string{"N", "S", "W", "E"}

	for i := 0; i < 10; i++ {
		proposedMovements := make(map[Point]Point)
		duplicateProtection := make(map[Point]int)

		for elf, _ := range elves {
			neighbours := getNeighbouringElves(elves, elf)
			if len(neighbours) == 0 {
				proposedMovements[elf] = elf
				duplicateProtection[elf] = duplicateProtection[elf] + 1
			} else {
				point := getProposedMovement(neighbours, elf, directions)
				proposedMovements[elf] = point
				duplicateProtection[point] = duplicateProtection[point] + 1
			}
		}

		newElves := make(map[Point]bool)
		for origin, target := range proposedMovements {
			//Check if not bumping
			if duplicateProtection[target] >= 2 {
				newElves[origin] = true
			} else {
				newElves[target] = true
			}
		}

		elves = newElves
		//shift directions
		var newDirections [4]string
		newDirections[0] = directions[1]
		newDirections[1] = directions[2]
		newDirections[2] = directions[3]
		newDirections[3] = directions[0]
		directions = newDirections
	}

	//Find boundaries
	xmin := 0
	xmax := 0
	ymin := 0
	ymax := 0
	first := true
	for elf, _ := range elves {
		if first {
			xmin = elf.x
			xmax = elf.x
			ymin = elf.y
			ymax = elf.y
			first = false
		} else {
			if xmin > elf.x {
				xmin = elf.x
			}
			if xmax < elf.x {
				xmax = elf.x
			}
			if ymin > elf.y {
				ymin = elf.y
			}
			if ymax < elf.y {
				ymax = elf.y
			}
		}
	}

	result := (xmax-xmin+1)*(ymax-ymin+1) - len(elves)
	fmt.Println(result)
}

func solve2() {
	lines := helper.ReadLines("input-23.txt")

	elves := make(map[Point]bool)

	for row := 0; row < len(lines); row++ {
		splitLine := strings.Split(lines[row], "")
		for column := 0; column < len(splitLine); column++ {
			if splitLine[column] == "#" {
				point := createPoint(column, row)
				elves[point] = true
			}
		}
	}

	directions := [4]string{"N", "S", "W", "E"}
	done := false
	round := 1

	for !done {
		proposedMovements := make(map[Point]Point)
		duplicateProtection := make(map[Point]int)
		moved := false

		for elf, _ := range elves {
			neighbours := getNeighbouringElves(elves, elf)
			if len(neighbours) == 0 {
				proposedMovements[elf] = elf
				duplicateProtection[elf] = duplicateProtection[elf] + 1
			} else {
				point := getProposedMovement(neighbours, elf, directions)
				proposedMovements[elf] = point
				duplicateProtection[point] = duplicateProtection[point] + 1
			}
		}

		newElves := make(map[Point]bool)
		for origin, target := range proposedMovements {
			//Check if not bumping
			if duplicateProtection[target] >= 2 {
				newElves[origin] = true
			} else {
				newElves[target] = true
				if target != origin {
					moved = true
				}
			}
		}

		elves = newElves
		//shift directions
		var newDirections [4]string
		newDirections[0] = directions[1]
		newDirections[1] = directions[2]
		newDirections[2] = directions[3]
		newDirections[3] = directions[0]
		directions = newDirections

		if !moved {
			done = true
		} else {
			round++
		}
	}

	fmt.Println(round)
}

func getNeighbouringElves(elves map[Point]bool, elf Point) map[Point]bool {
	neighbours := make(map[Point]bool)
	for x := elf.x - 1; x <= elf.x+1; x++ {
		for y := elf.y - 1; y <= elf.y+1; y++ {
			if !(x == elf.x && y == elf.y) {
				point := createPoint(x, y)
				_, exists := elves[point]
				if exists {
					neighbours[point] = true
				}
			}
		}
	}
	return neighbours
}

func getProposedMovement(neighbours map[Point]bool, elf Point, directions [4]string) Point {
	for i := 0; i < 4; i++ {
		point, valid := checkDirection(neighbours, elf, directions[i])
		if valid {
			return point
		}
	}
	return elf
}

func checkDirection(neighbours map[Point]bool, elf Point, direction string) (Point, bool) {
	if direction == "N" {
		nw := createPoint(elf.x-1, elf.y-1)
		n := createPoint(elf.x, elf.y-1)
		ne := createPoint(elf.x+1, elf.y-1)
		if neighbours[nw] == false && neighbours[n] == false && neighbours[ne] == false {
			return n, true
		}
	}
	if direction == "S" {
		sw := createPoint(elf.x-1, elf.y+1)
		s := createPoint(elf.x, elf.y+1)
		se := createPoint(elf.x+1, elf.y+1)
		if neighbours[sw] == false && neighbours[s] == false && neighbours[se] == false {
			return s, true
		}
	}
	if direction == "W" {
		nw := createPoint(elf.x-1, elf.y-1)
		w := createPoint(elf.x-1, elf.y)
		sw := createPoint(elf.x-1, elf.y+1)
		if neighbours[nw] == false && neighbours[w] == false && neighbours[sw] == false {
			return w, true
		}
	}
	if direction == "E" {
		ne := createPoint(elf.x+1, elf.y-1)
		e := createPoint(elf.x+1, elf.y)
		se := createPoint(elf.x+1, elf.y+1)
		if neighbours[ne] == false && neighbours[e] == false && neighbours[se] == false {
			return e, true
		}
	}
	return Point{}, false
}

func printElves(elves map[Point]bool) {
	//Find boundaries
	xmin := 0
	xmax := 0
	ymin := 0
	ymax := 0
	first := true
	for elf, _ := range elves {
		if first {
			xmin = elf.x
			xmax = elf.x
			ymin = elf.y
			ymax = elf.y
			first = false
		} else {
			if xmin > elf.x {
				xmin = elf.x
			}
			if xmax < elf.x {
				xmax = elf.x
			}
			if ymin > elf.y {
				ymin = elf.y
			}
			if ymax < elf.y {
				ymax = elf.y
			}
		}
	}

	//Printing
	for row := ymin; row <= ymax; row++ {
		for column := xmin; column <= xmax; column++ {
			point := createPoint(column, row)
			if elves[point] == true {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
