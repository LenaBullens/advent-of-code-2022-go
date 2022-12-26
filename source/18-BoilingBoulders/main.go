package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/LenaBullens/advent-of-code-2022-go/source/helper"
)

type Cube struct {
	x int
	y int
	z int
}

func createCube(x int, y int, z int) Cube {
	return Cube{x: x, y: y, z: z}
}

func main() {
	solve1()
}

func solve1() {
	var cubes []Cube

	lines := helper.ReadLines("input-18.txt")
	for i := 0; i < len(lines); i++ {
		splitLine := strings.Split(lines[i], ",")
		x, error := strconv.Atoi(splitLine[0])
		if error != nil {
			log.Fatal(error)
		}
		y, error := strconv.Atoi(splitLine[1])
		if error != nil {
			log.Fatal(error)
		}
		z, error := strconv.Atoi(splitLine[2])
		if error != nil {
			log.Fatal(error)
		}
		cubes = append(cubes, createCube(x, y, z))
	}

	var nbOfAdjacentPairs int

	// Make sure we only check each pair once. If we've checked (1,2) already we don't need to check (2,1) anymore.
	for i := 0; i < len(cubes); i++ {
		for j := i + 1; j < len(cubes); j++ {
			if areAdjacent(cubes[i], cubes[j]) {
				nbOfAdjacentPairs++
			}
		}
	}

	result := 6*len(cubes) - 2*nbOfAdjacentPairs
	fmt.Println(result)
}

func areAdjacent(cube1 Cube, cube2 Cube) bool {
	// Cubes are adjacent if their Manhattan distance is 1.
	manhattanDistance := absoluteValue(cube1.x-cube2.x) + absoluteValue(cube1.y-cube2.y) + absoluteValue(cube1.z-cube2.z)
	return manhattanDistance == 1
}

func absoluteValue(input int) int {
	if input > 0 {
		return input
	}
	return -1 * input
}
