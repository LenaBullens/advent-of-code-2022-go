package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/LenaBullens/advent-of-code-2022-go/source/helper"
)

func main() {
	solve2()
}

func solve1() {
	lines := helper.ReadLines("input-13.txt")

	var leftPackets [][]any
	var rightPackets [][]any

	i := 0
	for i < len(lines) {
		//Skip empty lines
		for lines[i] == "" {
			i++
		}
		leftPackets = append(leftPackets, parsePacket(lines[i]))
		i++

		rightPackets = append(rightPackets, parsePacket(lines[i]))
		i++
	}

	var result int

	for j := 0; j < len(leftPackets); j++ {
		leftPacket := leftPackets[j]
		rightPacket := rightPackets[j]

		comparison := compareLists(leftPacket, rightPacket)

		if comparison == 1 {
			result = result + j + 1
		}
	}

	fmt.Println(result)
}

func solve2() {
	lines := helper.ReadLines("input-13.txt")

	var packets [][]any

	i := 0
	for i < len(lines) {
		//Skip empty lines
		for lines[i] == "" {
			i++
		}
		packets = append(packets, parsePacket(lines[i]))
		i++
	}

	twoPacket := parsePacket("[[2]]")
	sixPacket := parsePacket("[[6]]")

	indexTwo := 1
	indexSix := 2 // [[2]] always before [[6]]

	for j := 0; j < len(packets); j++ {
		if compareLists(packets[j], twoPacket) == 1 {
			indexTwo++
		}
	}

	for j := 0; j < len(packets); j++ {
		if compareLists(packets[j], sixPacket) == 1 {
			indexSix++
		}
	}

	result := indexTwo * indexSix

	fmt.Println(result)
}

func parsePacket(input string) []any {
	var result []any
	error := json.Unmarshal([]byte(input), &result)
	if error != nil {
		log.Fatal(error)
	}
	return result
}

func isNumber(input any) bool {
	_, ok := input.(float64)
	return ok
}

func numberToList(input float64) []any {
	var result []any
	result = append(result, input)
	return result
}

func compareNumbers(a float64, b float64) int {
	if a < b {
		return 1
	} else if a > b {
		return -1
	} else {
		return 0
	}
}

func compareLists(a []any, b []any) int {
	i := 0
	done := false
	for !done {
		if i < len(a) && i < len(b) {
			elementA := a[i]
			elementB := b[i]
			if isNumber(elementA) && isNumber(elementB) {
				comparison := compareNumbers(elementA.(float64), elementB.(float64))
				if comparison == 1 {
					return 1
				} else if comparison == -1 {
					return -1
				}
				i++
			} else {
				var elementAList []any
				var elementBList []any
				if isNumber(elementA) {
					elementAList = numberToList(elementA.(float64))
				} else {
					elementAList = elementA.([]any)
				}
				if isNumber(elementB) {
					elementBList = numberToList(elementB.(float64))
				} else {
					elementBList = elementB.([]any)
				}
				comparison := compareLists(elementAList, elementBList)
				if comparison == 1 {
					return 1
				} else if comparison == -1 {
					return -1
				}
				i++
			}
		} else if len(a) >= i && i < len(b) {
			return 1
		} else if i < len(a) && len(b) >= i {
			return -1
		} else {
			done = true
		}
	}

	return 0
}
