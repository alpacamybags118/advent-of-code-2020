package main

import (
	"bufio"
	"fmt"
	"os"
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
	data := ReadInput("day15/input")
	result := PlaySpeakingGame(data, 30000000)

	fmt.Println(result)
}

func part1() {
	data := ReadInput("day15/input")
	result := PlaySpeakingGame(data, 2020)

	fmt.Println(result)
}

// PlaySpeakingGame runs a game of the speaking game, up to the number of turns given
func PlaySpeakingGame(input []int, turns int) int {
	memory := make(map[int][]int)
	currentTurn := 0
	lastNumberSpoken := input[len(input)-1]
	for _, val := range input {
		memory[val] = []int{currentTurn}
		currentTurn++
	}

	for currentTurn < turns {
		val, _ := memory[lastNumberSpoken]

		if len(val) == 1 {
			CommitToMemory(memory, 0, currentTurn)
			lastNumberSpoken = 0
		} else {
			lastNumberSpoken = val[1] - val[0]
			CommitToMemory(memory, lastNumberSpoken, currentTurn)
		}
		currentTurn++
	}

	return lastNumberSpoken
}

// CommitToMemory adds a number to our "memory", keeping track of the last two times we saw it
func CommitToMemory(memory map[int][]int, number int, turn int) map[int][]int {
	val, exists := memory[number]

	if !exists {
		memory[number] = []int{turn}
	} else if len(val) == 1 {
		memory[number] = append(memory[number], turn)
	} else {
		memory[number][0] = val[1]
		memory[number][1] = turn
	}

	return memory
}

// ReadInput - Reads input file from a path, line by line, and returns it as a string array
func ReadInput(path string) []int {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer file.Close()
	var data []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		val, _ := strconv.ParseInt(scanner.Text(), 10, 0)
		data = append(data, int(val))
	}

	return data
}
