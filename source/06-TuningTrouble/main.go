package main

import (
	"fmt"

	"github.com/LenaBullens/advent-of-code-2022-go/source/helper"
)

func main() {
	lines := helper.ReadLines("input-06.txt")
	result1 := solve(lines[0], 4)
	result2 := solve(lines[0], 14)
	fmt.Print("Result part 1: ")
	fmt.Println(result1)
	fmt.Print("Result part 2: ")
	fmt.Println(result2)
}

func solve(input string, width int) int {
	cursor := width
	for ; cursor < len(input); cursor++ {
		currentSection := input[cursor-width : cursor]
		if checkAllLettersUnique(currentSection) {
			return cursor
		}
	}
	return -1
}

func checkAllLettersUnique(input string) bool {
	set := make(map[rune]bool)
	for _, rune := range input {
		set[rune] = true
	}
	if len(set) == len(input) {
		return true
	}
	return false
}
