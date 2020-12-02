package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	var validPasswordCount int
	passwordData := ReadInput()

	for _, data := range passwordData {
		input := strings.Split(data, " ")
		first, _ := strconv.ParseInt(strings.Split(input[0], "-")[0], 10, 64)
		last, _ := strconv.ParseInt(strings.Split(input[0], "-")[1], 10, 64)
		char := strings.Split(input[1], ":")[0]
		password := input[2]

		var firstContained, lastContained bool

		firstContained = string(password[first-1]) == char
		lastContained = string(password[last-1]) == char

		if (firstContained || lastContained) && !(firstContained && lastContained) {
			validPasswordCount++
		}
	}

	fmt.Print(validPasswordCount)
}

func part1() {
	var validPasswordCount int
	passwordData := ReadInput()

	for _, data := range passwordData {
		input := strings.Split(data, " ")
		min, _ := strconv.ParseInt(strings.Split(input[0], "-")[0], 10, 64)
		max, _ := strconv.ParseInt(strings.Split(input[0], "-")[1], 10, 64)
		char := strings.Split(input[1], ":")[0]
		password := input[2]

		charCount := int64(strings.Count(password, char))
		if charCount >= min && charCount <= max {
			validPasswordCount++
		}
	}

	fmt.Print(validPasswordCount)
}

// Returns input from file
func ReadInput() []string {
	file, err := os.Open("./day2/input")

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
