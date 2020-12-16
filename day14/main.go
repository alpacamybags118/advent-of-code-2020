package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
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
	data := ReadInput("day14/input")
	i := 0
	var sum int64 = 0
	memory := make(map[string]int64)

	for i < len(data) {
		mask := strings.Replace(data[i], "mask = ", "", -1)
		i++

		for i < len(data) && !strings.Contains(data[i], "mask") {
			mem := strings.Split(data[i], " = ")
			reg := regexp.MustCompile("\\d+")
			memAddress, _ := strconv.ParseInt(reg.FindStringSubmatch(mem[0])[0], 10, 64)
			memVal, _ := strconv.ParseInt(mem[1], 10, 64)
			newAddress := ApplyMask(fmt.Sprintf("%036b", memAddress), mask, true)
			newCombinations := FindAllFloatingAddresses(newAddress)
			for _, add := range newCombinations {
				memory[add] = memVal
			}
			i++
		}
	}
	for _, val := range memory {
		sum += val
	}

	fmt.Println(sum)
}

func part1() {
	i := 0
	memory := make(map[int]string)
	data := ReadInput("day14/input")
	for i < len(data) {
		mask := strings.Replace(data[i], "mask = ", "", -1)
		i++

		for i < len(data) && !strings.Contains(data[i], "mask") {
			mem := strings.Split(data[i], " = ")
			reg := regexp.MustCompile("\\d+")
			memAddress, _ := strconv.ParseInt(reg.FindStringSubmatch(mem[0])[0], 10, 0)
			memVal, _ := strconv.ParseInt(mem[1], 10, 0)

			memory[int(memAddress)] = ApplyMask(fmt.Sprintf("%036b", memVal), mask, false)
			i++
		}

	}

	fmt.Println(SumBinaries(memory))
}

// ApplyMask runs the mask against the given binary number (in string format) and returns
// a string representation of the binary result
// if floating is true, X will be applied to the result, represeting a floating bit
func ApplyMask(binary string, mask string, floating bool) string {
	binaryRune := []rune(binary)
	maskRune := []rune(mask)

	for i := 0; i < len(mask); i++ {
		if floating {
			if mask[i] == '1' || mask[i] == 'X' {
				binaryRune[i] = maskRune[i]
			}
		} else if mask[i] != 'X' && mask[i] != binary[i] {
			binaryRune[i] = maskRune[i]
		}
	}

	return string(binaryRune)
}

// FindAllFloatingAddresses will find all combinations of addresses to be writing to
// based on floating bits
func FindAllFloatingAddresses(binary string) []string {
	combinations := []string{}
	xCombinations := GetBinaryUpToNumber(strings.Count(binary, "X"))
	for _, xVal := range xCombinations {
		xValRune := []rune(xVal)
		combination := []rune(binary)
		strLen := 0
		for i := 0; i < len(binary); i++ {
			if binary[i] == 'X' {
				combination[i] = xValRune[strLen]
				strLen++
			}
		}

		combinations = append(combinations, string(combination))
	}

	return combinations
}

// GetBinaryUpToNumber returns all binary representations for the number of bits given
func GetBinaryUpToNumber(bits int) []string {
	number := int64(math.Pow(2, float64(bits)))
	numbers := []string{}
	var i int64 = 0

	for i < number {
		bin := fmt.Sprintf("%036b", i)
		index := 36 - bits
		numbers = append(numbers, bin[index:])
		i++
	}

	return numbers
}

// SumBinaries will sum each binary number in a map and return an integer representation
func SumBinaries(binaryList map[int]string) int {
	sum := int64(0)

	for _, binary := range binaryList {
		val, _ := strconv.ParseInt(binary, 2, 0)
		sum += val
	}

	return int(sum)
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
