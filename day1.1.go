package main

import (
	"log"
	"os"
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

	var lines []string
	start := 0

	data = append(data, '\n')

	for i, b := range data {
		if b != '\n' {
			continue
		}

		if i == start {
			start = i + 1
			continue
		}

		lines = append(lines, string(data[start:i]))
		start = i + 1
	}

	var result []int
	for _, line := range lines {
		var firstNumber *int
		var lastNumber *int
		for _, b := range line {
			if b < '0' || b > '9' {
				continue
			}

			n := int(b - '0')
			if firstNumber == nil {
				firstNumber = &n
			}

			lastNumber = &n
		}

		result = append(result, *firstNumber*10+*lastNumber)
	}

	var sum int
	for _, n := range result {
		sum += n
	}

	println(sum)
	elapsed := time.Since(startTime)
	log.Printf("Day 1.1 took %s", elapsed)
}
