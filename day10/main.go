package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
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
}

func part1() {
	data := ReadInput("./day10/input")
	sort.Ints(data)
	//fmt.Println(data)
	product := CalculateAdapterChainJolts(data)

	fmt.Println(product)
}

// CalculateAdapterChainJolts - Returns the product of 1 jolt differences and 3 jolt differences
// when constructing an adapter chain
func CalculateAdapterChainJolts(sortedCollection []int) int {
	oneJoltDifference := 0
	threeJoltDifference := 0

	if sortedCollection[0]-0 == 1 {
		fmt.Println("one diff initial")
		oneJoltDifference++
	} else if sortedCollection[0]-0 == 3 {
		fmt.Println("three diff initial")
		threeJoltDifference++
	} else {
		fmt.Println("diff is not one or three")
	}

	for i, val := range sortedCollection {
		var next, prev int
		if i == len(sortedCollection)-1 {
			continue
		}

		next = sortedCollection[i+1]
		prev = val

		if next-prev == 1 {
			fmt.Printf("one diff found between %d and %d", next, prev)
			fmt.Println()
			oneJoltDifference++
		} else if next-prev == 3 {
			fmt.Printf("three diff found between %d and %d", next, prev)
			fmt.Println()
			threeJoltDifference++
		} else {
			fmt.Println("diff is not one or three")
		}
	}

	threeJoltDifference++ //the final three difference from your built-in adapter

	return oneJoltDifference * threeJoltDifference
}

func ReadInput(path string) []int {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer file.Close()
	var data []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		val, err := strconv.ParseInt(scanner.Text(), 10, 0)

		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		data = append(data, int(val))
	}

	return data
}
