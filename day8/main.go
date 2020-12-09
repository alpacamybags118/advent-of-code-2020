package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Bag struct {
	color string
	count int64
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
	instructions := ReadInput("./day8/input")
	for i := 0; i < len(instructions); i++ {
		instructon := strings.Split(instructions[i], " ")
		if instructon[0] == "jmp" {
			tmpInstruction := instructions[i]
			instructions[i] = strings.Replace(instructions[i], "jmp", "nop", -1)
			res, loop := RunComputer(instructions)

			if !loop {
				fmt.Println(res)
			}

			instructions[i] = tmpInstruction
		} else if instructon[0] == "nop" {
			tmpInstruction := instructions[i]
			instructions[i] = strings.Replace(instructions[i], "nop", "jmp", -1)

			res, loop := RunComputer(instructions)

			if !loop {
				fmt.Println(res)
			}

			instructions[i] = tmpInstruction
		}
	}
}

func part1() {
	instructions := ReadInput("./day8/input")

	acc, _ := RunComputer(instructions)

	fmt.Println(acc)
}

// RunComputer - Runs the instructions one by one until
// it either hits a previously ran insruction or ends
// returns the value of acc at stop and a bool indicating if it hit an infinite loop
func RunComputer(instructons []string) (int, bool) {
	previouslyRanInstructions := make(map[int]string)
	acc := 0
	i := 0
	infLoop := false

	for i < len(instructons) {
		_, exists := previouslyRanInstructions[i]

		if exists {
			infLoop = true
			break
		}

		previouslyRanInstructions[i] = instructons[i]

		accChange, indexChange := RunInstruction(instructons[i])

		acc += accChange
		i += indexChange
	}

	return acc, infLoop
}

// RunInstruction - Runs the instruction and returns two values, first being how much to increment the acc
// and the second how much to inc the index for the next instruction
func RunInstruction(instruction string) (int, int) {
	index := 0
	acc := 0
	instructonParts := strings.Split(instruction, " ")
	operator := string(instructonParts[1][0])
	value, _ := strconv.ParseInt(strings.Replace(instructonParts[1], operator, "", -1), 10, 0)

	switch instructonParts[0] {
	case "acc":
		switch operator {
		case "+":
			acc += int(value)

		case "-":
			acc -= int(value)
		}
		index++
	case "jmp":
		switch operator {
		case "+":
			index += int(value)

		case "-":
			index -= int(value)
		}

	case "nop":
		index++
	}
	return acc, index
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
