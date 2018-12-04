package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines := readFileString()
	sort.Strings(lines)

	//guard id -> { minute -> count }
	guards := make(map[string]map[int]int)

	currentId := ""
	startSleepTime := ""

	for _, line := range lines {
		if strings.Contains(line, "#") {
			// guard #XYZ going on duty
			// start tracking
			guardIdArray := strings.Split(line, "#")
			guardIdArray = strings.Split(guardIdArray[1], " ")
			guardId := guardIdArray[0]
			currentId = guardId
			if _, ok := guards[guardId]; !ok {
				guards[guardId] = make(map[int]int)
			}
		} else if strings.Contains(line, "falls asleep") {
			// mark start time for currentId sleep
			elements := strings.Split(line, "]")
			elements = strings.Split(elements[0], "[")
			startSleepTime = elements[1]
		} else if strings.Contains(line, "wakes up") {
			elements := strings.Split(line, "]")
			elements = strings.Split(elements[0], "[")
			endSleepTime := elements[1]
			minutesAsleep := calculateSleepMinutes(startSleepTime, endSleepTime) // [5, 6, 7, 8]

			for _, min := range minutesAsleep {
				guards[currentId][min]++
			}
		}
	}

	maxMinSlept := 0
	guardIdMaxMinSlept := ""
	var minSleptMost int
	minSleptMostCount := 0

	fmt.Println(guards)

	for id, minuteMap := range guards {
		// count minutes slept
		totalMinutesSlept := 0
		for _, times := range minuteMap {
			totalMinutesSlept += times
		}
		if totalMinutesSlept > maxMinSlept {
			maxMinSlept = totalMinutesSlept
			guardIdMaxMinSlept = id
			// also figure out what minute they slept most
			for min2, times2 := range minuteMap {
				if times2 > minSleptMostCount {
					minSleptMost = min2
					minSleptMostCount = times2
				}
			}
		}
	}

	fmt.Print("maxMinSlept: ")
	fmt.Println(maxMinSlept)
	fmt.Print("guardIdMaxMinSlept: ")
	fmt.Println(guardIdMaxMinSlept)
	fmt.Print("minSleptMost: ")
	fmt.Println(minSleptMost)
	fmt.Print("ANSWER PART 1: ")
	guardIdMaxMinSleptInt, _ := strconv.Atoi(guardIdMaxMinSlept)
	fmt.Println(guardIdMaxMinSleptInt * minSleptMost)

	// part 2
	var minSleptMost2 int
	countMinSleptMost2 := 0
	var guardIdMinSleptMost2 string

	for id, minuteMap := range guards {
		// figure out which minute they slept the most
		for minute, times := range minuteMap {
			if times > countMinSleptMost2 {
				minSleptMost2 = minute
				countMinSleptMost2 = times
				guardIdMinSleptMost2 = id
			}
		}
	}

	fmt.Print("minSleptMost2: ")
	fmt.Println(minSleptMost2)
	fmt.Print("countMinSleptMost2: ")
	fmt.Println(countMinSleptMost2)
	fmt.Print("guardIdMinSleptMost2: ")
	fmt.Println(guardIdMinSleptMost2)
	guardIdMinSleptMost2Int, _ := strconv.Atoi(guardIdMinSleptMost2)
	fmt.Print("ANSWER PART2: ")
	fmt.Println(guardIdMinSleptMost2Int * minSleptMost2)
}

func calculateSleepMinutes(start, end string) []int {
	startTime := strings.Split(start, " ")
	endTime := strings.Split(end, " ")
	startTimes := strings.Split(startTime[1], ":")
	startHour, _ := strconv.Atoi(startTimes[0])
	startMin, _ := strconv.Atoi(startTimes[1])
	endTimes := strings.Split(endTime[1], ":")
	endHour, _ := strconv.Atoi(endTimes[0])
	endMin, _ := strconv.Atoi(endTimes[1])

	var minsAsleep []int

	if startHour < endHour {
		// get the hours
		for counterHours := startHour; counterHours < endHour; counterHours++ {
			for min := startMin; min < 60; min++ {
				minsAsleep = append(minsAsleep, min)
			}
		}
		// get leftover minutes
		for min := 0; min < endMin; min++ {
			minsAsleep = append(minsAsleep, min)
		}
	} else {
		for min := startMin; min < endMin; min++ {
			minsAsleep = append(minsAsleep, min)
		}
	}

	return minsAsleep
}

func readFileString() []string {
	pwd, _ := os.Getwd()
	content, err := ioutil.ReadFile(pwd + "/day4/input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")

	return lines
}
