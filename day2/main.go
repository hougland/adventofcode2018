package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	lines := readFileString()

	var twos []string
	var threes []string

	for _, line := range lines {
		chars := strings.Split(line, "")

		counter := make(map[string]int)
		for _, char := range chars {
			counter[char]++
		}

		for key, value := range counter {
			if value == 2 {
				if !contains(twos, key) {
					twos = append(twos, key)
					fmt.Println("adding " + key + " to twos")
				} else {
					fmt.Println("NOT adding " + key + " to twos")
				}
			} else if value == 3 {
				if !contains(threes, key) {
					threes = append(threes, key)
					fmt.Println("adding " + key + " to threes")
				} else {
					fmt.Println("NOT adding " + key + " to threes")
				}
			}
		}
	}

	fmt.Println(twos)
	fmt.Println(threes)
	fmt.Print(len(twos))
	fmt.Print("*")
	fmt.Print(len(threes))
	fmt.Print("=\n")
	fmt.Print((len(twos)) * (len(threes)))
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

// wrong: 416
// wrong: 459
// wrong: 17631