package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/LenaBullens/advent-of-code-2022-go/source/helper"
)

// We keep track of the original index to distinguish between identical values
type Entry struct {
	index int
	value int64
}

func createEntry(index int, value int64) Entry {
	return Entry{index: index, value: value}
}

func main() {
	solve2()
}

func solve1() {
	lines := helper.ReadLines("input-20.txt")

	var numbers []int64
	for _, line := range lines {
		number, error := strconv.ParseInt(strings.TrimSpace(line), 10, 64)
		if error != nil {
			log.Fatal(error)
		}
		numbers = append(numbers, number)
	}

	var mutableList []Entry
	for i := 0; i < len(numbers); i++ {
		mutableList = append(mutableList, createEntry(i, numbers[i]))
	}

	for i := 0; i < len(numbers); i++ {
		entry := createEntry(i, numbers[i])
		var currentIndex int64 = int64(getIndexOfEntry(mutableList, entry))
		mutableList = removeEntry(mutableList, entry)
		currentIndex = currentIndex + entry.value
		currentIndex = remainder(currentIndex, int64(len(numbers)-1))
		if currentIndex == 0 {
			mutableList = append(mutableList, entry)
		} else {
			mutableList = append(mutableList[:currentIndex+1], mutableList[currentIndex:]...)
			mutableList[currentIndex] = entry
		}
	}

	zeroEntry := createEntry(getIndexOf(numbers, 0), 0)
	zeroIndex := getIndexOfEntry(mutableList, zeroEntry)

	result := mutableList[(zeroIndex+1000)%len(numbers)].value + mutableList[(zeroIndex+2000)%len(numbers)].value + mutableList[(zeroIndex+3000)%len(numbers)].value
	fmt.Println(result)
}

func solve2() {
	lines := helper.ReadLines("input-20.txt")

	var numbers []int64
	for _, line := range lines {
		number, error := strconv.ParseInt(strings.TrimSpace(line), 10, 64)
		if error != nil {
			log.Fatal(error)
		}
		numbers = append(numbers, number*811589153)
	}

	var mutableList []Entry
	for i := 0; i < len(numbers); i++ {
		mutableList = append(mutableList, createEntry(i, numbers[i]))
	}

	for j := 0; j < 10; j++ {
		for i := 0; i < len(numbers); i++ {
			entry := createEntry(i, numbers[i])
			var currentIndex int64 = int64(getIndexOfEntry(mutableList, entry))
			mutableList = removeEntry(mutableList, entry)
			currentIndex = currentIndex + entry.value
			currentIndex = remainder(currentIndex, int64(len(numbers)-1))
			if currentIndex == 0 {
				mutableList = append(mutableList, entry)
			} else {
				mutableList = append(mutableList[:currentIndex+1], mutableList[currentIndex:]...)
				mutableList[currentIndex] = entry
			}
		}
	}

	zeroEntry := createEntry(getIndexOf(numbers, 0), 0)
	zeroIndex := getIndexOfEntry(mutableList, zeroEntry)

	result := mutableList[(zeroIndex+1000)%len(numbers)].value + mutableList[(zeroIndex+2000)%len(numbers)].value + mutableList[(zeroIndex+3000)%len(numbers)].value
	fmt.Println(result)
}

func remainder(index int64, divisor int64) int64 {
	remainder := index % divisor
	if remainder < 0 {
		remainder = remainder + divisor
	}
	return remainder
}

func getIndexOf(input []int64, value int64) int {
	for i, v := range input {
		if value == v {
			return i
		}
	}
	return -1
}

func getIndexOfEntry(input []Entry, entryToGetIndexOf Entry) int {
	for i, entry := range input {
		if entryToGetIndexOf == entry {
			return i
		}
	}
	return -1
}

func removeEntry(input []Entry, entryToRemove Entry) []Entry {
	index := -1
	for i, entry := range input {
		if entryToRemove == entry {
			index = i
			break
		}
	}
	result := make([]Entry, len(input))
	copy(result, input)
	if index != -1 {
		if index == 0 {
			return result[1:]
		} else if index == len(input)-1 {
			return result[:index]
		}
		result := append(result[:index], result[index+1:]...)
		return result
	}
	return result
}
