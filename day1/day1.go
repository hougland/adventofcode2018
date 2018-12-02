package main

import (
	"fmt"
	"io"
	"os"
)

func main()  {
	pwd, _ := os.Getwd()
	numbers := readFile(pwd + "/day1/input.txt")

	foundRepeat := false
	m := make(map[int]bool)
	total := 0
	m[total] = true

	for foundRepeat != true {
		for _, num := range numbers {
			total += num

			if _, ok := m[total]; ok { // key found
				foundRepeat = true
				fmt.Print(num)
				break
			} else { // key not found
				m[total] = true
			}
		}
	}

	fmt.Print("\n")
	fmt.Print(total)
}

func readFile(filePath string) (numbers []int) {
	fd, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", filePath, err))
	}
	var line int
	for {
		_, err := fmt.Fscanf(fd, "%d\n", &line)

		if err != nil {
			fmt.Println(err)
			if err == io.EOF {
				return
			}
			panic(fmt.Sprintf("Scan Failed %s: %v", filePath, err))

		}
		numbers = append(numbers, line)
	}
	return
}
