package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1]

	if arg == "part1" {
		part1()
	} else if arg == "part2" {
		part2()
	}
}

func part2() {
	data := ReadInput("./day5/input")
	highestSeatId := 0
	lowestSeatId := 0
	seats := make(map[int]int)
	for _, seat := range data {
		var row, column string

		for i, c := range seat {
			if i < 7 {
				row = fmt.Sprintf("%s%c", row, c)
			} else {
				column = fmt.Sprintf("%s%c", column, c)
			}
		}

		rowVal := FindRowValue(row)
		colVal := FindColumnValue(column)
		seatID := (rowVal * 8) + colVal
		seats[seatID] = seatID

		if seatID > highestSeatId {
			highestSeatId = seatID
		}

		if lowestSeatId == 0 {
			lowestSeatId = seatID
		} else if seatID < lowestSeatId {
			lowestSeatId = seatID
		}
	}

	var x int = lowestSeatId
	for x <= highestSeatId {
		_, exists := seats[x]

		if !exists {
			fmt.Println(x)
		}

		x++
	}
}

func part1() {
	highestSeatId := 0
	data := ReadInput("./day5/input")

	for _, seat := range data {
		var row, column string

		for i, c := range seat {
			if i < 7 {
				row = fmt.Sprintf("%s%c", row, c)
			} else {
				column = fmt.Sprintf("%s%c", column, c)
			}
		}

		rowVal := FindRowValue(row)
		colVal := FindColumnValue(column)
		seatID := (rowVal * 8) + colVal
		if seatID > highestSeatId {
			highestSeatId = seatID
		}
	}

	fmt.Println(highestSeatId)
}

// FindRowValue - Finds the row the passenger is in
func FindRowValue(row string) int {
	high := 127
	low := 0
	var rowVal int

	for i, char := range row {
		rowVal = (high + low) / 2
		if char == 'F' {
			high = rowVal - 1
		} else {
			low = rowVal + 1

			if i == len(row)-1 {
				rowVal++
			}
		}
	}

	return rowVal
}

// FindColumnValue - Finds the column the passenger is in
func FindColumnValue(col string) int {
	high := 7
	low := 0
	var colVal int

	for i, char := range col {
		colVal = (high + low) / 2
		if char == 'L' {
			high = colVal - 1
		} else {
			low = colVal + 1

			if i == len(col)-1 {
				colVal++
			}
		}
	}

	return colVal
}

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
