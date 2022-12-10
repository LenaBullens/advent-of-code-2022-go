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
	part1()
	part2()
}

func part1() {
	lines := helper.ReadLines("input-09.txt")
	fmt.Println(solve1(lines))
}

func part2() {
	lines := helper.ReadLines("input-09.txt")
	fmt.Println(solve2(lines))
}

func solve1(input []string) int {
	head := createPoint(0, 0)
	tail := createPoint(0, 0)
	visitedPoints := make(map[Point]bool)

	instructions := convertInstructions(input)

	for _, instruction := range instructions {
		head = moveHead(head, instruction)
		tail = moveTail(tail, head)
		visitedPoints[tail] = true
	}

	return len(visitedPoints)
}

func solve2(input []string) int {
	head := createPoint(0, 0)
	tail1 := createPoint(0, 0)
	tail2 := createPoint(0, 0)
	tail3 := createPoint(0, 0)
	tail4 := createPoint(0, 0)
	tail5 := createPoint(0, 0)
	tail6 := createPoint(0, 0)
	tail7 := createPoint(0, 0)
	tail8 := createPoint(0, 0)
	tail9 := createPoint(0, 0)

	visitedPoints := make(map[Point]bool)

	instructions := convertInstructions(input)

	for _, instruction := range instructions {
		head = moveHead(head, instruction)
		tail1 = moveTail(tail1, head)
		tail2 = moveTail(tail2, tail1)
		tail3 = moveTail(tail3, tail2)
		tail4 = moveTail(tail4, tail3)
		tail5 = moveTail(tail5, tail4)
		tail6 = moveTail(tail6, tail5)
		tail7 = moveTail(tail7, tail6)
		tail8 = moveTail(tail8, tail7)
		tail9 = moveTail(tail9, tail8)

		visitedPoints[tail9] = true
	}

	return len(visitedPoints)
}

func convertInstructions(input []string) []string {
	var instructions []string
	for _, line := range input {
		instructionArray := strings.Split(line, " ")
		direction := instructionArray[0]
		distance, error := strconv.Atoi(instructionArray[1])
		if error != nil {
			log.Fatal(error)
		}
		for i := 0; i < distance; i++ {
			instructions = append(instructions, direction)
		}
	}
	return instructions
}

func moveHead(input Point, command string) Point {
	if "R" == command {
		return createPoint(input.x+1, input.y)
	} else if "D" == command {
		return createPoint(input.x, input.y+1)
	} else if "L" == command {
		return createPoint(input.x-1, input.y)
	} else if "U" == command {
		return createPoint(input.x, input.y-1)
	}
	return input
}

func moveTail(tail Point, head Point) Point {
	//Same row and two or more apart
	if tail.y == head.y {
		if tail.x <= head.x-2 {
			return createPoint(head.x-1, tail.y)
		} else if tail.x >= head.x+2 {
			return createPoint(head.x+1, tail.y)
		}
	}
	//Same column and two or more apart
	if tail.x == head.x {
		if tail.y <= head.y-2 {
			return createPoint(tail.x, head.y-1)
		} else if tail.y >= head.y+2 {
			return createPoint(tail.x, head.y+1)
		}
	}
	//Different row & column and not touching
	if tail.x != head.x && tail.y != head.y && !areAdjacent(tail, head) {
		if tail.x > head.x && tail.y > head.y {
			return createPoint(tail.x-1, tail.y-1)
		} else if tail.x < head.x && tail.y > head.y {
			return createPoint(tail.x+1, tail.y-1)
		} else if tail.x < head.x && tail.y < head.y {
			return createPoint(tail.x+1, tail.y+1)
		} else if tail.x > head.x && tail.y < head.y {
			return createPoint(tail.x-1, tail.y+1)
		}
	}
	return tail
}

func areAdjacent(p1 Point, p2 Point) bool {
	if p1.x >= p2.x-1 && p1.x <= p2.x+1 && p1.y >= p2.y-1 && p1.y <= p2.y+1 {
		return true
	}
	return false
}
