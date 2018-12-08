package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func main() {
	steps := readFileString()
	result := part1(steps)

	// step 2 - process the steps
	numWorkers := 5
	letterToSecondsRequired := getLetterToSeconds()
	seconds := 0
	inProcess := make(map[string]int)
	toProcess := strings.Split(result, "")
	var processed []string

	for len(processed) < len(result) {
		for i := 1; i <= numWorkers - len(inProcess); i++ { // for each available worker, see if it can do work
			for _, letter := range toProcess { // check all the letters to be processed to see if we can work on one
				if preReqsMet(letter, processed, steps) && !(inProcess[letter] > 0) && !(contains(processed, letter)) {
					inProcess[letter] = 1 // start processing
				}
				if len(inProcess) >= numWorkers {
					break
				}
			}
		}
		for letter, secondsProcessed := range inProcess {
			if secondsProcessed == letterToSecondsRequired[letter] { // off by 1 error? finished processing
				delete(inProcess, letter) // remove from inProcess
				processed = append(processed, letter) // add to processed
			} else {
				inProcess[letter] += 1
			}
		}
		seconds++
	}

	fmt.Print("Seconds: ")
	fmt.Println(seconds)
}
// too high: 947
// too high: 1267

func immediatePreReqsMet(letter, result string, steps []Step) bool {
	for _, step := range steps {
		if letter == step.prerequisite && !strings.Contains(result, step.letter) {
			return false
		}
	}
	return true
}

// if results contains all prereq
func preReqsMet(letter string, result []string, steps []Step) bool {
	for _, step := range steps {
		if step.prerequisite == letter && !contains(result, step.letter) {
			return false
		}
	}
	return true
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func getLetterToSeconds() map[string]int {
	secondsMap := make(map[string]int)
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphabetSlice := strings.Split(alphabet, "")
	for i := 0; i < len(alphabetSlice); i++ {
		letter := alphabetSlice[i]
		secondsMap[letter] = 60 + i + 1
	}
	return secondsMap
}

func findAllConnections(letter string, steps []Step, lastLetter string) []string {
	var connections []string
	for _, step := range steps {
		if step.letter == letter && step.prerequisite != lastLetter {
			connections = append(connections, step.prerequisite)
		}
	}
	return connections
}

// not mentioned as a prerequisite for any step
func getFirstLetters(steps []Step) []string {
	letterMentioned := make(map[string]bool)
	for _, step := range steps {
		letterMentioned[step.prerequisite] = true
	}
	var lastLetters []string
	for _, step := range steps {
		if !letterMentioned[step.letter] {
			lastLetters = append(lastLetters, step.letter)
		}
	}
	return lastLetters
}
// if not mentioned as a letter in any step
func getLastLetter(steps []Step) string {
	letters := make(map[string]bool)
	for _, step := range steps {
		letters[step.letter] = true
	}
	var firstLetter string
	for _, step := range steps {
		if !letters[step.prerequisite] {
			firstLetter = step.prerequisite
			break
		}
	}
	return firstLetter
}

func readFileString() []Step {
	pwd, _ := os.Getwd()
	content, _ := ioutil.ReadFile(pwd + "/day7/input.txt")
	lines := strings.Split(string(content), "\n")

	var steps []Step
	for _, line := range lines {
		step := lineToStep(line)
		steps = append(steps, step)
	}

	return steps
}

// Step C must be finished before step A can begin.
func lineToStep(line string) Step {
	words := strings.Split(line, " ")
	letter1 := ""
	letter2 := ""
	for _, word := range words {
		if len(word) == 1 {
			if letter1 == "" {
				letter1 = word
			} else {
				letter2 = word
			}
		}
	}

	return Step{letter1, letter2}
}

type Step struct {
	letter string
	prerequisite string
}

func getUniqueLetters(letters []string) []string {
	uniqueLetters := make(map[string]bool)
	for _, letter := range letters {
		uniqueLetters[letter] = true
	}
	var keys []string
	for k := range uniqueLetters {
		keys = append(keys, k)
	}
	return keys
}

func part1(steps []Step) string {
	lastLetter := getLastLetter(steps)
	firstLetter := getFirstLetters(steps)
	fmt.Println("lastLetter: " + lastLetter)
	fmt.Print("firstLetters: ")
	fmt.Println(firstLetter)

	result := ""

	unlocked := firstLetter

	for len(unlocked) > 0 {
		// sort and remove dups from unlocked
		unlocked = getUniqueLetters(unlocked)
		sort.Sort(sort.StringSlice(unlocked))

		// pop first element off of "unlocked" and add to result
		letter := unlocked[0]
		unlocked = unlocked[1:]
		result += letter

		connectionsToLetter := findAllConnections(letter, steps, lastLetter)
		// find connections whose reqs are met already
		var connectionsWithPreReqsMet []string
		for _, connection := range connectionsToLetter {
			if immediatePreReqsMet(connection, result, steps) {
				connectionsWithPreReqsMet = append(connectionsWithPreReqsMet, connection)
			}
		}
		// then add all to unlocked
		unlocked = append(unlocked, connectionsWithPreReqsMet...)
	}

	result += lastLetter
	return result
}
