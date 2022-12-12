package main

import (
	"container/heap"
	"fmt"
	"log"

	"github.com/LenaBullens/advent-of-code-2022-go/source/helper"
)

const MAXINT = int(^uint(0) >> 1)

type Point struct {
	row    int
	column int
}

func createPoint(row int, column int) Point {
	return Point{row: row, column: column}

}

type Item struct {
	value    Point
	priority int
	index    int
	visited  bool
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, value Point, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func main() {
	solve1()
	solve2()
}

func solve1() {
	grid := convertToRuneGrid(helper.ReadGrid("input-12.txt"))

	height := len(grid)
	width := len(grid[0])

	var start *Point
	var end *Point

	tentativeDistances := make(map[*Item]int, height*width)
	priorityQueue := make(PriorityQueue, height*width)
	tiles := make(map[Point]*Item, height*width)

	i := 0
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			if grid[r][c] == 'S' {
				p := createPoint(r, c)
				item := Item{value: p, priority: 0, index: i, visited: false}
				tentativeDistances[&item] = 0
				priorityQueue[i] = &item
				tiles[p] = &item
				start = &p
			} else {
				p := createPoint(r, c)
				item := Item{value: p, priority: MAXINT, index: i, visited: false}
				tentativeDistances[&item] = MAXINT
				priorityQueue[i] = &item
				tiles[p] = &item
				if grid[r][c] == 'E' {
					end = &p
				}
			}
			i = i + 1
		}
	}

	if start == nil || end == nil {
		log.Fatal()
	}

	heap.Init(&priorityQueue)

	for priorityQueue.Len() > 0 {
		nextNode := heap.Pop(&priorityQueue).(*Item)
		unvisitedNeighbours := findUnvisitedNeighbours(nextNode, tiles, height, width)
		for i := 0; i < len(unvisitedNeighbours); i++ {
			row := unvisitedNeighbours[i].value.row
			column := unvisitedNeighbours[i].value.column
			if reachable1(grid[nextNode.value.row][nextNode.value.column], grid[row][column]) {
				newDistance := tentativeDistances[nextNode] + 1
				if newDistance < tentativeDistances[unvisitedNeighbours[i]] {
					tentativeDistances[unvisitedNeighbours[i]] = newDistance
					priorityQueue.update(unvisitedNeighbours[i], unvisitedNeighbours[i].value, newDistance)
				}
			}
		}
		nextNode.visited = true
	}

	endItem := tiles[*end]
	result := tentativeDistances[endItem]
	fmt.Printf("Result: %d\n", result)
}

func solve2() {
	grid := convertToRuneGrid(helper.ReadGrid("input-12.txt"))

	height := len(grid)
	width := len(grid[0])

	var start *Point

	tentativeDistances := make(map[*Item]int, height*width)
	priorityQueue := make(PriorityQueue, height*width)
	tiles := make(map[Point]*Item, height*width)

	i := 0
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			if grid[r][c] == 'E' {
				p := createPoint(r, c)
				item := Item{value: p, priority: 0, index: i, visited: false}
				tentativeDistances[&item] = 0
				priorityQueue[i] = &item
				tiles[p] = &item
				start = &p
			} else {
				p := createPoint(r, c)
				item := Item{value: p, priority: MAXINT, index: i, visited: false}
				tentativeDistances[&item] = MAXINT
				priorityQueue[i] = &item
				tiles[p] = &item
			}
			i = i + 1
		}
	}

	if start == nil {
		log.Fatal()
	}

	heap.Init(&priorityQueue)

	for priorityQueue.Len() > 0 {
		nextNode := heap.Pop(&priorityQueue).(*Item)
		if grid[nextNode.value.row][nextNode.value.column] == 'a' || grid[nextNode.value.row][nextNode.value.column] == 'S' {
			result := tentativeDistances[nextNode]
			fmt.Printf("Result: %d\n", result)
			break
		}
		unvisitedNeighbours := findUnvisitedNeighbours(nextNode, tiles, height, width)
		for i := 0; i < len(unvisitedNeighbours); i++ {
			row := unvisitedNeighbours[i].value.row
			column := unvisitedNeighbours[i].value.column
			if reachable2(grid[nextNode.value.row][nextNode.value.column], grid[row][column]) {
				newDistance := tentativeDistances[nextNode] + 1
				if newDistance < tentativeDistances[unvisitedNeighbours[i]] {
					tentativeDistances[unvisitedNeighbours[i]] = newDistance
					priorityQueue.update(unvisitedNeighbours[i], unvisitedNeighbours[i].value, newDistance)
				}
			}
		}
		nextNode.visited = true
	}
}

func findUnvisitedNeighbours(item *Item, tiles map[Point]*Item, height int, width int) []*Item {
	var neighbours []*Item
	//Left neighbour
	if item.value.column > 0 {
		neighbour := createPoint(item.value.row, item.value.column-1)
		_, exists := tiles[neighbour]
		if exists && !tiles[neighbour].visited {
			neighbours = append(neighbours, tiles[neighbour])
		}
	}
	//Top neighbour
	if item.value.row > 0 {
		neighbour := createPoint(item.value.row-1, item.value.column)
		_, exists := tiles[neighbour]
		if exists && !tiles[neighbour].visited {
			neighbours = append(neighbours, tiles[neighbour])
		}
	}
	//Right neighbour
	if item.value.column < width {
		neighbour := createPoint(item.value.row, item.value.column+1)
		_, exists := tiles[neighbour]
		if exists && !tiles[neighbour].visited {
			neighbours = append(neighbours, tiles[neighbour])
		}
	}
	//Bottom neighbour
	if item.value.row < height {
		neighbour := createPoint(item.value.row+1, item.value.column)
		_, exists := tiles[neighbour]
		if exists && !tiles[neighbour].visited {
			neighbours = append(neighbours, tiles[neighbour])
		}
	}
	return neighbours
}

func reachable1(start rune, end rune) bool {
	if start == 'S' {
		start = 'a'
	}
	if end == 'E' {
		end = 'z'
	}
	return end-start <= 1
}

func reachable2(start rune, end rune) bool {
	if start == 'E' {
		start = 'z'
	}
	if end == 'S' {
		end = 'a'
	}
	return end-start >= -1
}

func convertToRuneGrid(input [][]string) [][]rune {
	var result [][]rune
	for r := 0; r < len(input); r++ {
		var row []rune
		for c := 0; c < len(input[r]); c++ {
			row = append(row, stringToRune(input[r][c]))
		}
		result = append(result, row)
	}
	return result
}

func stringToRune(input string) rune {
	return []rune(input)[0]
}
