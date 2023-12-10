package main

import (
	"bufio"
	"slices"
	"strings"
)

func P1(scanner *bufio.Scanner) int {

	sum := 0

	// card_totals := make(map[string]int)

	for scanner.Scan() {
		card := scanner.Text()
		split := strings.Split(card, ":")
		// card_name := split[0]
		card_game := split[1]

		card_game_split := strings.Split(card_game, "|")
		my_numbers := card_game_split[0]
		winning_numbers := card_game_split[1]

		matching_numbers := 0
		my_numbers_sliced := removeDuplicate(strings.Split(my_numbers, " "))
		// fmt.Println(my_numbers_sliced)
		winning_numbers_sliced := strings.Split(winning_numbers, " ")
		// fmt.Println(winning_numbers_sliced)
		for _, my_num := range my_numbers_sliced {
			// fmt.Println(my_num)
			if slices.Contains(winning_numbers_sliced, my_num) {
				matching_numbers += 1
			}
		}
		// fmt.Println(matching_numbers)

		x := 0
		for i := 1; i <= matching_numbers; i++ {
			if x == 0 {
				x = 1
				continue
			}

			x *= 2

		}

		sum += x

	}

	return sum
}

func removeDuplicate[T string | int](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
