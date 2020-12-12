package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

type Path struct {
	direction   string
	startingPos []int
}

func main() {
	arg := os.Args[1]

	if arg == "part1" {
		part1()
	} else if arg == "part2" {
		part2()
	}
}

func part2() {
	done := false
	data := ReadInput("day11/input")
	newArrange := []string{}
	prevArrange := data

	for !done {
		newArrange = RunSeatingModel(prevArrange, true)
		if reflect.DeepEqual(prevArrange, newArrange) {
			done = true
		}

		prevArrange = newArrange
	}

	fmt.Println(GetOccupiedSeatCount(newArrange))
}

func part1() {
	done := false
	data := ReadInput("day11/input")
	newArrange := []string{}
	prevArrange := data

	for !done {
		newArrange = RunSeatingModel(prevArrange, false)
		if reflect.DeepEqual(prevArrange, newArrange) {
			done = true
		}

		prevArrange = newArrange
	}

	fmt.Println(GetOccupiedSeatCount(newArrange))

}

// RunSeatingModel - Runs seating ruleset against given seating chart and returns the new seatchart
// with applied rules. in addition, a boolean is provided to turn on the additional sight rules
// when running the model.
func RunSeatingModel(seatingChart []string, sightRules bool) []string {
	newArrange := []string{}
	for rowIndex, row := range seatingChart {
		rowString := ""
		for posIndex, pos := range row {
			adjacentSpots := 0
			maxAdjacent := 0
			switch string(pos) {
			case ".":
				rowString = rowString + "."
			case "L":
				if sightRules {
					adjacentSpots = FindFirstSeenAdjacentOccupiedSeatCount([]int{rowIndex, posIndex}, seatingChart)
				} else {
					adjacentSpots = FindAdjacentOccupiedSeatCount([]int{rowIndex, posIndex}, seatingChart)
				}

				if adjacentSpots == 0 {
					rowString = rowString + "#"
				} else {
					rowString = rowString + "L"
				}
			case "#":
				if sightRules {
					adjacentSpots = FindFirstSeenAdjacentOccupiedSeatCount([]int{rowIndex, posIndex}, seatingChart)
					maxAdjacent = 5
				} else {
					adjacentSpots = FindAdjacentOccupiedSeatCount([]int{rowIndex, posIndex}, seatingChart)
					maxAdjacent = 4
				}

				if adjacentSpots >= maxAdjacent {
					rowString = rowString + "L"
				} else {
					rowString = rowString + "#"
				}
			}
		}
		newArrange = append(newArrange, rowString)
	}

	return newArrange
}

// FindAdjacentOccupiedSeatCount finds the number of adjacent seats to the given seat that
// are occupied and returns the count
func FindAdjacentOccupiedSeatCount(currentPos []int, seatingChart []string) int {
	positionsFilled := 0
	positionsToSearch := [][]int{[]int{currentPos[0] - 1, currentPos[1]},
		[]int{currentPos[0] - 1, currentPos[1] + 1},
		[]int{currentPos[0] - 1, currentPos[1] - 1},
		[]int{currentPos[0], currentPos[1] - 1},
		[]int{currentPos[0], currentPos[1] + 1},
		[]int{currentPos[0] + 1, currentPos[1]},
		[]int{currentPos[0] + 1, currentPos[1] + 1},
		[]int{currentPos[0] + 1, currentPos[1] - 1},
	}

	for _, pos := range positionsToSearch {
		if pos[0] < 0 || pos[0] >= len(seatingChart) {
			continue
		}

		if pos[1] < 0 || pos[1] >= len(seatingChart[pos[0]]) {
			continue
		}

		if string(seatingChart[pos[0]][pos[1]]) == "#" {
			positionsFilled++
		}
	}

	return positionsFilled
}

// FindFirstSeenAdjacentOccupiedSeatCount finds the number of adjacent seats to the given seat that
// are occupied based on first seen rules and returns the count
func FindFirstSeenAdjacentOccupiedSeatCount(currentPos []int, seatingChart []string) int {
	positionsFilled := 0

	startingPositions := []Path{
		Path{direction: "N", startingPos: []int{currentPos[0] - 1, currentPos[1]}},
		Path{direction: "NE", startingPos: []int{currentPos[0] - 1, currentPos[1] + 1}},
		Path{direction: "E", startingPos: []int{currentPos[0], currentPos[1] + 1}},
		Path{direction: "SE", startingPos: []int{currentPos[0] + 1, currentPos[1] + 1}},
		Path{direction: "S", startingPos: []int{currentPos[0] + 1, currentPos[1]}},
		Path{direction: "SW", startingPos: []int{currentPos[0] + 1, currentPos[1] - 1}},
		Path{direction: "W", startingPos: []int{currentPos[0], currentPos[1] - 1}},
		Path{direction: "NW", startingPos: []int{currentPos[0] - 1, currentPos[1] - 1}},
	}

	for _, pos := range startingPositions {
		direction := pos.direction
		current := pos.startingPos

		// dont go past any boundries in the seating chart
		for current[0] >= 0 && current[0] < len(seatingChart) && current[1] >= 0 && current[1] < len(seatingChart[0]) {
			if string(seatingChart[current[0]][current[1]]) == "#" {
				positionsFilled++
				break
			} else if string(seatingChart[current[0]][current[1]]) == "L" {
				break
			}

			switch direction {
			case "N":
				current[0]--
			case "NE":
				current[0]--
				current[1]++
			case "E":
				current[1]++
			case "SE":
				current[0]++
				current[1]++
			case "S":
				current[0]++
			case "SW":
				current[0]++
				current[1]--
			case "W":
				current[1]--
			case "NW":
				current[0]--
				current[1]--
			default:
				fmt.Println("direction invalid")
			}
		}

	}
	return positionsFilled
}

// GetOccupiedSeatCount returns the count of occupied seats for a given seat chart
func GetOccupiedSeatCount(seatingChart []string) int {
	count := 0
	for _, row := range seatingChart {
		count += strings.Count(row, "#")
	}

	return count
}

// ReadInput - Reads input file from a path, line by line, and returns it as a string array
func ReadInput(path string) []string {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer file.Close()
	var data []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	return data
}
