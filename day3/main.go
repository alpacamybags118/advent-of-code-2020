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
	treeProduct := 1
	data := ReadInput()
	slopes := [][]int{[]int{1, 1}, []int{3, 1}, []int{5, 1}, []int{7, 1}, []int{1, 2}}

	for _, slope := range slopes {
		treeProduct *= findTreesEncountered(data, slope[0], slope[1])
	}

	fmt.Println(treeProduct)
}

func part1() {
	data := ReadInput()
	totalTrees := findTreesEncountered(data, 3, 1)

	fmt.Println(totalTrees)
}

func findTreesEncountered(dataMap []string, xSlope int, ySlope int) int {
	var x, y int = 0, 0
	var totalTrees int = 0

	for y < len(dataMap)-1 {
		if (x + xSlope) >= len(dataMap[y]) {
			x = (x + xSlope) - len(dataMap[y])
		} else {
			x = x + xSlope
		}
		y += ySlope

		if dataMap[y][x] == '#' {
			totalTrees++
		}
	}

	return totalTrees
}

// Returns input from file
func ReadInput() []string {
	file, err := os.Open("./day3/input")

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
