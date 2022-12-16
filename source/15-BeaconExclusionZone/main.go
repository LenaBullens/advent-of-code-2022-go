package main

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/LenaBullens/advent-of-code-2022-go/source/helper"
)

type Point struct {
	x int
	y int
}

func createPoint(x int, y int) Point {
	return Point{x: x, y: y}
}

type SBPair struct {
	sensor   Point
	beacon   Point
	distance int
}

func createSBPair(sensor Point, beacon Point) SBPair {
	distance := calculateDistance(sensor, beacon)
	return SBPair{sensor: sensor, beacon: beacon, distance: distance}
}

type Interval struct {
	start int
	end   int
}

func createInterval(start int, end int) Interval {
	return Interval{start: start, end: end}
}

func main() {
	solve1()
	solve2()
}

func solve1() {
	lines := helper.ReadLines("input-15.txt")

	var sbPairs []SBPair

	for _, line := range lines {
		splitLine := strings.Split(line, ":")
		sensorSection := splitLine[0]
		beaconSection := splitLine[1]

		sensorCoordinates := sensorSection[10:]
		sensorCoordinatesSplit := strings.Split(sensorCoordinates, ",")
		sensorX, error := strconv.Atoi(sensorCoordinatesSplit[0][2:])
		if error != nil {
			log.Fatal(error)
		}
		sensorY, error := strconv.Atoi(sensorCoordinatesSplit[1][3:])
		if error != nil {
			log.Fatal(error)
		}
		sensor := createPoint(sensorX, sensorY)

		beaconCoordinates := beaconSection[22:]
		beaconCoordinatesSplit := strings.Split(beaconCoordinates, ",")
		beaconX, error := strconv.Atoi(beaconCoordinatesSplit[0][2:])
		if error != nil {
			log.Fatal(error)
		}
		beaconY, error := strconv.Atoi(beaconCoordinatesSplit[1][3:])
		if error != nil {
			log.Fatal(error)
		}
		beacon := createPoint(beaconX, beaconY)

		sbPairs = append(sbPairs, createSBPair(sensor, beacon))
	}

	var sbPairsToCheck []SBPair

	target := 2000000

	for _, sbPair := range sbPairs {
		if absoluteValue(target-sbPair.sensor.y) <= sbPair.distance {
			sbPairsToCheck = append(sbPairsToCheck, sbPair)
		}
	}

	var intervals []Interval
	for _, sb := range sbPairsToCheck {
		interval, error := calculateBlockedTilesAtHeightForSensor(sb, target)
		if error != nil {
			log.Fatal(error)
		}
		intervals = append(intervals, interval)
	}

	intervals = mergeIntervals(intervals)

	totalPoints := 0
	for _, interval := range intervals {
		totalPoints = totalPoints + (interval.end - interval.start)
	}

	fmt.Println(totalPoints)
}

func solve2() {
	lines := helper.ReadLines("input-15.txt")

	var sbPairs []SBPair

	for _, line := range lines {
		splitLine := strings.Split(line, ":")
		sensorSection := splitLine[0]
		beaconSection := splitLine[1]

		sensorCoordinates := sensorSection[10:]
		sensorCoordinatesSplit := strings.Split(sensorCoordinates, ",")
		sensorX, error := strconv.Atoi(sensorCoordinatesSplit[0][2:])
		if error != nil {
			log.Fatal(error)
		}
		sensorY, error := strconv.Atoi(sensorCoordinatesSplit[1][3:])
		if error != nil {
			log.Fatal(error)
		}
		sensor := createPoint(sensorX, sensorY)

		beaconCoordinates := beaconSection[22:]
		beaconCoordinatesSplit := strings.Split(beaconCoordinates, ",")
		beaconX, error := strconv.Atoi(beaconCoordinatesSplit[0][2:])
		if error != nil {
			log.Fatal(error)
		}
		beaconY, error := strconv.Atoi(beaconCoordinatesSplit[1][3:])
		if error != nil {
			log.Fatal(error)
		}
		beacon := createPoint(beaconX, beaconY)

		sbPairs = append(sbPairs, createSBPair(sensor, beacon))
	}

	xmin := 0
	xmax := 4000000
	ymin := 0
	ymax := 4000000

	var result Point

	for i := ymin; i <= ymax; i++ {
		var intervals []Interval
		for _, sb := range sbPairs {
			interval, error := calculateBlockedTilesAtHeightForSensor2(sb, i, xmin, xmax)
			if error != nil {
				//Don't append
			} else {
				intervals = append(intervals, interval)
			}
		}

		intervals = mergeIntervals(intervals)
		if len(intervals) > 1 {
			//Found our y!
			y := i
			x := intervals[0].end + 1
			result = createPoint(x, y)
			break
		} else if intervals[0].start == xmin+1 {
			//Found our y
			y := i
			x := xmin
			result = createPoint(x, y)
			break
		} else if intervals[0].end == xmax-1 {
			//Found our y
			y := i
			x := xmax
			result = createPoint(x, y)
			break
		}
	}

	fmt.Println(4000000*result.x + result.y)
}

func calculateDistance(a Point, b Point) int {
	return absoluteValue(a.x-b.x) + absoluteValue(a.y-b.y)
}

func absoluteValue(input int) int {
	if input < 0 {
		return -1 * input
	}
	return input
}

func calculateBlockedTilesAtHeightForSensor(sb SBPair, height int) (Interval, error) {
	yDif := absoluteValue(sb.sensor.y - height)
	if yDif > sb.distance {
		return Interval{}, errors.New("no exclusion from sensor at height")
	}
	leftBoundary := sb.sensor.x - sb.distance + yDif
	rightBoundary := sb.sensor.x + sb.distance - yDif
	return createInterval(leftBoundary, rightBoundary), nil
}

func calculateBlockedTilesAtHeightForSensor2(sb SBPair, height int, xmin int, xmax int) (Interval, error) {
	yDif := absoluteValue(sb.sensor.y - height)
	if yDif > sb.distance {
		return Interval{}, errors.New("no exclusion from sensor at height")
	}
	leftBoundary := sb.sensor.x - sb.distance + yDif
	rightBoundary := sb.sensor.x + sb.distance - yDif
	if leftBoundary < xmin {
		leftBoundary = xmin
	}
	if rightBoundary > xmax {
		rightBoundary = xmax
	}
	return createInterval(leftBoundary, rightBoundary), nil
}

func mergeIntervals(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})

	index := 0

	for i := 1; i < len(intervals); i++ {
		if intervals[index].end >= intervals[i].start {
			intervals[index].end = max(intervals[index].end, intervals[i].end)
		} else {
			index++
			intervals[index] = intervals[i]
		}
	}

	return intervals[:index+1]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
