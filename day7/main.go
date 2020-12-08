package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Bag struct {
	color string
	holds []Bag
}

var total int

func main() {
	arg := os.Args[1]

	if arg == "part1" {
		part1()
	} else if arg == "part2" {
		part2()
	}
}

func part2() {

}

func part1() {
	totalBags := 0
	data := ReadInput("./day7/input")
	bags := MakeBagMap(data)

	for bag, _ := range bags {
		if FindGoldenBag(bag, bags) {
			totalBags++
		}
	}

	fmt.Println(totalBags)

}

// MakeBagMap - Creates a key-value map of bags, which describe the color and what bags each can hold
func MakeBagMap(data []string) map[string][]string {
	bags := make(map[string][]string)
	for _, in := range data {
		in = strings.Replace(in, "bags", "", -1)
		in = strings.Replace(in, "bag", "", -1)
		in = strings.Replace(in, ".", "", -1)

		rule := strings.Split(in, "contain")
		sourceBagName := strings.TrimSpace(rule[0])
		destinationBags := []string{}

		for _, val := range strings.Split(rule[1], ",") {
			bagDetails := strings.Split(val, " ")
			bagName := strings.TrimSpace(fmt.Sprintf("%s %s", bagDetails[2], bagDetails[3]))

			destinationBags = append(destinationBags, bagName)
		}

		bags[sourceBagName] = destinationBags
	}

	return bags
}

func FindGoldenBag(bag string, bagMap map[string][]string) bool {
	if bag == "shiny gold" {
		return true
	} else if val, _ := bagMap[bag]; len(val) == 0 {
		return false
	} else {
		return Any(bagMap[bag], bagMap, FindGoldenBag)
	}
}

func Any(vs []string, bagMap map[string][]string, f func(string, map[string][]string) bool) bool {
	for _, v := range vs {
		if f(v, bagMap) {
			return true
		}
	}
	return false
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
