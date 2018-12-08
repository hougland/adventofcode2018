package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readFileString()
	originalCoordinates := stringsToCoordinates(input)
	//fmt.Println(originalCoordinates)

	grid := buildGrid(originalCoordinates)
	fmt.Print("height of grid (y): ")
	fmt.Println(len(grid))
	fmt.Print("width of grid (x): ")
	fmt.Println(len(grid[0]))
	grid = addCoordsToGrid(grid, originalCoordinates)
	//printGrid(grid)

	addedLastRound := originalCoordinates

	// this loop is taking FOREVERRRRRRRR - at least 15 min so far - killed overnight
	// could I terminate some trains of execution early? like if count is below max count?
	for len(addedLastRound) > 0 {
		var coordinatesToAdd []Coordinate

		for _, coord := range addedLastRound {
			adding := getCoordsToAddToGrid(coord, grid, coord.value)
			coordinatesToAdd = append(coordinatesToAdd, adding...)
		}

		grid = addCoordsToGrid(grid, coordinatesToAdd)
		//printGrid(grid)

		addedLastRound = coordinatesToAdd
		//fmt.Print("addedLastRound: ")
		//fmt.Println(addedLastRound)
	}

	printGrid(grid)

	// doesn't take into account spaces that are equidistant between main points
	//for index, row := range grid {
	//	for index2, value := range row {
	//		if value == "" {
	//			grid[index][index2] = "."
	//		}
	//		// find if this space is equidistant ??????
	//	}
	//}

	// TODO: optimization - could remove this and keep track in loop above
	// go through grid and count areas
	nonInfiniteCoords := getNonInfiniteCoords(originalCoordinates) // value -> count
	for _, row := range grid {
		for _, value := range row {
			nonInfiniteCoords[value] ++
		}
	}

	maxArea := 0
	for _, value := range nonInfiniteCoords {
		if value > maxArea {
			maxArea = value
		}
	}

	fmt.Printf("MAX AREA: %d", maxArea)
}

func getCoordsToAddToGrid(coordinate Coordinate, grid [][]string, value string) []Coordinate {
	var coordinatesToAdd []Coordinate
	// first make sure it can fit in grid
	if coordinate.y < len(grid) - 2 {
		targetUp := Coordinate{coordinate.x, coordinate.y + 1, value}
		if targetUp.x >= 354 {
			fmt.Println("1111111111111")
		}
		equal := allDirectionsFreeOrEqual(targetUp, grid, value)
		if equal {
			if grid[targetUp.y][targetUp.x] != value {
				coordinatesToAdd = append(coordinatesToAdd, Coordinate{targetUp.x, targetUp.y, value})
			}
			//grid[targetUp.y][targetUp.x] = value
		}
	}
	if coordinate.y != 0 {
		targetDown := Coordinate{coordinate.x, coordinate.y - 1, value}
		if targetDown.x >= 354 {
			fmt.Println("222222222222222")
		}
		equal := allDirectionsFreeOrEqual(targetDown, grid, value)
		if equal {
			if grid[targetDown.y][targetDown.x] != value {
				coordinatesToAdd = append(coordinatesToAdd, Coordinate{targetDown.x, targetDown.y, value})
			}
			//grid[targetDown.y][targetDown.x] = value
		}
	}
	if coordinate.x < len(grid[0]) - 2 {
		targetLeft := Coordinate{coordinate.x + 1, coordinate.y, value}
		equal := allDirectionsFreeOrEqual(targetLeft, grid, value)
		if equal {
			if grid[targetLeft.y][targetLeft.x] != value {
				coordinatesToAdd = append(coordinatesToAdd, Coordinate{targetLeft.x, targetLeft.y, value})
			}
			//grid[targetLeft.y][targetLeft.x] = value
		}
	}
	if coordinate.x != 0 {
		targetRight := Coordinate{coordinate.x - 1, coordinate.y, value}
		if targetRight.x >= 354 {
			fmt.Println("444444444444444")
		}
		equal := allDirectionsFreeOrEqual(targetRight, grid, value)
		if equal {
			if grid[targetRight.y][targetRight.x] != value {
				coordinatesToAdd = append(coordinatesToAdd, Coordinate{targetRight.x, targetRight.y, value})
			}
			//grid[targetRight.y][targetRight.x] = value
		}
	}
	return coordinatesToAdd
}

func allDirectionsFreeOrEqual(coordinate Coordinate, grid [][]string, value string) bool {
	// check that current coordinate isn't in the grid already?
	//if coordinate.x >= 354 {
	//	fmt.Println("WTFFFFFFFF")
	//}
	return upFree(coordinate, grid, value) && downFree(coordinate, grid, value) && leftFree(coordinate, grid, value) && rightFree(coordinate, grid, value)
}
func upFree(coordinate Coordinate, grid [][]string, value string) bool {
	if coordinate.y + 1 == len(grid) { // outside grid
		return true
	}
	val := grid[coordinate.y+1][coordinate.x]
	return val == "" || val == value
}
func downFree(coordinate Coordinate, grid [][]string, value string) bool {
	if coordinate.y - 1 < 0 { // outside grid
		return true
	}
	testing := grid[coordinate.y-1]
	//if coordinate.x >= 354 {
	//	fmt.Println("WTFFFFFFFF")
	//}
	val := testing[coordinate.x]
	return val == "" || val == value
}
func leftFree(coordinate Coordinate, grid [][]string, value string) bool {
	if coordinate.x + 1 == len(grid[0]) { // edge of grid
		return true
	}
	val := grid[coordinate.y][coordinate.x + 1]
	return val == "" || val == value
}
func rightFree(coordinate Coordinate, grid [][]string, value string) bool {
	if coordinate.x - 1 < 0 { // outside grid
		return true
	}
	val := grid[coordinate.y][coordinate.x - 1]
	return val == "" || val == value
}

func stringsToCoordinates(lines []string) []Coordinate {
	var coordsList []Coordinate
	for index, line := range lines {
		coords := strings.Split(line, ", ")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])

		var coord Coordinate
		coord.x = x
		coord.y = y
		coord.value = strconv.Itoa(index)

		coordsList = append(coordsList, coord)
	}
	return coordsList
}

func buildGrid(coordinates []Coordinate) [][]string {
	// get x and y dimensions
	// no negative numbers in input
	maxX := 0
	maxY := 0
	for _, coord := range coordinates {
		if coord.x > maxX {
			maxX = coord.x
		}
		if coord.y > maxY {
			maxY = coord.y
		}
	}

	grid := make([][]string, maxY + 1)
	for i := range grid {
		grid[i] = make([]string, maxX + 1)
	}

	return grid
}

func addCoordsToGrid(grid [][]string, coordinates []Coordinate) [][]string {
	for _, coordinate := range coordinates {
		grid[coordinate.y][coordinate.x] = coordinate.value
	}
	return grid
}

func getNonInfiniteCoords(coordinates []Coordinate) map[string]int {
	nonInfininite := make(map[string]int)
	for index, coord1 := range coordinates {
		topBound := false
		bottomBound := false
		leftBound := false
		rightBound := false
		for _, coord2 := range coordinates {
			if coord1.y < coord2.y {
				topBound = true
			}
			if coord1.y > coord2.y {
				bottomBound = true
			}
			if coord1.x > coord2.x {
				rightBound = true
			}
			if coord1.x < coord2.x {
				leftBound = true
			}
		}
		if topBound && bottomBound && leftBound && rightBound {
			coordinates[index] = coord1
			nonInfininite[coord1.value] = 0
		}
	}
	return nonInfininite
}

type Coordinate struct {
	x int
	y int
	value string
}

func printGrid(grid [][]string) {
	fmt.Println("printing grid:")
	for _, row := range grid {
		printRow(row)
	}
}

func printRow(row []string) {
	fmt.Print("[ ")
	for _, value := range row {
		if value != "" {
			fmt.Print(value)
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Println(" ]")
}

func readFileString() []string {
	pwd, _ := os.Getwd()
	content, err := ioutil.ReadFile(pwd + "/day6/input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	return lines
}

//func getCoordsFromGrid(grid [][]string) []Coordinate {
//	var coordinates []Coordinate
//	for index, row := range grid {
//		for index2 := range row {
//			coordinates = append(coordinates, Coordinate{index, index2})
//		}
//	}
//	return coordinates
//}