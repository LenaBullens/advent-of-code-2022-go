package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/LenaBullens/advent-of-code-2022-go/source/helper"
)

func main() {
	solve2()
}

func solve1() {
	lines := helper.ReadLines("input-08.txt")
	height := len(lines)
	width := len(lines[0])
	heightGrid := make([][]int, height)
	visibilityGrid := make([][]bool, height)

	//Assume all trees are initially invisible (excluding outer edge)
	for y := 0; y < height; y++ {
		line := lines[y]
		splitLine := strings.Split(line, "")
		heightRow := make([]int, width)
		visibilityRow := make([]bool, width)
		for x := 0; x < width; x++ {
			treeHeight, error := strconv.Atoi(splitLine[x])
			if error != nil {
				log.Fatal(error)
			}
			heightRow[x] = treeHeight
			if x == 0 {
				visibilityRow[x] = true
			} else if y == 0 {
				visibilityRow[x] = true
			} else if x == width-1 {
				visibilityRow[x] = true
			} else if y == height-1 {
				visibilityRow[x] = true
			} else {
				visibilityRow[x] = false
			}
		}
		heightGrid[y] = heightRow
		visibilityGrid[y] = visibilityRow
	}

	//Start scanning
	//Left to right
	for y := 1; y < height-1; y++ {
		highestTree := heightGrid[y][0]
		for x := 1; x < width-1; x++ {
			currentHeight := heightGrid[y][x]
			if currentHeight > highestTree {
				visibilityGrid[y][x] = true
				highestTree = currentHeight
			}
		}
	}

	//Top to bottom
	for x := 1; x < width-1; x++ {
		highestTree := heightGrid[0][x]
		for y := 1; y < height-1; y++ {
			currentHeight := heightGrid[y][x]
			if currentHeight > highestTree {
				visibilityGrid[y][x] = true
				highestTree = currentHeight
			}
		}
	}

	//Right to left
	for y := 1; y < height-1; y++ {
		highestTree := heightGrid[y][width-1]
		for x := width - 2; x > 0; x-- {
			currentHeight := heightGrid[y][x]
			if currentHeight > highestTree {
				visibilityGrid[y][x] = true
				highestTree = currentHeight
			}
		}
	}

	//Bottom to top
	for x := 1; x < width-1; x++ {
		highestTree := heightGrid[height-1][x]
		for y := height - 2; y > 0; y-- {
			currentHeight := heightGrid[y][x]
			if currentHeight > highestTree {
				visibilityGrid[y][x] = true
				highestTree = currentHeight
			}
		}
	}

	result := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if visibilityGrid[y][x] {
				result++
			}
		}
	}

	fmt.Println(result)
}

func solve2() {
	lines := helper.ReadLines("input-08.txt")
	height := len(lines)
	width := len(lines[0])
	heightGrid := make([][]int, height)
	scenicGrid := make([][]int, height)

	for y := 0; y < height; y++ {
		line := lines[y]
		splitLine := strings.Split(line, "")
		heightRow := make([]int, width)
		scenicRow := make([]int, width)
		for x := 0; x < width; x++ {
			treeHeight, error := strconv.Atoi(splitLine[x])
			if error != nil {
				log.Fatal(error)
			}
			heightRow[x] = treeHeight
			scenicRow[x] = 0
		}
		heightGrid[y] = heightRow
		scenicGrid[y] = scenicRow
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			currentHeight := heightGrid[y][x]
			//Check left

			if x == 0 || y == 0 || x == width-1 || y == height-1 {
				scenicGrid[y][x] = 0
			} else {
				//Check left
				leftDistance := 0
				for i := x - 1; i >= 0; i-- {
					leftDistance++
					if heightGrid[y][i] >= currentHeight {
						break
					}
				}
				//Check left
				rightDistance := 0
				for i := x + 1; i <= width-1; i++ {
					rightDistance++
					if heightGrid[y][i] >= currentHeight {
						break
					}
				}
				//Check top
				topDistance := 0
				for i := y - 1; i >= 0; i-- {
					topDistance++
					if heightGrid[i][x] >= currentHeight {
						break
					}
				}
				//Check bottom
				bottomDistance := 0
				for i := y + 1; i <= height-1; i++ {
					bottomDistance++
					if heightGrid[i][x] >= currentHeight {
						break
					}
				}
				scenicGrid[y][x] = leftDistance * rightDistance * topDistance * bottomDistance
			}
		}
	}

	result := 0

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if scenicGrid[y][x] > result {
				result = scenicGrid[y][x]
			}
		}
	}

	fmt.Println(result)
}
