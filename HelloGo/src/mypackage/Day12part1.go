package main

import (
    "bufio"
    "fmt"
	"os"
	"strconv"
	"math"
)

type Ship struct{
	current_direction string
	vertical string // E+ W-
	vertical_distance int
	horizontal string // S+ N-
	horizontal_distance int
}

var directionsE = [4]string{"E","S","W","N"}
var directionsS = [4]string{"S","W","N","E"}
var directionsW = [4]string{"W","N","E","S"}
var directionsN = [4]string{"N","E","S","W"}


func main() {
	direction_distances, _ := readLines("Day12 input.txt")
	var ship = Ship{current_direction:"E",vertical:"E",vertical_distance:0, horizontal:"N",horizontal_distance:0 }
	for _, direction_distance := range direction_distances {
		var direction = string(direction_distance[0])
		distance, _ := strconv.Atoi(direction_distance[1:])
		switch string(direction_distance[0]){
		case "E","S","W","N":
			ship = moveToDirection(direction,distance,ship)
		case "F":
			ship = moveToDirection(ship.current_direction, distance, ship)
		case "L","R":
			ship = turnDirection(direction,distance,ship)
		}
		fmt.Println(ship)
	}

	totalDistance := math.Abs(float64(ship.horizontal_distance)) + math.Abs(float64(ship.vertical_distance))
	fmt.Println(totalDistance)
}

func turnDirection(direction string, degress int, ship Ship) (Ship){
	var numOfDirection int = degress/90
	var directions [4] string
	switch ship.current_direction {
	case "E":
		directions = directionsE
	case "S":
		directions = directionsS
	case "W":
		directions = directionsW
	case "N":
		directions = directionsN
	}
	switch direction {
	case "R":
		ship.current_direction = directions[numOfDirection]
	case "L":
		ship.current_direction = directions[4-numOfDirection]
	}
	return ship
}

func moveToDirection(direction string, distance int, ship Ship) (Ship){
	switch direction {
	case "E":
		ship.vertical_distance += distance
	case "S":
		ship.horizontal_distance += distance
	case "W":
		ship.vertical_distance -= distance
	case "N":
		ship.horizontal_distance -= distance
	}
	return ship
}

func readLines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}