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
	data := ReadInput("./day9/input")

	sums := FindAllSumsOfInvalid(data, 530627549)

	sort.Ints(sums)

	fmt.Println(sums[0] + sums[len(sums)-1])
}

func part1() {
	data := ReadInput("./day9/input")

	for i := 25; i < len(data); i++ {
		if !IsSumOfPrevious(data, i, i-25) {
			fmt.Println(data[i])
		}
	}
}

//IsSumOfPrevious determines whether the item at the current index of a collection can be generated
// as a sum of two previous values in the array, starting at starting index
func IsSumOfPrevious(collection []int, currentIndex int, startIndex int) bool {
	for i := startIndex; i < currentIndex; i++ {
		for x := startIndex + 1; x < currentIndex; x++ {
			if collection[x]+collection[i] == collection[currentIndex] {
				return true
			}
		}
	}

	return false
}

// FindAllSumsOfInvalid - Returns a list of all the integers that, summed in some order,
// return the value of the collection at the current index
func FindAllSumsOfInvalid(sortedCollection []int, invalid int) []int {
	numberList := []int{}

	for i := 0; i < len(sortedCollection); i++ {
		index := i
		currentSum := 0
		for currentSum <= invalid {
			currentSum += sortedCollection[index]

			if currentSum == invalid {
				for x := i; x < index; x++ {
					numberList = append(numberList, sortedCollection[x])
				}
			}

			index++
		}
	}

	return numberList
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
