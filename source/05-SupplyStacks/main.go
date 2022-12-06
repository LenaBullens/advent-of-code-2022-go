package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/LenaBullens/advent-of-code-2022-go/source/helper"
)

type Stack []string

func (stack *Stack) IsEmpty() bool {
	return len(*stack) == 0
}

func (stack *Stack) Push(value string) {
	*stack = append(*stack, value)
}

func (stack *Stack) Pop() (string, error) {
	if stack.IsEmpty() {
		return "", errors.New("stack is empty")
	} else {
		index := len(*stack) - 1
		top := (*stack)[index]
		*stack = (*stack)[:index]
		return top, nil
	}
}

func main() {
	part1()
	part2()
}

func part1() {
	solve(true)
}

func part2() {
	solve(false)
}

func solve(oneByone bool) {
	lines := helper.ReadLines("input-05.txt")

	//Identify stacks section and instructions section
	var stackInput []string
	var instructionInput []string
	scanningStacks := true

	i := 0
	for ; i < len(lines) && scanningStacks; i++ {
		if len(strings.TrimSpace(lines[i])) != 0 {
			stackInput = append(stackInput, lines[i])
		} else {
			scanningStacks = false
		}
	}

	for ; i < len(lines); i++ {
		if len(strings.TrimSpace(lines[i])) != 0 {
			instructionInput = append(instructionInput, lines[i])
		}
	}

	//Determine how many stacks there are based on last line of stacks sections
	amountOfStacksString := strings.Fields(stackInput[len(stackInput)-1])

	amountOfStacks, err := strconv.Atoi(amountOfStacksString[len(amountOfStacksString)-1])
	if err != nil {
		log.Fatal(err)
	}

	stacks := make([]Stack, amountOfStacks)

	for j := len(stackInput) - 2; j >= 0; j-- {
		crateLine := stackInput[j]
		for k := 0; k < amountOfStacks; k++ {
			crate := crateLine[1+4*k : 2+4*k]
			if len(strings.TrimSpace(crate)) > 0 {
				stacks[k].Push(crate)
			}
		}
	}

	//Apply the instructions
	for j := 0; j < len(instructionInput); j++ {
		instructionLine := instructionInput[j]
		instructionFields := strings.Split(instructionLine, " ")
		amount, err := strconv.Atoi(instructionFields[1])
		if err != nil {
			log.Fatal(err)
		}
		source, err := strconv.Atoi(instructionFields[3])
		if err != nil {
			log.Fatal(err)
		}
		target, err := strconv.Atoi(instructionFields[5])
		if err != nil {
			log.Fatal(err)
		}

		if oneByone {
			for k := 0; k < amount; k++ {
				crateToMove, err := stacks[source-1].Pop()
				if err != nil {
					log.Fatal(err)
				}
				stacks[target-1].Push(crateToMove)
			}

		} else {
			if amount == 1 {
				crateToMove, err := stacks[source-1].Pop()
				if err != nil {
					log.Fatal(err)
				}
				stacks[target-1].Push(crateToMove)
			} else {
				var tempStack Stack
				for k := 0; k < amount; k++ {
					crateToMove, err := stacks[source-1].Pop()
					if err != nil {
						log.Fatal(err)
					}
					tempStack.Push(crateToMove)
				}
				for k := 0; k < amount; k++ {
					crateToMove, err := tempStack.Pop()
					if err != nil {
						log.Fatal(err)
					}
					stacks[target-1].Push(crateToMove)
				}
			}
		}
	}

	var result string

	for _, stack := range stacks {
		topCrate, err := stack.Pop()
		if err != nil {
			log.Fatal(err)
		}
		result = result + topCrate
	}

	fmt.Println(result)
}
