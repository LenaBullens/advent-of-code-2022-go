package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/LenaBullens/advent-of-code-2022-go/source/helper"
)

type Operation int

const (
	Invalid Operation = iota
	Add
	Subtract
	Multiply
	Divide
)

type Monkey struct {
	name        string
	value       int
	operation   Operation
	leftMonkey  string
	rightMonkey string
	monkeyMap   map[string]Monkey
}

func createMonkey(name string, value int, operation Operation, leftMonkey string, rightMonkey string, monkeyMap map[string]Monkey) Monkey {
	return Monkey{name: name, value: value, operation: operation, leftMonkey: leftMonkey, rightMonkey: rightMonkey, monkeyMap: monkeyMap}
}

func (monkey Monkey) getLeftMonkey() Monkey {
	return monkey.monkeyMap[monkey.leftMonkey]
}

func (monkey Monkey) getRightMonkey() Monkey {
	return monkey.monkeyMap[monkey.rightMonkey]
}

func (monkey Monkey) getValue() int {
	if monkey.operation == Invalid {
		return monkey.value
	} else {
		leftMonkey := monkey.getLeftMonkey()
		rightMonkey := monkey.getRightMonkey()
		if monkey.operation == Add {
			return leftMonkey.getValue() + rightMonkey.getValue()
		}
		if monkey.operation == Subtract {
			return leftMonkey.getValue() - rightMonkey.getValue()
		}
		if monkey.operation == Multiply {
			return leftMonkey.getValue() * rightMonkey.getValue()
		}
		if monkey.operation == Divide {
			return leftMonkey.getValue() / rightMonkey.getValue()
		}
	}
	return 0
}

func main() {
	solve1()
}

func solve1() {
	lines := helper.ReadLines("input-21.txt")
	monkeyMap := make(map[string]Monkey)

	for _, line := range lines {
		splitLine := strings.Split(line, ": ")
		name := splitLine[0]
		var operation Operation
		var value int
		var leftMonkey string
		var rightMonkey string
		if strings.Contains(splitLine[1], "+") {
			operation = Add
			operationSplit := strings.Split(splitLine[1], " + ")
			leftMonkey = operationSplit[0]
			rightMonkey = operationSplit[1]
		} else if strings.Contains(splitLine[1], "-") {
			operation = Subtract
			operationSplit := strings.Split(splitLine[1], " - ")
			leftMonkey = operationSplit[0]
			rightMonkey = operationSplit[1]
		} else if strings.Contains(splitLine[1], "*") {
			operation = Multiply
			operationSplit := strings.Split(splitLine[1], " * ")
			leftMonkey = operationSplit[0]
			rightMonkey = operationSplit[1]
		} else if strings.Contains(splitLine[1], "/") {
			operation = Divide
			operationSplit := strings.Split(splitLine[1], " / ")
			leftMonkey = operationSplit[0]
			rightMonkey = operationSplit[1]
		} else {
			value, _ = strconv.Atoi(splitLine[1])
			operation = Invalid
			leftMonkey = ""
			rightMonkey = ""
		}

		monkey := createMonkey(name, value, operation, leftMonkey, rightMonkey, monkeyMap)
		monkeyMap[name] = monkey
	}

	fmt.Println(monkeyMap["root"].getValue())
}
