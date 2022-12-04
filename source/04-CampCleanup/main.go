package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/LenaBullens/advent-of-code-2022-go/source/helper"
)

type assignment struct {
	start int
	end   int
}

func newAssignment(start int, end int) *assignment {
	assignment := assignment{start, end}
	return &assignment
}

func main() {
	part2()
}

func part1() {
	lines := helper.ReadLines("input-04.txt")

	score := 0

	for _, line := range lines {
		assignmentStrings := strings.Split(line, ",")

		split1 := strings.Split(assignmentStrings[0], "-")
		start1, err := strconv.Atoi(split1[0])
		if err != nil {
			log.Fatal(err)
		}
		end1, err := strconv.Atoi(split1[1])
		if err != nil {
			log.Fatal(err)
		}
		assignment1 := newAssignment(start1, end1)

		split2 := strings.Split(assignmentStrings[1], "-")
		start2, err := strconv.Atoi(split2[0])
		if err != nil {
			log.Fatal(err)
		}
		end2, err := strconv.Atoi(split2[1])
		if err != nil {
			log.Fatal(err)
		}
		assignment2 := newAssignment(start2, end2)

		if (assignment1.start >= assignment2.start && assignment1.end <= assignment2.end) || (assignment2.start >= assignment1.start && assignment2.end <= assignment1.end) {
			score = score + 1
		}
	}

	fmt.Println(score)
}

func part2() {
	lines := helper.ReadLines("input-04.txt")

	score := 0

	for _, line := range lines {
		assignmentStrings := strings.Split(line, ",")

		split1 := strings.Split(assignmentStrings[0], "-")
		start1, err := strconv.Atoi(split1[0])
		if err != nil {
			log.Fatal(err)
		}
		end1, err := strconv.Atoi(split1[1])
		if err != nil {
			log.Fatal(err)
		}
		assignment1 := newAssignment(start1, end1)

		split2 := strings.Split(assignmentStrings[1], "-")
		start2, err := strconv.Atoi(split2[0])
		if err != nil {
			log.Fatal(err)
		}
		end2, err := strconv.Atoi(split2[1])
		if err != nil {
			log.Fatal(err)
		}
		assignment2 := newAssignment(start2, end2)

		if (assignment1.start >= assignment2.start && assignment1.start <= assignment2.end) || (assignment1.end >= assignment2.start && assignment1.end <= assignment2.end) {
			score = score + 1
		} else if (assignment2.start >= assignment1.start && assignment2.start <= assignment1.end) || (assignment2.end >= assignment1.start && assignment2.end <= assignment1.end) {
			score = score + 1
		}
	}

	fmt.Println(score)
}
