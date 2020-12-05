package main

import (
	"bufio"
	"fmt"
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
	data := ReadInput("./day4/input")
	validPassportCount := 0

	for i := 0; i < len(data); i++ {
		var passportData []string
		for i < len(data) && data[i] != "" {
			passportData = append(passportData, data[i])
			i++
		}

		if len(passportData) > 1 {
			passportData = []string{strings.Join(passportData, " ")}
		}

		passport := ProcessPassport(strings.Split(passportData[0], " "))

		if ValidatePassport(passport, true) {
			validPassportCount++
		}
	}

	fmt.Println(validPassportCount)
}

func part1() {
	data := ReadInput("./day4/input")
	validPassportCount := 0

	for i := 0; i < len(data); i++ {
		var passportData []string
		for i < len(data) && data[i] != "" {
			passportData = append(passportData, data[i])
			i++
		}

		if len(passportData) > 1 {
			passportData = []string{strings.Join(passportData, " ")}
		}

		passport := ProcessPassport(strings.Split(passportData[0], " "))

		if ValidatePassport(passport, false) {
			validPassportCount++
		}
	}

	fmt.Println(validPassportCount)
}

// ProcessPassport - converts string data into a map of passport data
func ProcessPassport(input []string) map[string]string {
	passport := make(map[string]string)
	for _, field := range input {
		data := strings.Split(field, ":")
		passport[data[0]] = data[1]
	}

	return passport
}

// ValidatePassport - valids the passport is valid by checking if required fields exist
// optionally can check if each field is valid
func ValidatePassport(passport map[string]string, validateData bool) bool {
	requiredFields := [...]string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, field := range requiredFields {
		val, ok := passport[field]

		if !ok {
			return false
		}

		if validateData {
			if !ValidatePassportData(field, val) {
				return false
			}
		}
	}

	return true
}

// ValidatePassportData - validates each field in the passport
func ValidatePassportData(field string, val string) bool {
	switch field {
	case "byr":
		byr, err := strconv.ParseInt(val, 10, 0)

		if err != nil {
			return false
		}

		if byr < 1920 || byr > 2002 {
			return false
		}

	case "iyr":
		iyr, err := strconv.ParseInt(val, 10, 0)
		if err != nil {
			return false
		}

		if iyr < 2010 || iyr > 2020 {
			return false
		}

	case "eyr":
		eyr, err := strconv.ParseInt(val, 10, 0)

		if err != nil {
			return false
		}

		if eyr < 2020 || eyr > 2030 {
			return false
		}

	case "hgt":
		system := string([]byte{val[len(val)-2], val[len(val)-1]})
		height, err := strconv.ParseInt(strings.TrimSuffix(val, system), 10, 0)

		if err != nil {
			return false
		}

		if system == "cm" && (height < 150 || height > 193) {
			return false
		} else if system == "in" && (height < 59 || height > 76) {
			return false
		}

	case "hcl":
		if match, err := regexp.MatchString("#[0-9a-f]{6}", val); !match || err != nil {
			return false
		}

	case "ecl":
		found := false
		for _, color := range []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"} {
			if val == color {
				found = true
			}
		}

		return found

	case "pid":
		if len(val) != 9 {
			return false
		}
		if _, err := strconv.ParseInt(val, 10, 0); err != nil {
			return false
		}
	}

	return true
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
