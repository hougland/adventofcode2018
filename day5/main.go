package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	input := readFileString()[0]

	// part 2
	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	unitToCount := make(map[string]int) // letter -> resulting polymer length

	for _, letter := range alphabet {
		fmt.Print("current")
		regex := regexp.MustCompile("(?i)" + letter)
		strippedLine := regex.ReplaceAllString(input, "")
		polymer := breakdownUnits(strings.Split(strippedLine, ""))
		unitToCount[letter] = len(polymer)
	}
	
	shortestLength := 100000
	for _, value := range unitToCount {
		if value < shortestLength {
			shortestLength = value
		}
	}

	fmt.Println(shortestLength)
}

// part 1
func breakdownUnits(lines []string) []string {
	brokeDown := true
	for brokeDown {
		brokeDown = false

		for index := 0; index < len(lines) - 1; index++ {
			if isReactive(lines[index], lines[index + 1]) {
				lines = remove(lines, index)
				brokeDown = true
			}
		}
	}
	return lines
}

func isReactive(element1, element2 string) bool {
	if element1 == strings.ToLower(element1) {
		// lowercase must match uppercase
		return element2 == strings.ToUpper(element1)
	} else if element1 == strings.ToUpper(element1) {
		return element2 == strings.ToLower(element1)
	}

	return false
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+2:]...)
}

func readFileString() []string {
	pwd, _ := os.Getwd()
	content, err := ioutil.ReadFile(pwd + "/day5/input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")

	return lines
}
