package main

import (
	"bufio"
	"strconv"
	"strings"
)

func P2(scanner *bufio.Scanner) int {
	sum := 0

	for scanner.Scan() {

		minimum_balls_needed := make(map[string]int)
		split := strings.Split(scanner.Text(), ":")
		rounds_played := split[1]
		rounds := strings.Split(rounds_played, ";")

		for _, round := range rounds {

			balls := strings.Split(round, ",")

			for _, picked := range balls {

				picked_trimmed := strings.Trim(picked, " ")

				picked_split := strings.Split(picked_trimmed, " ")
				count, err := strconv.Atoi(picked_split[0])

				if err != nil {
					panic(err)
				}

				colour := picked_split[1]

				test := minimum_balls_needed[colour]

				if count > test {
					minimum_balls_needed[colour] = count
				}
			}
		}

		value := 1
		for _, count := range minimum_balls_needed {
			value *= count
		}

		sum += value
	}

	return sum
}
