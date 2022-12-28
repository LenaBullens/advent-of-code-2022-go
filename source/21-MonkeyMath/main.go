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
	parent      string
	leftMonkey  string
	rightMonkey string
	monkeyMap   map[string]Monkey
}

func createMonkey(name string, value int, operation Operation, leftMonkey string, rightMonkey string, monkeyMap map[string]Monkey) Monkey {
	return Monkey{name: name, value: value, operation: operation, leftMonkey: leftMonkey, rightMonkey: rightMonkey, monkeyMap: monkeyMap}
}

func (monkey Monkey) getParent() Monkey {
	return monkey.monkeyMap[monkey.parent]
}

func (monkey Monkey) getLeftMonkey() Monkey {
	return monkey.monkeyMap[monkey.leftMonkey]
}

func (monkey Monkey) getRightMonkey() Monkey {
	return monkey.monkeyMap[monkey.rightMonkey]
}

func (monkey Monkey) getOtherChild(child Monkey) Monkey {
	if child.name == monkey.leftMonkey {
		return monkey.getRightMonkey()
	} else {
		return monkey.getLeftMonkey()
	}
}

func (monkey Monkey) fixParents() {
	if monkey.leftMonkey != "" {
		leftMonkey := monkey.monkeyMap[monkey.leftMonkey]
		leftMonkey.parent = monkey.name
		monkey.monkeyMap[monkey.leftMonkey] = leftMonkey
	}
	if monkey.rightMonkey != "" {
		rightMonkey := monkey.monkeyMap[monkey.rightMonkey]
		rightMonkey.parent = monkey.name
		monkey.monkeyMap[monkey.rightMonkey] = rightMonkey
	}
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

func (monkey Monkey) getRequiredValue() int {
	// If we're already at the level below root, we just need to equate the value of the
	// other branch of the tree.
	if monkey.parent == "root" {
		return monkey.getParent().getOtherChild(monkey).getValue()
	} else {
		// Else we go up one level, and 'revert' the operation of the parent.
		if monkey.name == monkey.getParent().leftMonkey {
			if monkey.getParent().operation == Add {
				return monkey.getParent().getRequiredValue() - monkey.getParent().getOtherChild(monkey).getValue()
			}
			if monkey.getParent().operation == Subtract {
				return monkey.getParent().getRequiredValue() + monkey.getParent().getOtherChild(monkey).getValue()
			}
			if monkey.getParent().operation == Multiply {
				return monkey.getParent().getRequiredValue() / monkey.getParent().getOtherChild(monkey).getValue()
			}
			if monkey.getParent().operation == Divide {
				return monkey.getParent().getRequiredValue() * monkey.getParent().getOtherChild(monkey).getValue()
			}
		} else {
			if monkey.getParent().operation == Add {
				return monkey.getParent().getRequiredValue() - monkey.getParent().getOtherChild(monkey).getValue()
			}
			if monkey.getParent().operation == Subtract {
				return monkey.getParent().getOtherChild(monkey).getValue() - monkey.getParent().getRequiredValue()
			}
			if monkey.getParent().operation == Multiply {
				return monkey.getParent().getRequiredValue() / monkey.getParent().getOtherChild(monkey).getValue()
			}
			if monkey.getParent().operation == Divide {
				return monkey.getParent().getOtherChild(monkey).getValue() / monkey.getParent().getRequiredValue()
			}
		}
	}
	return 0
}

func main() {
	solve2()
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

func solve2() {
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

	for _, v := range monkeyMap {
		v.fixParents()
	}

	fmt.Println(monkeyMap["humn"].getRequiredValue())
}
