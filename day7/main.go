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
	totalBags := 0
	data := ReadInput("./day7/input")
	bags := MakeBagMap(data)
	fmt.Println(bags["shiny gold"])
	totalBags = TotalBags("shiny gold", bags)

	fmt.Println(totalBags)
}

func part1() {
	totalBags := 0
	data := ReadInput("./day7/input")
	bags := MakeBagMap(data)

	for bag, _ := range bags {
		if(bag == "shiny gold") {
			continue
		}
		if FindGoldenBag(bag, bags) {
			totalBags++
		}
	}

	fmt.Println(totalBags)

}

// MakeBagMap - Creates a key-value map of bags, which describe the color and what bags each can hold
func MakeBagMap(data []string) map[string][]Bag {
	bags := make(map[string][]Bag)
	for _, in := range data {
		in = strings.Replace(in, "bags", "", -1)
		in = strings.Replace(in, "bag", "", -1)
		in = strings.Replace(in, ".", "", -1)

		rule := strings.Split(in, "contain")
		sourceBagName := strings.TrimSpace(rule[0])
		destinationBags := []Bag{}

		for _, val := range strings.Split(rule[1], ",") {
			bagDetails := strings.Split(val, " ")
			bagName := strings.TrimSpace(fmt.Sprintf("%s %s", bagDetails[2], bagDetails[3]))
			bagCount, _ := strconv.ParseInt(bagDetails[1], 10 ,0)
			bag := Bag{
				color: bagName,
				count: bagCount,
			}
			destinationBags = append(destinationBags, bag)
		}

		bags[sourceBagName] = destinationBags
	}

	return bags
}

func TotalBags(bag string, bagMap map[string][]Bag) int {
	total := 0
	val, _ := bagMap[bag]
	for _, destBag := range(val) {
		total += int(destBag.count) + int(destBag.count) * TotalBags(destBag.color, bagMap)
	}
	return total
}

func FindGoldenBag(bag string, bagMap map[string][]Bag) bool {
	if bag == "shiny gold" {
		return true
	} else if val, _ := bagMap[bag]; len(val) == 0 {
		return false
	} else {
		return Any(bagMap[bag], bagMap, FindGoldenBag)
	}
}

func Any(vs []Bag, bagMap map[string][]Bag, f func(string, map[string][]Bag) bool) bool {
	for _, v := range vs {
		if f(v.color, bagMap) {
			return true
		}
	}
	return false
}

func All(vs []Bag, bagMap map[string][]Bag, f func(string, map[string][]Bag) int) int {
	total := 1
	for _, v := range vs {
		total += int(v.count) * f(v.color, bagMap) 
	}
	return total
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
