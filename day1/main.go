package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var x, i, z int
	sum := int64(2020)
	numArr := ReadInput()

	for i = 0; i < len(numArr); i++ {
		for x = i + 1; x < len(numArr); x++ {
			for z = i + 2; z < len(numArr); z++ {
				if numArr[i]+numArr[x]+numArr[z] == sum {
					fmt.Println(numArr[i] * numArr[x] * numArr[z])
				}
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
