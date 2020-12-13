package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Waypoint struct {
	exists      bool
	east_west   int
	north_south int
}

type Ship struct {
	currentlyFacing int
	east_west       int
	north_south     int
	waypoint        Waypoint
}

func main() {
	arg := os.Args[1]

	if arg == "part1" {
		part1()
	} else if arg == "part2" {
		part2()
	}
}

// 51124 is too low
func part2() {
	data := ReadInput("day12/input")
	ship := Ship{
		currentlyFacing: 90,
		east_west:       0,
		north_south:     0,
		waypoint: Waypoint{
			exists:      true,
			east_west:   10,
			north_south: 1,
		},
	}

	ship = CalculatePosition(ship, data)
	manhattanX := int(math.Abs(float64(ship.east_west)))
	manhattanY := int(math.Abs(float64(ship.north_south)))

	fmt.Println(manhattanX + manhattanY)
}

func part1() {
	data := ReadInput("day12/input")
	ship := Ship{
		currentlyFacing: 90,
		east_west:       0,
		north_south:     0,
		waypoint: Waypoint{
			exists: false,
		},
	}

	ship = CalculatePosition(ship, data)
	manhattanX := int(math.Abs(float64(ship.east_west)))
	manhattanY := int(math.Abs(float64(ship.north_south)))

	fmt.Println(manhattanX + manhattanY)
}

// CalculatePosition returns a set of coordinations that are the position after
// running the ruleset on the initially provided position
func CalculatePosition(ship Ship, instructions []string) Ship {
	for _, instruction := range instructions {
		direction := instruction[0]
		distance, err := strconv.ParseInt(strings.Replace(instruction, string(direction), "", -1), 10, 0)

		if err != nil {
			fmt.Println(instruction)
			fmt.Println(err.Error())
			return ship
		}

		ship = MoveShip(ship, string(direction), int(distance))
		//fmt.Println(ship.position)
	}

	return ship
}

// MoveShip returns the position of the ship after running the direction and distance rules
func MoveShip(ship Ship, direction string, distance int) Ship {
	switch direction {
	case "N":
		if ship.waypoint.exists {
			ship = MoveInDirection(ship, 0, distance, true)
			fmt.Println("Moved the waypoint north")
			fmt.Println(ship)
		} else {
			ship = MoveInDirection(ship, 0, distance, false)
		}
	case "S":
		if ship.waypoint.exists {
			ship = MoveInDirection(ship, 180, distance, true)

		} else {
			ship = MoveInDirection(ship, 180, distance, false)
		}
	case "E":
		if ship.waypoint.exists {
			ship = MoveInDirection(ship, 90, distance, true)

		} else {
			ship = MoveInDirection(ship, 90, distance, false)
		}
	case "W":
		if ship.waypoint.exists {
			ship = MoveInDirection(ship, 270, distance, true)

		} else {
			ship = MoveInDirection(ship, 270, distance, false)
		}
	case "R":
		if ship.waypoint.exists {
			ship = RotateWaypoint(ship, distance)
		} else {
			ship = RotateShip(ship, distance)
		}
		fmt.Println(ship)
	case "L":
		if ship.waypoint.exists {
			ship = RotateWaypoint(ship, 360-distance)
		} else {
			ship = RotateShip(ship, 360-distance)
		}
		fmt.Println(ship)
	case "F":
		if ship.waypoint.exists {
			distanceToEast := distance * ship.waypoint.east_west
			distanceToNorth := distance * ship.waypoint.north_south

			ship = MoveInDirection(ship, 0, distanceToNorth, false)
			ship = MoveInDirection(ship, 90, distanceToEast, false)

			//ship = MoveInDirection(ship, 0, ship.north_south, true)
			//ship = MoveInDirection(ship, 90, ship.east_west, true)

			fmt.Println(ship)
		} else {
			ship = MoveInDirection(ship, ship.currentlyFacing, distance, false)
		}
		//fmt.Println(ship)
	}

	return ship
}

// MoveInDirection moves the ship forward based on direction and distance
func MoveInDirection(ship Ship, direction int, distance int, moveWaypoint bool) Ship {
	switch direction {
	case 0:
		if moveWaypoint {
			ship.waypoint.north_south += distance
		} else {
			ship.north_south += distance
		}
	case 90:
		if moveWaypoint {
			ship.waypoint.east_west += distance
		} else {
			ship.east_west += distance
		}
	case 180:
		if moveWaypoint {
			ship.waypoint.north_south -= distance
		} else {
			ship.north_south -= distance
		}
	case 270:
		if moveWaypoint {
			ship.waypoint.east_west -= distance
		} else {
			ship.east_west -= distance
		}
	}
	return ship
}

// RotateShip returns the rotated ship according to degrees
func RotateShip(ship Ship, degrees int) Ship {

	ship.currentlyFacing = int(math.Abs(float64(ship.currentlyFacing + degrees)))

	if ship.currentlyFacing >= 360 {
		ship.currentlyFacing = ship.currentlyFacing - 360
	}

	return ship
}

// RotateWaypoint returns the ship with a rotated waypoint according to degrees
func RotateWaypoint(ship Ship, degrees int) Ship {
	x := ship.waypoint.east_west
	y := ship.waypoint.north_south
	switch degrees {
	case 90:
		ship.waypoint.east_west = y
		ship.waypoint.north_south = -x
	case 180:
		ship.waypoint.east_west = -x
		ship.waypoint.north_south = -y
	case 270:
		ship.waypoint.east_west = -y
		ship.waypoint.north_south = x
	}

	return ship
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
