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

type Blizzard struct {
	point     Point
	direction string
}

func createBlizzard(point Point, direction string) Blizzard {
	return Blizzard{point: point, direction: direction}
}

type Boundary struct {
	xmin int
	xmax int
	ymin int
	ymax int
}

func createBoundary(xmin int, xmax int, ymin int, ymax int) Boundary {
	return Boundary{xmin: xmin, xmax: xmax, ymin: ymin, ymax: ymax}
}

func main() {
	solve2()
}

func solve1() {
	blizzards := make(map[*Blizzard]bool)

	lines := helper.ReadLines("input-24.txt")
	for row := 0; row < len(lines); row++ {
		splitLine := strings.Split(lines[row], "")
		for column := 0; column < len(splitLine); column++ {
			point := createPoint(column, row)
			if strings.Contains("^v><", splitLine[column]) {
				blizzard := createBlizzard(point, splitLine[column])
				blizzards[&blizzard] = true
			}
		}
	}

	boundary := createBoundary(0, len(lines[0])-1, 0, len(lines)-1)
	start := createPoint(1, 0)
	destination := createPoint(len(lines[0])-2, len(lines)-1)

	fmt.Println(travelTo(boundary, blizzards, start, destination))
}

func solve2() {
	blizzards := make(map[*Blizzard]bool)

	lines := helper.ReadLines("input-24.txt")
	for row := 0; row < len(lines); row++ {
		splitLine := strings.Split(lines[row], "")
		for column := 0; column < len(splitLine); column++ {
			point := createPoint(column, row)
			if strings.Contains("^v><", splitLine[column]) {
				blizzard := createBlizzard(point, splitLine[column])
				blizzards[&blizzard] = true
			}
		}
	}

	boundary := createBoundary(0, len(lines[0])-1, 0, len(lines)-1)
	start := createPoint(1, 0)
	destination := createPoint(len(lines[0])-2, len(lines)-1)

	var result int

	result = travelTo(boundary, blizzards, start, destination)
	result = result + travelTo(boundary, blizzards, destination, start)
	result = result + travelTo(boundary, blizzards, start, destination)

	fmt.Println(result)
}

func travelTo(boundary Boundary, blizzards map[*Blizzard]bool, start Point, destination Point) int {
	if start == destination {
		return 0
	} else {
		possibleLocations := make(map[Point]bool)
		possibleLocations[start] = true

		var blizzardLocations map[Point]bool

		var minute int

		for !possibleLocations[destination] {

			//Advance time
			minute++

			//Move blizzards
			for blizz := range blizzards {
				if blizz.direction == ">" {
					if blizz.point.x >= boundary.xmax-1 {
						newLocation := createPoint(boundary.xmin+1, blizz.point.y)
						blizz.point = newLocation
					} else {
						newLocation := createPoint(blizz.point.x+1, blizz.point.y)
						blizz.point = newLocation
					}
				}
				if blizz.direction == "<" {
					if blizz.point.x <= boundary.xmin+1 {
						newLocation := createPoint(boundary.xmax-1, blizz.point.y)
						blizz.point = newLocation
					} else {
						newLocation := createPoint(blizz.point.x-1, blizz.point.y)
						blizz.point = newLocation
					}
				}
				if blizz.direction == "v" {
					if blizz.point.y >= boundary.ymax-1 {
						newLocation := createPoint(blizz.point.x, boundary.ymin+1)
						blizz.point = newLocation
					} else {
						newLocation := createPoint(blizz.point.x, blizz.point.y+1)
						blizz.point = newLocation
					}
				}
				if blizz.direction == "^" {
					if blizz.point.y <= boundary.ymin+1 {
						newLocation := createPoint(blizz.point.x, boundary.ymax-1)
						blizz.point = newLocation
					} else {
						newLocation := createPoint(blizz.point.x, blizz.point.y-1)
						blizz.point = newLocation
					}
				}
			}

			//Update blizzard locations
			blizzardLocations = make(map[Point]bool)
			for blizz := range blizzards {
				blizzardLocations[blizz.point] = true
			}

			//Evaluate possible locations
			newPossibleLocations := make(map[Point]bool)
			for loc := range possibleLocations {
				neighbours := getTraversableNeighbours(loc, start, destination, boundary, blizzardLocations)
				if len(neighbours) > 0 {
					for _, n := range neighbours {
						newPossibleLocations[n] = true
					}
				}
			}
			possibleLocations = newPossibleLocations
		}
		return minute
	}
}

func getTraversableNeighbours(location Point, start Point, destination Point, boundary Boundary, blizzardLocations map[Point]bool) []Point {
	var validNeighbours []Point

	//Staying stationary
	stationary := location
	if !blizzardLocations[stationary] {
		validNeighbours = append(validNeighbours, stationary)
	}

	//West
	westCandidate := createPoint(location.x-1, location.y)
	if isPointInValley(westCandidate, start, destination, boundary) && !blizzardLocations[westCandidate] {
		validNeighbours = append(validNeighbours, westCandidate)
	}

	//East
	eastCandidate := createPoint(location.x+1, location.y)
	if isPointInValley(eastCandidate, start, destination, boundary) && !blizzardLocations[eastCandidate] {
		validNeighbours = append(validNeighbours, eastCandidate)
	}

	//South
	southCandidate := createPoint(location.x, location.y+1)
	if isPointInValley(southCandidate, start, destination, boundary) && !blizzardLocations[southCandidate] {
		validNeighbours = append(validNeighbours, southCandidate)
	}

	//North
	northCandidate := createPoint(location.x, location.y-1)
	if isPointInValley(northCandidate, start, destination, boundary) && !blizzardLocations[northCandidate] {
		validNeighbours = append(validNeighbours, northCandidate)
	}

	return validNeighbours
}

func isPointInValley(location Point, start Point, destination Point, boundary Boundary) bool {
	if location == start {
		return true
	}
	if location == destination {
		return true
	}
	if location.x > boundary.xmin && location.x < boundary.xmax && location.y > boundary.ymin && location.y < boundary.ymax {
		return true
	}
	return false
}
