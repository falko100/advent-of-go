package main

import (
	"log"
	"os"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	startTime := time.Now()

	data, err := os.ReadFile("./inputs/day-1.1.txt")
	check(err)

	data = append(data, '\n')
	dataString := string(data)
	// search replace one to 1, two to 2, etc
	replacements := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	for word, digit := range replacements {
		dataString = strings.ReplaceAll(string(dataString), word, digit)
	}

	var lines = strings.Split(dataString, "\n")

	var result []int
	for _, line := range lines {
		var firstNumberFound bool = false
		var firstNumber int
		var lastNumber int
		for _, b := range line {
			if b < '0' || b > '9' {
				continue
			}

			n := int(b - '0')
			if !firstNumberFound {
				firstNumber = n
				firstNumberFound = true
			}

			lastNumber = n
		}

		result = append(result, firstNumber*10+lastNumber)
	}

	var sum int
	for _, n := range result {
		sum += n
	}

	println(sum)
	elapsed := time.Since(startTime)
	log.Printf("Day 1.1 took %s", elapsed)
}
