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

type WayPoint struct{
	point1 string // E+ W-
	point1_distance int // E+ W-
	point2 string // S+ N-
	point2_distance int // S+ N-
}

var directionsE = [4]string{"E","S","W","N"}
var directionsS = [4]string{"S","W","N","E"}
var directionsW = [4]string{"W","N","E","S"}
var directionsN = [4]string{"N","E","S","W"}


func main() {
	direction_distances, _ := readLines("Day12 input.txt")
	var ship = Ship{current_direction:"E",vertical:"E",vertical_distance:0, horizontal:"N",horizontal_distance:0 }
	var wayPoint = WayPoint{point1:"E",point1_distance:10,point2:"N", point2_distance:1}
	for _, direction_distance := range direction_distances {
		var direction = string(direction_distance[0])
		distance, _ := strconv.Atoi(direction_distance[1:])
		switch string(direction_distance[0]){
		case "E","S","W","N":
			wayPoint = wayPointmoveToDirection(direction,distance,wayPoint)
		case "F":
			ship = moveShip(wayPoint, ship, distance)
		case "L","R":
			wayPoint = turnWayPointDirection(direction,distance,wayPoint)
		}
		fmt.Println(direction_distance)
		fmt.Println(ship)
		fmt.Println(wayPoint)
	}

	totalDistance := math.Abs(float64(ship.horizontal_distance)) + math.Abs(float64(ship.vertical_distance))
	fmt.Println(totalDistance)
}

func wayPointmoveToDirection(direction string, distance int, wayPoint WayPoint) (WayPoint){
	switch direction {
	case "E","W":
		switch wayPoint.point1{
		case "E","W":
			if wayPoint.point1 == direction{
				wayPoint.point1_distance += distance
			}else{
				wayPoint.point1_distance -= distance
			}
			if wayPoint.point1_distance<0{
				wayPoint.point1 = direction
				wayPoint.point1_distance = int(math.Abs(float64(wayPoint.point1_distance)))
			}
		}
		switch wayPoint.point2{
		case "E","W":
			if wayPoint.point2 == direction{
				wayPoint.point2_distance += distance
			}else{
				wayPoint.point2_distance -= distance
			}
			if wayPoint.point2_distance<0{
				wayPoint.point2 = direction
				wayPoint.point2_distance = int(math.Abs(float64(wayPoint.point2_distance)))
			}
		}
	case "S","N":
		switch wayPoint.point1{
		case "S","N":
			if wayPoint.point1 == direction{
				wayPoint.point1_distance += distance
			}else{
				wayPoint.point1_distance -= distance
			}
			if wayPoint.point1_distance<0{
				wayPoint.point1 = direction
				wayPoint.point1_distance = int(math.Abs(float64(wayPoint.point1_distance)))
			}
		}
		switch wayPoint.point2{
		case "S","N":
			if wayPoint.point2 == direction{
				wayPoint.point2_distance += distance
			}else{
				wayPoint.point2_distance -= distance
			}
			if wayPoint.point2_distance<0{
				wayPoint.point2 = direction
				wayPoint.point2_distance = int(math.Abs(float64(wayPoint.point2_distance)))
			}
		}
	}
	return wayPoint
}

func moveShip(wayPoint WayPoint, ship Ship, distance int) (Ship){
	ship = moveToDirection(wayPoint.point2, wayPoint.point2_distance*distance, ship)
	ship = moveToDirection(wayPoint.point1, wayPoint.point1_distance*distance, ship)
	return ship
}

func moveToDirection(direction string, distance int, ship Ship) (Ship){
	switch direction{
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

func turnWayPointDirection(direction string, degress int, wayPoint WayPoint) (WayPoint){
	wayPoint.point1 = turnDirection(wayPoint.point1, direction, degress)
	wayPoint.point2 = turnDirection(wayPoint.point2, direction, degress)
	return wayPoint
}

func turnDirection(current_direction, to_direction string, degress int) (string){
	var numOfDirection int = degress/90
	var directions [4] string
	var turnedDirection string
	switch current_direction {
	case "E":
		directions = directionsE
	case "S":
		directions = directionsS
	case "W":
		directions = directionsW
	case "N":
		directions = directionsN
	}
	switch to_direction {
	case "R":
		turnedDirection = directions[numOfDirection]
	case "L":
		turnedDirection = directions[4-numOfDirection]
	}
	return turnedDirection
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