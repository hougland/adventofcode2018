package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	lines := readFileString()

	// part 1
	var twos []string
	var threes []string

	for _, line := range lines {
		chars := strings.Split(line, "")

		counter := make(map[string]int)
		for _, char := range chars {
			counter[char]++
		}

		for _, value := range counter {
			if value == 2 && !contains(twos, line) {
				twos = append(twos, line)
			} else if value == 3 && !contains(threes, line) {
				threes = append(threes, line)
			}
		}
	}

	fmt.Print(len(twos))
	fmt.Print("*")
	fmt.Print(len(threes))
	fmt.Print("=")
	fmt.Println((len(twos)) * (len(threes)))

	// part 2
	mostSimilarChars := 0
	var mostSimilarCharsStringLine1 string
	var mostSimilarCharsStringLine2 string

	for index1, line1 := range lines {
		for index2, line2 := range lines {
			if index2 <= index1 { // skip current and previous lines
				continue
			}

			similarCharsInLine := 0
			for charIndex, char := range strings.Split(line1, "") {
				if char == string(line2[charIndex]) {
					similarCharsInLine++
				}
			}

			if similarCharsInLine > mostSimilarChars {
				mostSimilarChars = similarCharsInLine
				mostSimilarCharsStringLine1 = line1
				mostSimilarCharsStringLine2 = line2
			}
		}
	}

	fmt.Println(mostSimilarCharsStringLine1)
	fmt.Println(mostSimilarCharsStringLine2)
	// manually grab the answer
}

func readFileString() []string {
	pwd, _ := os.Getwd()
	content, err := ioutil.ReadFile(pwd + "/day2/input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	return lines
}

func contains(chars []string, target string) bool {
	for _, char := range chars {
		if char == target {
			return true
		}
	}
	return false
}
