package main

import (
	"bufio"
	"strconv"
	"unicode"
)

func P1(scanner *bufio.Scanner) int {

	var sum = 0

	for scanner.Scan() {
		var number_as_string string
		for _, ch := range scanner.Text() {
			if unicode.IsDigit(ch) {
				number_as_string = string(ch)
				break
			}
		}

		for _, ch := range reverse(scanner.Text()) {
			if unicode.IsDigit(ch) {
				number_as_string += string(ch)
				break
			}
		}

		if len(number_as_string) < 2 {
			break
		}

		as_int, err := strconv.Atoi(number_as_string)

		if err != nil {
			panic(err)
		}

		sum += as_int
	}

	return sum
}
