package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/LenaBullens/advent-of-code-2022-go/source/helper"
)

func main() {

	lines := helper.ReadLines("input-01.txt")
	var elves []int

	calorieTotal := 0
	for _, str := range lines {
		if strings.TrimSpace(str) != "" {
			nb, err := strconv.Atoi(str)
			if err != nil {
				log.Fatal(err)
			}
			calorieTotal = calorieTotal + nb
		} else {
			elves = append(elves, calorieTotal)
			calorieTotal = 0
		}
	}

	sort.Ints(elves)

	highest := elves[len(elves)-1]
	secondHighest := elves[len(elves)-2]
	thirdHighest := elves[len(elves)-3]

	fmt.Println("Highest calorie count is: " + strconv.Itoa(highest))

	fmt.Println("Sum of three highest calorie counts is: " + strconv.Itoa(highest+secondHighest+thirdHighest))
}
