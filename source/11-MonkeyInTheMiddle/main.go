package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/LenaBullens/advent-of-code-2022-go/source/helper"
)

type Monkey struct {
	id              int
	items           []int
	operationType   string
	operationAmount int
	testFactor      int
	positiveMonkey  int
	negativeMonkey  int
}

func createMonkey(id int, items []int, operationType string, operationAmount int, testFactor int, positiveMonkey int, negativeMonkey int) Monkey {
	return Monkey{id: id, items: items, operationType: operationType, operationAmount: operationAmount, testFactor: testFactor, positiveMonkey: positiveMonkey, negativeMonkey: negativeMonkey}
}

func main() {
	solve2()
}

func solve1() {
	lines := helper.ReadLines("input-11.txt")

	monkeys := createMonkeys(lines)

	inspections := make([]int, len(monkeys))

	for n := 0; n < 20; n++ {
		for i := 0; i < len(monkeys); i++ {
			currentMonkey := monkeys[i]
			currentItems := currentMonkey.items
			var newItems []int
			for _, item := range currentItems {
				//Inspect
				inspections[i] = inspections[i] + 1

				//Worry
				if currentMonkey.operationType == "square" {
					item = item * item
				} else if currentMonkey.operationType == "multiply" {
					item = item * currentMonkey.operationAmount
				} else if currentMonkey.operationType == "add" {
					item = item + currentMonkey.operationAmount
				}

				//Calm down
				item = item / 3

				//Decide which monkey to throw to
				testFactor := currentMonkey.testFactor
				if item%testFactor == 0 {
					newMonkeyId := currentMonkey.positiveMonkey
					receivingMonkey := monkeys[newMonkeyId]
					receivingMonkey.items = append(receivingMonkey.items, item)
					monkeys[newMonkeyId] = receivingMonkey
				} else {
					newMonkeyId := currentMonkey.negativeMonkey
					receivingMonkey := monkeys[newMonkeyId]
					receivingMonkey.items = append(receivingMonkey.items, item)
					monkeys[newMonkeyId] = receivingMonkey
				}
			}
			currentMonkey.items = newItems
			monkeys[i] = currentMonkey
		}
	}

	sort.Ints(inspections)

	top := inspections[len(inspections)-1]
	second := inspections[len(inspections)-2]
	monkeyBusiness := top * second

	fmt.Println(monkeyBusiness)
}

func solve2() {
	lines := helper.ReadLines("input-11.txt")

	monkeys := createMonkeys(lines)

	inspections := make([]int, len(monkeys))

	magicNumber := 1
	for _, monkey := range monkeys {
		magicNumber = magicNumber * monkey.testFactor
	}

	for n := 0; n < 10000; n++ {
		for i := 0; i < len(monkeys); i++ {
			currentMonkey := monkeys[i]
			currentItems := currentMonkey.items
			var newItems []int
			for _, item := range currentItems {
				//Inspect
				inspections[i] = inspections[i] + 1

				//Worry
				if currentMonkey.operationType == "square" {
					item = item * item
				} else if currentMonkey.operationType == "multiply" {
					item = item * currentMonkey.operationAmount
				} else if currentMonkey.operationType == "add" {
					item = item + currentMonkey.operationAmount
				}

				//Reduce worry through ~magic~
				item = item % magicNumber

				//Decide which monkey to throw to
				testFactor := currentMonkey.testFactor
				if item%testFactor == 0 {
					newMonkeyId := currentMonkey.positiveMonkey
					receivingMonkey := monkeys[newMonkeyId]
					receivingMonkey.items = append(receivingMonkey.items, item)
					monkeys[newMonkeyId] = receivingMonkey
				} else {
					newMonkeyId := currentMonkey.negativeMonkey
					receivingMonkey := monkeys[newMonkeyId]
					receivingMonkey.items = append(receivingMonkey.items, item)
					monkeys[newMonkeyId] = receivingMonkey
				}
			}
			currentMonkey.items = newItems
			monkeys[i] = currentMonkey
		}
	}

	sort.Ints(inspections)

	top := inspections[len(inspections)-1]
	second := inspections[len(inspections)-2]
	monkeyBusiness := top * second

	fmt.Println(monkeyBusiness)
}

func createMonkeys(input []string) []Monkey {
	lineNb := 0
	var monkeys []Monkey

	for lineNb < len(input) {
		//Id
		id := getId(input[lineNb])
		lineNb++

		//Items
		items := getItems(input[lineNb])
		lineNb++

		//Operation
		operationType := getOperationType(input[lineNb])
		operationAmount := getOperationAmount(input[lineNb])
		lineNb++

		//Test factor
		testFactor := getTestFactor(input[lineNb])
		lineNb++

		//Test passes
		positiveMonkey := getPositiveMonkey(input[lineNb])
		lineNb++

		//Test fails
		negativeMonkey := getNegativeMonkey(input[lineNb])
		lineNb++

		//Empty line
		lineNb++

		monkey := createMonkey(id, items, operationType, operationAmount, testFactor, positiveMonkey, negativeMonkey)
		monkeys = append(monkeys, monkey)
	}

	return monkeys
}

func getId(input string) int {
	idSection := strings.Split(input, " ")[1]
	id, error := strconv.Atoi(idSection[:len(idSection)-1])
	if error != nil {
		log.Fatal(error)
	}
	return id
}

func getItems(input string) []int {
	itemSection := input[18:]
	splitItems := strings.Split(itemSection, ", ")
	var result []int
	for _, item := range splitItems {
		itemNb, error := strconv.Atoi(item)
		if error != nil {
			log.Fatal(error)
		}
		result = append(result, itemNb)
	}
	return result
}

func getOperationType(input string) string {
	operationTypeSection := input[23:]
	if strings.HasPrefix(operationTypeSection, "* old") {
		return "square"
	} else if strings.HasPrefix(operationTypeSection, "*") {
		return "multiply"
	} else {
		return "add"
	}
}

func getOperationAmount(input string) int {
	operationAmountSection := input[25:]
	if operationAmountSection == "old" {
		return 0
	}
	operationAmount, error := strconv.Atoi(operationAmountSection)
	if error != nil {
		log.Fatal(error)
	}
	return operationAmount
}

func getTestFactor(input string) int {
	testFactorSection := input[21:]
	testFactor, error := strconv.Atoi(testFactorSection)
	if error != nil {
		log.Fatal(error)
	}
	return testFactor
}

func getPositiveMonkey(input string) int {
	positiveMonkeySection := input[29:]
	positiveMonkey, error := strconv.Atoi(positiveMonkeySection)
	if error != nil {
		log.Fatal(error)
	}
	return positiveMonkey
}

func getNegativeMonkey(input string) int {
	negativeMonkeySection := input[30:]
	negativeMonkey, error := strconv.Atoi(negativeMonkeySection)
	if error != nil {
		log.Fatal(error)
	}
	return negativeMonkey
}
