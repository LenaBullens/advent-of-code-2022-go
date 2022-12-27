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

type Boundary struct {
	minx int
	maxx int
	miny int
	maxy int
	minz int
	maxz int
}

func createBoundary(minx int, maxx int, miny int, maxy int, minz int, maxz int) Boundary {
	return Boundary{minx: minx, maxx: maxx, miny: miny, maxy: maxy, minz: minz, maxz: maxz}
}

type Material int

const (
	Air Material = iota
	Lava
)

func main() {
	solve1()
	solve2()
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

func solve2() {
	cubes := make(map[Cube]bool)
	var xmin int
	var xmax int
	var ymin int
	var ymax int
	var zmin int
	var zmax int

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
		if i == 0 {
			xmin = x
			xmax = x
			ymin = y
			ymax = y
			zmin = z
			zmax = z
		} else {
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
			if z < zmin {
				zmin = z
			}
			if z > zmax {
				zmax = z
			}
		}
		cubes[createCube(x, y, z)] = true
	}
	boundary := createBoundary(xmin-1, xmax+1, ymin-1, ymax+1, zmin-1, zmax+1)

	space := make(map[Cube]Material)
	visited := make(map[Cube]bool)

	for x := boundary.minx; x <= boundary.maxx; x++ {
		for y := boundary.miny; y <= boundary.maxy; y++ {
			for z := boundary.minz; z <= boundary.maxz; z++ {
				cube := createCube(x, y, z)
				_, exists := cubes[cube]
				if exists {
					space[cube] = Lava
				} else {
					space[cube] = Air
				}
			}
		}
	}

	//Floodfill the cube containing our droplet with water from the outside. Each surface of the droplet
	//is only in contact with one water cube, so when we encounter a lava cube, we count it as an encountered
	//surface.

	var surfacesEncountered int
	start := createCube(boundary.minx, boundary.miny, boundary.minz)
	visited[start] = true
	var queue []Cube
	queue = append(queue, start)
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		neighbours := getNeighbours(current, boundary)
		for _, neighbour := range neighbours {
			if space[neighbour] == Air {
				_, v := visited[neighbour]
				if !v {
					visited[neighbour] = true
					queue = append(queue, neighbour)
				}
			} else if space[neighbour] == Lava {
				surfacesEncountered++
			}
		}
	}

	fmt.Println(surfacesEncountered)
}

func areAdjacent(cube1 Cube, cube2 Cube) bool {
	// Cubes are adjacent if their Manhattan distance is 1.
	manhattanDistance := absoluteValue(cube1.x-cube2.x) + absoluteValue(cube1.y-cube2.y) + absoluteValue(cube1.z-cube2.z)
	return manhattanDistance == 1
}

func getNeighbours(cube Cube, boundary Boundary) []Cube {

	var neighbours []Cube
	if cube.x > boundary.minx {
		neighbours = append(neighbours, createCube(cube.x-1, cube.y, cube.z))
	}
	if cube.x < boundary.maxx {
		neighbours = append(neighbours, createCube(cube.x+1, cube.y, cube.z))
	}
	if cube.y > boundary.miny {
		neighbours = append(neighbours, createCube(cube.x, cube.y-1, cube.z))
	}
	if cube.y < boundary.maxy {
		neighbours = append(neighbours, createCube(cube.x, cube.y+1, cube.z))
	}
	if cube.z > boundary.minz {
		neighbours = append(neighbours, createCube(cube.x, cube.y, cube.z-1))
	}
	if cube.z < boundary.maxz {
		neighbours = append(neighbours, createCube(cube.x, cube.y, cube.z+1))
	}
	return neighbours
}

func absoluteValue(input int) int {
	if input > 0 {
		return input
	}
	return -1 * input
}
