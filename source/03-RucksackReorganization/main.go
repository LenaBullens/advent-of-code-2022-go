package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/LenaBullens/advent-of-code-2022-go/source/helper"
)

func main() {
	part1()
	part2()
}

func part1() {
	lines := helper.ReadLines("input-03.txt")

	var totalPriority = 0

	for _, line := range lines {
		found := false
		middle := len(line) / 2
		firstCompartment := line[:middle]
		secondCompartment := line[middle:]
		for _, r1 := range firstCompartment {
			for _, r2 := range secondCompartment {
				if r1 == r2 && !found {
					priority, err := runeToPriority(r1)
					if err != nil {
						log.Fatal(err)
					}
					totalPriority = totalPriority + priority
					found = true
				}
			}
		}
	}
	fmt.Println(totalPriority)
}

func part2() {
	lines := helper.ReadLines("input-03.txt")

	var totalPriority = 0

	options := map[rune]bool{
		'a': true,
		'b': true,
		'c': true,
		'd': true,
		'e': true,
		'f': true,
		'g': true,
		'h': true,
		'i': true,
		'j': true,
		'k': true,
		'l': true,
		'm': true,
		'n': true,
		'o': true,
		'p': true,
		'q': true,
		'r': true,
		's': true,
		't': true,
		'u': true,
		'v': true,
		'w': true,
		'x': true,
		'y': true,
		'z': true,
		'A': true,
		'B': true,
		'C': true,
		'D': true,
		'E': true,
		'F': true,
		'G': true,
		'H': true,
		'I': true,
		'J': true,
		'K': true,
		'L': true,
		'M': true,
		'N': true,
		'O': true,
		'P': true,
		'Q': true,
		'R': true,
		'S': true,
		'T': true,
		'U': true,
		'V': true,
		'W': true,
		'X': true,
		'Y': true,
		'Z': true,
	}

	for i := 0; i < len(lines); i = i + 3 {
		elf1 := lines[i]
		elf2 := lines[i+1]
		elf3 := lines[i+2]

		//Check all the items that aren't carried by each elf. Each item not carried by an elf can't be the
		//badge. If we then intersect all the items not carried by at least one elf the one that is left will
		//be the badge.

		missing1 := copySet(options)
		missing2 := copySet(options)
		missing3 := copySet(options)

		for _, r := range elf1 {
			delete(missing1, r)
		}

		for _, r := range elf2 {
			delete(missing2, r)
		}

		for _, r := range elf3 {
			delete(missing3, r)
		}

		present := copySet(options)
		for k, _ := range missing1 {
			delete(present, k)
		}
		for k, _ := range missing2 {
			delete(present, k)
		}
		for k, _ := range missing3 {
			delete(present, k)
		}

		for k, _ := range present {
			priority, err := runeToPriority(k)
			if err != nil {
				log.Fatal(err)
			}
			totalPriority = totalPriority + priority
		}
	}

	fmt.Println(totalPriority)
}

func runeToPriority(r rune) (int, error) {
	if r >= 65 && r <= 90 {
		return int(r) - 38, nil
	}
	if r >= 97 && r <= 122 {
		return int(r) - 96, nil
	}
	return -1, errors.New("invalid rune")
}

func copySet(input map[rune]bool) map[rune]bool {
	output := make(map[rune]bool)
	for k, v := range input {
		output[k] = v
	}
	return output
}
