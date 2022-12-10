package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/LenaBullens/advent-of-code-2022-go/source/helper"
)

func main() {
	solve1()
	solve2()
}

func solve1() {
	lines := helper.ReadLines("input-10.txt")

	x := 1
	cycle := 0
	row := 0
	done := false

	addInQueue := false
	amountInQueue := 0

	result := 0

	for !done {
		cycle = cycle + 1

		result = result + checkSignalStrength(cycle, x)

		if addInQueue {
			x = x + amountInQueue
			addInQueue = false
		} else {
			line := lines[row]
			row = row + 1
			if len(line) > 4 {
				amount, error := strconv.Atoi(strings.Split(line, " ")[1])
				if error != nil {
					log.Fatal(error)
				}
				addInQueue = true
				amountInQueue = amount
			}
		}
		done = !addInQueue && row >= len(lines)
	}
	fmt.Println(result)
}

func solve2() {
	lines := helper.ReadLines("input-10.txt")

	x := 1
	cycle := 0
	row := 0
	done := false

	addInQueue := false
	amountInQueue := 0

	for !done {
		cycle = cycle + 1

		draw(cycle, x)

		if addInQueue {
			x = x + amountInQueue
			addInQueue = false
		} else {
			line := lines[row]
			row = row + 1
			if len(line) > 4 {
				amount, error := strconv.Atoi(strings.Split(line, " ")[1])
				if error != nil {
					log.Fatal(error)
				}
				addInQueue = true
				amountInQueue = amount
			}
		}
		done = !addInQueue && row >= len(lines)
	}
}

func checkSignalStrength(cycle int, x int) int {
	if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
		return cycle * x
	}
	return 0
}

func draw(cycle int, x int) {
	horizontalPosition := (cycle - 1) % 40
	if horizontalPosition >= x-1 && horizontalPosition <= x+1 {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
	if horizontalPosition == 39 {
		fmt.Println()
	}
}
