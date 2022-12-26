package main

import (
	"fmt"
	"strconv"

	"github.com/LenaBullens/advent-of-code-2022-go/source/helper"
)

func main() {
	lines := helper.ReadLines("input-25.txt")
	var totalInDec int
	for _, line := range lines {
		totalInDec = totalInDec + snafuToDec(line)
	}
	fmt.Println(decToSnafu(totalInDec))
}

func snafuToDec(input string) int {
	factor := 1
	var result int
	for i := 0; i < len(input); i++ {
		current := input[len(input)-i-1 : len(input)-i]
		value := 0
		if current == "=" {
			value = -2
		} else if current == "-" {
			value = -1
		} else if current == "0" {
			value = 0
		} else if current == "1" {
			value = 1
		} else if current == "2" {
			value = 2
		}
		result = result + value*factor
		factor = factor * 5
	}
	return result
}

func decToSnafu(input int) string {
	factor := 1
	for input/factor > 0 {
		factor = factor * 5
	}
	factor = factor / 5

	var resultBase5 []int
	for factor >= 1 {
		current := input / factor
		resultBase5 = append(resultBase5, current)
		input = input % factor
		factor = factor / 5
	}

	var carry int
	var result string
	for i := len(resultBase5) - 1; i >= 0; i-- {
		current := strconv.Itoa(resultBase5[i] + carry)
		if current == "5" {
			carry = 1
			current = "0"
		} else if current == "4" {
			carry = 1
			current = "-"
		} else if current == "3" {
			carry = 1
			current = "="
		} else {
			carry = 0
		}
		result = current + result
	}
	if carry != 0 {
		result = strconv.Itoa(carry) + result
	}
	return result
}
