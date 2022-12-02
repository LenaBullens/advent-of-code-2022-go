package main

import (
	"fmt"

	"github.com/LenaBullens/advent-of-code-2022-go/source/helper"
)

var stateToScore1 = map[string]int{
	"A X": 4,
	"A Y": 8,
	"A Z": 3,
	"B X": 1,
	"B Y": 5,
	"B Z": 9,
	"C X": 7,
	"C Y": 2,
	"C Z": 6,
}

var stateToScore2 = map[string]int{
	"A X": 3,
	"A Y": 4,
	"A Z": 8,
	"B X": 1,
	"B Y": 5,
	"B Z": 9,
	"C X": 2,
	"C Y": 6,
	"C Z": 7,
}

func main() {
	part1()
	part2()
}

func part1() {
	lines := helper.ReadLines("input-02.txt")
	totalScore := 0

	for _, str := range lines {
		totalScore = totalScore + stateToScore1[str]
	}

	fmt.Println(totalScore)
}

func part2() {
	lines := helper.ReadLines("input-02.txt")
	totalScore := 0

	for _, str := range lines {
		totalScore = totalScore + stateToScore2[str]
	}

	fmt.Println(totalScore)
}
