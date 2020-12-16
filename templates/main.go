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
	data := ReadInput("#DAY#/input")
	//do more stuff
}

func part1() {
	data := ReadInput("#DAY#/input")
	//do stuff
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
