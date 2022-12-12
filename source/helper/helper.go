package helper

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ReadLines(path string) []string {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	var lines []string

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func ReadGrid(path string) [][]string {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	var grid [][]string

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		rowArray := strings.Split(scanner.Text(), "")
		grid = append(grid, rowArray)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return grid
}
