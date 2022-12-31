package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/LenaBullens/advent-of-code-2022-go/source/helper"
)

type Point struct {
	row    int
	column int
}

func createPoint(row int, column int) Point {
	return Point{row: row, column: column}
}

type Me struct {
	point  Point
	facing string
}

func createMe(point Point, facing string) Me {
	return Me{point: point, facing: facing}
}

func main() {
	solve1()
}

func solve1() {
	lines := helper.ReadLines("input-22.txt")

	leftMostPoints := make(map[int]int)
	rightMostPoints := make(map[int]int)
	topMostPoints := make(map[int]int)
	bottomMostPoints := make(map[int]int)

	tiles := make(map[Point]string)
	var me Me

	var steps []int
	var turns []string

	first := true
	scanningGrid := true
	for row, line := range lines {
		if scanningGrid {
			if len(line) > 0 {
				splitLine := strings.Split(line, "")
				for column, terrainType := range splitLine {
					if terrainType != " " {
						point := createPoint(row, column)
						if first {
							me = createMe(point, ">")
							first = false
						}
						tiles[point] = terrainType
						//Update boundaries

						//For left-most and top-most it's the first tile encountered since we parse from left to right and top to bottom.
						_, exists := leftMostPoints[row]
						if !exists {
							leftMostPoints[row] = column
						}
						_, exists = topMostPoints[column]
						if !exists {
							topMostPoints[column] = row
						}

						//For right-most and bottom-most it's the last tile we encounter.
						rightMostPoints[row] = column
						bottomMostPoints[column] = row
					}
				}
			} else {
				scanningGrid = false
			}
		} else {
			replaceLine1 := strings.ReplaceAll(line, "R", "L")
			splitLine1 := strings.Split(replaceLine1, "L")
			for i := 0; i < len(splitLine1); i++ {
				number, error := strconv.Atoi(splitLine1[i])
				if error != nil {
					log.Fatal(error)
				}
				steps = append(steps, number)
			}

			replaceLine2 := strings.ReplaceAll(line, "0", "")
			replaceLine2 = strings.ReplaceAll(replaceLine2, "1", "")
			replaceLine2 = strings.ReplaceAll(replaceLine2, "2", "")
			replaceLine2 = strings.ReplaceAll(replaceLine2, "3", "")
			replaceLine2 = strings.ReplaceAll(replaceLine2, "4", "")
			replaceLine2 = strings.ReplaceAll(replaceLine2, "5", "")
			replaceLine2 = strings.ReplaceAll(replaceLine2, "6", "")
			replaceLine2 = strings.ReplaceAll(replaceLine2, "7", "")
			replaceLine2 = strings.ReplaceAll(replaceLine2, "8", "")
			replaceLine2 = strings.ReplaceAll(replaceLine2, "9", "")

			splitLine2 := strings.Split(replaceLine2, "")
			for i := 0; i < len(splitLine2); i++ {
				turns = append(turns, splitLine2[i])
			}
		}
	}

	//Do the move
	done := false
	index := 0
	for !done {
		if index < len(steps) {
			currentSteps := steps[index]
			for i := 0; i < currentSteps; i++ {
				//Moving right
				if me.facing == ">" {
					candidate := createPoint(me.point.row, me.point.column+1)
					terrainType, exists := tiles[candidate]
					if exists {
						if terrainType == "." {
							me = createMe(candidate, me.facing)
						}
					} else {
						candidateColumn := leftMostPoints[me.point.row]
						candidate = createPoint(me.point.row, candidateColumn)
						terrainType, exists = tiles[candidate]
						if exists {
							if terrainType == "." {
								me = createMe(candidate, me.facing)
							}
						}

					}
				} else if me.facing == "v" {
					candidate := createPoint(me.point.row+1, me.point.column)
					terrainType, exists := tiles[candidate]
					if exists {
						if terrainType == "." {
							me = createMe(candidate, me.facing)
						}
					} else {
						candidateRow := topMostPoints[me.point.column]
						candidate = createPoint(candidateRow, me.point.column)
						terrainType, exists = tiles[candidate]
						if exists {
							if terrainType == "." {
								me = createMe(candidate, me.facing)
							}
						}

					}
				} else if me.facing == "<" {
					candidate := createPoint(me.point.row, me.point.column-1)
					terrainType, exists := tiles[candidate]
					if exists {
						if terrainType == "." {
							me = createMe(candidate, me.facing)
						}
					} else {
						candidateColumn := rightMostPoints[me.point.row]
						candidate = createPoint(me.point.row, candidateColumn)
						terrainType, exists = tiles[candidate]
						if exists {
							if terrainType == "." {
								me = createMe(candidate, me.facing)
							}
						}
					}
				} else if me.facing == "^" {
					candidate := createPoint(me.point.row-1, me.point.column)
					terrainType, exists := tiles[candidate]
					if exists {
						if terrainType == "." {
							me = createMe(candidate, me.facing)
						}
					} else {
						candidateRow := bottomMostPoints[me.point.column]
						candidate = createPoint(candidateRow, me.point.column)
						terrainType, exists = tiles[candidate]
						if exists {
							if terrainType == "." {
								me = createMe(candidate, me.facing)
							}
						}

					}
				}
			}
		}

		if index < len(turns) {
			currentTurn := turns[index]
			if currentTurn == "L" {
				me = createMe(me.point, turnLeft(me.facing))
			} else {
				me = createMe(me.point, turnRight(me.facing))
			}
		}

		if index >= len(steps) && index >= len(turns) {
			done = true
		} else {
			index++
		}
	}

	result := 1000*(me.point.row+1) + 4*(me.point.column+1) + facingToInt(me.facing)
	fmt.Println(result)
}

func turnRight(direction string) string {
	if direction == ">" {
		return "v"
	}
	if direction == "v" {
		return "<"
	}
	if direction == "<" {
		return "^"
	}
	if direction == "^" {
		return ">"
	}
	return ""
}

func turnLeft(direction string) string {
	if direction == ">" {
		return "^"
	}
	if direction == "^" {
		return "<"
	}
	if direction == "<" {
		return "v"
	}
	if direction == "v" {
		return ">"
	}
	return ""
}

func facingToInt(direction string) int {
	if direction == ">" {
		return 0
	}
	if direction == "v" {
		return 1
	}
	if direction == "<" {
		return 2
	}
	if direction == "^" {
		return 3
	}
	return 0
}
