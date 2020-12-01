package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var i, j int
	sum := int64(2020)
	numArr := ReadInput()
	cache := make(map[int64]int64)

	for i = 0; i < len(numArr); i++ {
		for j = i + 1; j < len(numArr); j++ {
			val, exists := cache[sum-numArr[i]-numArr[j]]
			if exists {
				fmt.Println(numArr[i] * numArr[j] * val)
			} else {
				cache[numArr[j]] = numArr[j]
				cache[numArr[i]] = numArr[i]
			}
		}

	}
}

// Returns input from file
func ReadInput() []int64 {
	file, err := os.Open("./day1/input")

	if err != nil {
		fmt.Println(err.Error())
	}

	defer file.Close()
	var numbers []int64
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		num, err := strconv.ParseInt(scanner.Text(), 10, 64)

		if err != nil {
			fmt.Println(err.Error())
		}

		numbers = append(numbers, num)
	}

	return numbers
}
