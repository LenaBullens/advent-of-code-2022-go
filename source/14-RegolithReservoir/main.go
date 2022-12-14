package main

import (
	"fmt"
	"log"
	"strconv"
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
	solve2()
}

func solve1() {
	grid, ymax := setupGrid("input-14.txt")

	settledSands := 0

	fullyDone := false
	for !fullyDone {
		currentSand := createPoint(500, 0)

		done := false

		for !done {
			//If we fall beneath the lowest solid tile we're done
			if currentSand.y > ymax+5 {
				done = true
				fullyDone = true
				continue
			}

			//Check below
			checker := createPoint(currentSand.x, currentSand.y+1)
			_, exists := grid[checker]
			if !exists {
				currentSand = checker
				continue
			}

			//Check bottom left
			checker = createPoint(currentSand.x-1, currentSand.y+1)
			_, exists = grid[checker]
			if !exists {
				currentSand = checker
				continue
			}

			//Check bottom right
			checker = createPoint(currentSand.x+1, currentSand.y+1)
			_, exists = grid[checker]
			if !exists {
				currentSand = checker
				continue
			}

			//If we got here the sand settled
			grid[currentSand] = true
			done = true
			settledSands++
		}
	}

	fmt.Println(settledSands)
}

func solve2() {
	grid, ymax := setupGrid("input-14.txt")

	settledSands := 0

	fullyDone := false
	for !fullyDone {
		currentSand := createPoint(500, 0)

		done := false

		for !done {
			//If we at ymax + 1 we can't drop lower because of the floor, so we settle
			if currentSand.y == ymax+1 {
				grid[currentSand] = true
				done = true
				settledSands++
				continue
			}

			//Check below
			checker := createPoint(currentSand.x, currentSand.y+1)
			_, exists := grid[checker]
			if !exists {
				currentSand = checker
				continue
			}

			//Check bottom left
			checker = createPoint(currentSand.x-1, currentSand.y+1)
			_, exists = grid[checker]
			if !exists {
				currentSand = checker
				continue
			}

			//Check bottom right
			checker = createPoint(currentSand.x+1, currentSand.y+1)
			_, exists = grid[checker]
			if !exists {
				currentSand = checker
				continue
			}

			//If we got here the sand settled
			grid[currentSand] = true
			done = true
			settledSands++

			//Check that we're not currently blocking the input
			if currentSand.x == 500 && currentSand.y == 0 {
				fullyDone = true
			}
		}
	}

	fmt.Println(settledSands)
}

func setupGrid(input string) (map[Point]bool, int) {
	lines := helper.ReadLines(input)

	grid := make(map[Point]bool)

	xmin := 500
	xmax := 500
	ymin := 0
	ymax := 0

	for _, line := range lines {
		splitLine := strings.Split(line, " -> ")
		var pointsOnLine []Point
		for _, pointString := range splitLine {
			pointSplit := strings.Split(pointString, ",")
			x, error := strconv.Atoi(pointSplit[0])
			if error != nil {
				log.Fatal(error)
			}
			y, error := strconv.Atoi(pointSplit[1])
			if error != nil {
				log.Fatal(error)
			}

			if x < xmin {
				xmin = x
			}
			if x > xmax {
				xmax = x
			}
			if y < ymin {
				ymin = y
			}
			if y > ymax {
				ymax = y
			}

			p := createPoint(x, y)
			pointsOnLine = append(pointsOnLine, p)
		}

		for i := 0; i < len(pointsOnLine)-1; i++ {
			point1 := pointsOnLine[i]
			point2 := pointsOnLine[i+1]

			//Horizontal line
			if point1.y == point2.y {
				start := helper.Min(point1.x, point2.x)
				end := helper.Max(point1.x, point2.x)
				for j := start; j <= end; j++ {
					p := createPoint(j, point1.y)
					grid[p] = true
				}
			}
			//Vertical line
			if point1.x == point2.x {
				start := helper.Min(point1.y, point2.y)
				end := helper.Max(point1.y, point2.y)
				for j := start; j <= end; j++ {
					p := createPoint(point1.x, j)
					grid[p] = true
				}
			}
		}
	}

	return grid, ymax
}

func printGrid(grid map[Point]bool, xmin int, xmax int, ymin int, ymax int) {
	for i := ymin; i <= ymax; i++ {
		for j := xmin; j <= xmax; j++ {
			p := createPoint(j, i)
			_, exists := grid[p]
			if exists {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
