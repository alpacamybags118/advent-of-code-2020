package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	name              string
	rules             [][]int
	valuesNotIncludes []int
}

type TicketField struct {
	name  string
	index int
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
	data := ReadInput("day16/input")

	ruleList := GetRules(data)
	myTicket := GetMyTicket(data)
	tickets := GetNearbyTickets(data)

	tickets = GenerateValidTicketList(tickets, GetInvalidValues(tickets, ruleList, true))

	fields := FindTicketFields(tickets, ruleList)
	product := 1
	for _, field := range fields {
		if strings.Contains(field.name, "departure") {
			product = product * myTicket[field.index]
		}
	}

	fmt.Println(product)

}

func part1() {
	data := ReadInput("day16/input")

	// parse out the rules
	ruleList := GetRules(data)
	tickets := GetNearbyTickets(data)

	invalidVals := GetInvalidValues(tickets, ruleList, false)

	fmt.Println(invalidVals)
	sum := 0

	for _, val := range invalidVals {
		sum += val
	}

	fmt.Println(sum)
}

// GetRules extracts the rules from the given input
func GetRules(data []string) []Rule {
	ruleList := []Rule{}
	i := 0

	for data[i] != "" {
		numRanges := [][]int{}
		rule := strings.Split(data[i], ":")
		ranges := strings.Split(rule[1], "or")

		for _, r := range ranges {
			newRange := []int{}
			nums := strings.Split(r, "-")
			for _, num := range nums {
				numInt, err := strconv.ParseInt(strings.TrimSpace(num), 10, 0)

				if err != nil {
					fmt.Println(err.Error())
				}

				newRange = append(newRange, int(numInt))
			}

			numRanges = append(numRanges, newRange)
		}

		missingNumbers := FindMissingNumbers(numRanges)

		ruleList = append(ruleList, Rule{
			name:              rule[0],
			rules:             numRanges,
			valuesNotIncludes: missingNumbers,
		})
		i++
	}

	return ruleList
}

//GetMyTicket gets my ticket
func GetMyTicket(data []string) []int {
	i := 0
	ticket := []int{}
	for !strings.Contains(data[i], "your") {
		i++
	}
	i++

	values := strings.Split(data[i], ",")
	for _, value := range values {
		val, _ := strconv.ParseInt(value, 10, 0)
		ticket = append(ticket, int(val))
	}

	return ticket
}

// GetNearbyTickets gets a list of all the tickets, with their data
func GetNearbyTickets(data []string) [][]int {
	i := 0
	tickets := [][]int{}
	for !strings.Contains(data[i], "nearby") {
		i++
	}
	i++
	for i < len(data) {
		values := strings.Split(data[i], ",")
		ints := []int{}
		for _, value := range values {
			val, _ := strconv.ParseInt(value, 10, 0)
			ints = append(ints, int(val))
		}

		tickets = append(tickets, ints)
		i++
	}

	return tickets
}

//GetInvalidValues evaluates the rulesets against each ticket and returns all values in the tickets that are invald
// set returnTicketIndexes to true to return an array of the invalid ticket indexes instead of values in each ticket
func GetInvalidValues(tickets [][]int, rules []Rule, returnTicketIndexes bool) []int {
	invalidVals := []int{}

	for index, ticket := range tickets {
		for _, value := range ticket {
			if !DoesValueFollowRules(value, rules) {
				if returnTicketIndexes {
					invalidVals = append(invalidVals, index)
					break
				}
				invalidVals = append(invalidVals, value)
			}
		}
	}

	return invalidVals
}

// FindTicketFields Finds where each field is in a ticket
func FindTicketFields(tickets [][]int, rules []Rule) []TicketField {
	ticketFields := []TicketField{}
	foundRuleList := [][]int{}
	i := 0

	for i < len(tickets[0]) {
		foundRule := GenerateNumberArray(0, len(rules)-1)
		for _, ticket := range tickets {
			for index, rule := range rules {
				found := InIntArray(rule.valuesNotIncludes, ticket[i])

				if found {
					foundRule = RemoveItemFromArray(foundRule, index)
				}
			}

			if len(foundRule) == 1 {
				break
			}
		}

		foundRuleList = append(foundRuleList, foundRule)
		i++
	}

	foundRuleList = FindUniques(foundRuleList)

	for index, val := range foundRuleList {
		ticketFields = append(ticketFields, TicketField{
			name:  rules[val[0]].name,
			index: index,
		})
	}

	return ticketFields
}

// DoesValueFollowRules Evaluates a val against all rulesets and returns if it follows all rules
func DoesValueFollowRules(value int, rules []Rule) bool {
	for _, rule := range rules {
		for _, condition := range rule.rules {
			if value >= condition[0] && value <= condition[1] {
				return true
			}
		}
	}

	return false
}

// GenerateValidTicketList filters out invalid tickets from ticket list and returns valid tickets
func GenerateValidTicketList(tickets [][]int, invalidTickets []int) [][]int {
	validTickets := [][]int{}

	for index, val := range tickets {
		containsInvald := false
		for _, invalid := range invalidTickets {
			if index == invalid {
				containsInvald = true
				break
			}
		}

		if !containsInvald {
			validTickets = append(validTickets, val)
		}
	}

	return validTickets
}

// FindMissingNumbers returns an array of numbers not included in the range of the rulesets
func FindMissingNumbers(rules [][]int) []int {
	min := rules[0][0]
	max := rules[1][1]
	fullNumberList := GenerateNumberArray(min, max)
	containedNumberList := []int{}
	containedNumberList = append(containedNumberList, GenerateNumberArray(rules[0][0], rules[0][1])...)
	containedNumberList = append(containedNumberList, GenerateNumberArray(rules[1][0], rules[1][1])...)
	missingNumbers := []int{}

	for _, num := range fullNumberList {
		found := false

		for _, val := range containedNumberList {
			if num == val {
				found = true
				break
			}
		}

		if !found {
			missingNumbers = append(missingNumbers, num)
		}
	}

	return missingNumbers
}

// GenerateNumberArray generate an array of numbers from the start to end, inclusive
func GenerateNumberArray(start int, end int) []int {
	intArr := []int{}

	for i := start; i <= end; i++ {
		intArr = append(intArr, i)
	}

	return intArr
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

// InIntArray searches for a value in an int array
func InIntArray(array []int, number int) bool {
	found := false

	for _, val := range array {
		if val == number {
			found = true
			break
		}
	}

	return found
}

// RemoveItemFromArray returns an array with the number provided removed
func RemoveItemFromArray(array []int, number int) []int {
	newArr := []int{}
	for _, val := range array {
		if val != number {
			newArr = append(newArr, val)
		}
	}

	return newArr
}

// RemoveRuleFromArray returns the rule array with the removed value
func RemoveRuleFromArray(array []Rule, index int) []Rule {
	newArr := []Rule{}
	for i, val := range array {
		if i != index {
			newArr = append(newArr, val)
		}
	}

	return newArr
}

// FindUniques breaks down the array of array ints into an array of array ints w/ individual values by removeing dupes
// from each
func FindUniques(list [][]int) [][]int {
	analyzedRules := 0
	usedRules := []int{}

	for analyzedRules < len(list) {
		for i, rule := range list {
			if len(rule) == 1 && !InIntArray(usedRules, i) {
				usedRules = append(usedRules, i)
				for index, val := range list {
					if InIntArray(val, rule[0]) && !InIntArray(usedRules, index) {
						list[index] = RemoveItemFromArray(list[index], rule[0])
					}
				}
			}
		}

		analyzedRules++
	}

	return list
}
