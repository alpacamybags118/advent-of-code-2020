package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

type Bus struct {
	ID                  int
	subsequentTimeStamp int
}

func main() {
	arg := os.Args[1]

	if arg == "part1" {
		part1()
	} else if arg == "part2" {
		part2()
	}
}

// notes on first example
// 4199
// 782
// 17 * 13 * 19 = 782
// firstoccurnace = 3417
// firstoccurance + 4199 = nextoccurance , and so on
func part2() {
	data := ReadInput("day13/input")
	busList := GetValidBusIds(strings.Split(data[1], ","))
	n := []*big.Int{}
	a := []*big.Int{}
	N := 1

	for _, bus := range busList {
		N *= bus.ID
		n = append(n, big.NewInt(int64(bus.ID)))
		a = append(a, big.NewInt(int64(bus.subsequentTimeStamp)))
	}
	result, _ := crt(a, n)

	fmt.Println(int64(N) - result.Int64())
}

func part1() {
	data := ReadInput("day13/input")
	arrivalTime, _ := strconv.ParseInt(data[0], 10, 0)
	busList := GetValidBusIds(strings.Split(data[1], ","))
	foundBus := false
	foundBusID := 0
	searchTime := int(arrivalTime)

	fmt.Println(searchTime)
	fmt.Println(busList)
	for !foundBus {
		for _, bus := range busList {
			if IsBusAvailable(bus.ID, searchTime) {
				fmt.Println("found")
				foundBus = true
				foundBusID = bus.ID
				break
			}
		}
		if !foundBus {
			searchTime++
		}
	}

	fmt.Println((searchTime - int(arrivalTime)) * foundBusID)
}

// had to look up how to code this. this is the chinese remainder theorem
// source: https://rosettacode.org/wiki/Chinese_remainder_theorem#Go
func crt(a, n []*big.Int) (*big.Int, error) {
	p := new(big.Int).Set(n[0])
	for _, n1 := range n[1:] {
		p.Mul(p, n1)
	}
	var x, q, s, z big.Int
	for i, n1 := range n {
		q.Div(p, n1)
		z.GCD(nil, &s, n1, &q)
		if z.Cmp(big.NewInt(1)) != 0 {
			return nil, fmt.Errorf("%d not coprime", n1)
		}
		x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
	}
	return x.Mod(&x, p), nil
}

//GetValidBusIds returns a list of in-operation busses based on a list of all busses
func GetValidBusIds(busList []string) []Bus {
	validBusses := []Bus{}

	for index, val := range busList {
		busID, err := strconv.ParseInt(val, 10, 0)

		if err != nil {
			continue
		}
		bus := Bus{
			ID:                  int(busID),
			subsequentTimeStamp: index,
		}

		validBusses = append(validBusses, bus)
	}

	return validBusses
}

// IsBusAvailable will return if a bus is available at the given time
func IsBusAvailable(busID int, time int) bool {
	if time%busID == 0 {
		return true
	}

	return false
}

// HasSubsequentBusArrivals determines if each bus after the staringbus follows the
// subsequent time arrival rules
func HasSubsequentBusArrivals(startTime int, busList []Bus) bool {
	for index, bus := range busList {
		if index == 0 {
			continue
		}

		if !IsBusAvailable(bus.ID, startTime+bus.subsequentTimeStamp) {
			return false
		}
	}
	return true
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
