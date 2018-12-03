package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	claims := readFileString()

	claimsMap := make(map[Point]int)

	for _, claim := range claims {
		for x := 0; x < claim.width; x++ {
			for y := 0; y < claim.height; y++ {
				var point Point
				point.x = claim.x + x
				point.y = claim.y + y

				claimsMap[point] += 1
			}
		}
	}

	overlap := 0
	for _, val := range claimsMap {
		if val >= 2 {
			overlap++
		}
	}

	// part 1
	fmt.Print("Overlap count: ")
	fmt.Println(overlap)

	// part 2
	for _, claim := range claims { // for each claim - was it overlapping?
		claimOverlaps := false

		for x := 0; x < claim.width; x++ {
			for y := 0; y < claim.height; y++ {
				var point Point
				point.x = claim.x + x
				point.y = claim.y + y

				if claimsMap[point] != 1 {
					claimOverlaps = true
				}
			}
		}

		if !claimOverlaps {
			fmt.Println("ID: " + claim.id)
		}
	}
}

func readFileString() []Claim {
	pwd, _ := os.Getwd()
	content, err := ioutil.ReadFile(pwd + "/day3/input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")

	var claims []Claim

	for _, line := range lines {
		claim := parseClaim(line)
		claims = append(claims, claim)
	}

	return claims
}

func parseClaim(line string) Claim {
	var claim Claim

	id := ""
	finishedId := false
	x := ""
	finishedX := false
	y := ""
	finishedY := false
	width := ""
	finishedWidth := false
	height := ""

	for index, char := range line {
		if index == 0 {
			continue
		}
		if string(char) == "@" {
			continue
		}

		if !finishedId { // finding id
			if string(char) == " " {
				finishedId = true
			} else {
				id += string(char)
			}
		} else if string(char) == " " {
			continue
		} else if !finishedX {
			if string(char) == "," {
				finishedX = true
			} else {
				x += string(char)
			}
		} else if !finishedY {
			if string(char) == ":" {
				finishedY = true
			} else {
				y += string(char)
			}
		} else if !finishedWidth {
			if string(char) == "x" {
				finishedWidth = true
			} else {
				width += string(char)
			}
		} else {
			height += string(char)
		}
	}

	claim.id = id
	xInt, _ := strconv.Atoi(x)
	claim.x = xInt
	yInt, _ := strconv.Atoi(y)
	claim.y = yInt
	widthInt, _ := strconv.Atoi(width)
	claim.width = widthInt
	heightInt, _ := strconv.Atoi(height)
	claim.height = heightInt

	return claim
}

type Claim struct {
	id string
	x int
	y int
	width int
	height int
}

type Point struct {
	x int
	y int
}
