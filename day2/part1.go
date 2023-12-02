package main

import (
	"bufio"
	"strconv"
	"strings"
)

func P21(scanner *bufio.Scanner) int {
	balls_available := make(map[string]int)
	balls_available["red"] = 12
	balls_available["green"] = 13
	balls_available["blue"] = 14

	sum := 0

out:
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), ":")
		game_name := split[0]
		rounds_played := split[1]
		game_split := strings.Split(game_name, " ")
		game_id, err := strconv.Atoi(game_split[1])

		if err != nil {
			panic(err)
		}

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

				if balls_available[colour] == 0 {
					continue out
				}

				for available_colour, available_count := range balls_available {
					if colour == available_colour && available_count < count {
						continue out
					}
				}
			}
		}

		sum += game_id
	}

	return sum
}
