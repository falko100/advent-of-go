package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Game struct {
	id     int
	rounds []Round
}

type Round struct {
	red   int
	blue  int
	green int
}

func main() {
	startTime := time.Now()

	data, err := os.ReadFile("./inputs/day-2.txt")
	check(err)

	lines := strings.Split(string(data), "\n")

	var games []Game
	for _, line := range lines {
		// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
		var game Game
		game.id = toInt(regexp.MustCompile(`Game (\d+):`).FindStringSubmatch(line)[1])
		var roundsString = regexp.MustCompile(`Game \d+: (.+)`).FindStringSubmatch(line)[1]
		var roundsStringArray = strings.Split(roundsString, "; ")
		for _, roundString := range roundsStringArray {
			var round Round
			var roundStringArray = strings.Split(roundString, ", ")
			for _, colorString := range roundStringArray {
				var colorStringArray = strings.Split(colorString, " ")
				var color = colorStringArray[1]
				var amount = toInt(colorStringArray[0])
				switch color {
				case "red":
					round.red = amount
				case "green":
					round.green = amount
				case "blue":
					round.blue = amount
				}
			}
			game.rounds = append(game.rounds, round)
		}

		games = append(games, game)
	}

	totalRed := 12
	totalGreen := 13
	totalBlue := 14

	sum := 0
	sum2 := 0

	for _, game := range games {
		var highestRed = 0
		var highestGreen = 0
		var highestBlue = 0

		for _, round := range game.rounds {
			if round.red > highestRed {
				highestRed = round.red
			}
			if round.green > highestGreen {
				highestGreen = round.green
			}
			if round.blue > highestBlue {
				highestBlue = round.blue
			}
		}

		if highestRed <= totalRed && highestGreen <= totalGreen && highestBlue <= totalBlue {
			sum = sum + game.id
		}

		sum2 += highestRed * highestGreen * highestBlue
	}

	log.Printf("Sum: %d", sum)
	log.Printf("Sum2: %d", sum2)

	elapsed := time.Since(startTime)
	log.Printf("Day 2 took %s", elapsed)
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	check(err)
	return i
}
