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
	totalAnsweredUnaimous := 0
	data := ReadInput("./day6/input")

	for i := 0; i < len(data); i++ {
		answers := []string{}

		for i < len(data) && data[i] != "" {
			answers = append(answers, data[i])
			i++
		}

		groupAnswers := GetGroupAnswers(answers)

		for _, v := range groupAnswers {
			if v == len(answers) {
				totalAnsweredUnaimous++
			}
		}
	}

	fmt.Println(totalAnsweredUnaimous)
}

func part1() {
	totalAnswered := 0
	data := ReadInput("./day6/input")

	for i := 0; i < len(data); i++ {
		answers := []string{}

		for i < len(data) && data[i] != "" {
			answers = append(answers, data[i])
			i++
		}

		groupAnswers := GetGroupAnswers(answers)

		for range groupAnswers {
			totalAnswered++
		}
	}

	fmt.Println(totalAnswered)
}

// GetGroupAnswers - gets a list of answers in which anyone in the group answered yes, along with the count of how many answered yes
func GetGroupAnswers(answers []string) map[rune]int {
	answerMap := make(map[rune]int)
	for _, answer := range answers {
		for _, char := range answer {
			_, exists := answerMap[char]

			if !exists {
				answerMap[char] = 1
			} else {
				answerMap[char]++
			}
		}
	}

	return answerMap
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
