package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/LenaBullens/advent-of-code-2022-go/source/helper"
)

type File struct {
	name   string
	size   int
	parent *Dir
}

type Dir struct {
	name   string
	parent *Dir
	dirs   []*Dir
	files  []*File
}

func (d Dir) Size() int {
	var size int
	for _, f := range d.files {
		size = size + f.size
	}
	for _, dir := range d.dirs {
		size = size + dir.Size()
	}
	return size
}

func (d Dir) getAllDirs() []Dir {
	dirsToReturn := make([]Dir, 0)
	for _, dir := range d.dirs {
		dirsToReturn = append(dirsToReturn, dir.getAllDirs()...)
	}
	dirsToReturn = append(dirsToReturn, d)
	return dirsToReturn
}

func main() {
	solve()
}

func solve() {
	lines := helper.ReadLines("input-07.txt")

	currentDir := &Dir{name: "/"}
	rootDir := currentDir

	for lineNb := 1; lineNb < len(lines); {
		if strings.HasPrefix(lines[lineNb], "$ ls") {
			lineNb++
			for lineNb < len(lines) && !strings.HasPrefix(lines[lineNb], "$") {
				if strings.HasPrefix(lines[lineNb], "dir") {
					dirName := lines[lineNb][4:]
					dir := Dir{name: dirName, parent: currentDir}
					currentDir.dirs = append(currentDir.dirs, &dir)
					lineNb++
				} else {
					fields := strings.Fields(lines[lineNb])
					fileSize, err := strconv.Atoi(fields[0])
					if err != nil {
						log.Fatal(err)
					}
					fileName := fields[1]
					file := File{name: fileName, size: fileSize, parent: currentDir}
					currentDir.files = append(currentDir.files, &file)
					lineNb++
				}
			}
		} else if strings.HasPrefix(lines[lineNb], "$ cd ..") {
			currentDir = currentDir.parent
			lineNb++
		} else if strings.HasPrefix(lines[lineNb], "$ cd") {
			dirName := lines[lineNb][5:]
			for _, d := range currentDir.dirs {
				if d.name == dirName {
					currentDir = d
				}
			}
			lineNb++
		}
	}

	r := *rootDir

	directories := r.getAllDirs()

	var result1 int

	for _, directory := range directories {
		size := directory.Size()
		if size <= 100000 {
			result1 = result1 + size
		}
	}

	fmt.Println(result1)

	currentSize := r.Size()
	unusedSize := 70000000 - currentSize
	amountToFreeUp := 30000000 - unusedSize

	var candidate Dir

	for _, directory := range directories {
		size := directory.Size()
		if size >= amountToFreeUp {
			if candidate.name == "" {
				candidate = directory
			} else if candidate.Size() > size {
				candidate = directory
			}
		}
	}

	fmt.Println(candidate.Size())
}
