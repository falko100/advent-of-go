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
		"one":   "one1one",
		"two":   "two2two",
		"three": "three3three",
		"four":  "four4four",
		"five":  "five5five",
		"six":   "six6six",
		"seven": "seven7seven",
		"eight": "eight8eight",
		"nine":  "nine9nine",
	}

	for word, digit := range replacements {
		dataString = strings.ReplaceAll(dataString, word, digit)
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
