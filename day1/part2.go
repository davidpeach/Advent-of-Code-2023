package main

import (
	"bufio"
	"strconv"
	"strings"
	"unicode"
)

func P2(scanner *bufio.Scanner) int {

	var word_map = make(map[string]int)
	word_map["one"] = 1
	word_map["two"] = 2
	word_map["three"] = 3
	word_map["four"] = 4
	word_map["five"] = 5
	word_map["six"] = 6
	word_map["seven"] = 7
	word_map["eight"] = 8
	word_map["nine"] = 9

	var sum = 0

	for scanner.Scan() {
		var number_as_string string
		var word_string strings.Builder
		var prepend_string strings.Builder
		var word = scanner.Text()

		for _, ch := range word {
			var found_word_string string

			if unicode.IsDigit(ch) {
				number_as_string = string(ch)
				break
			}

			word_string.WriteString(string(ch))

			found_word_string = findWordString(word_string.String(), word_map)
			if found_word_string != "" {
				number := word_map[found_word_string]
				number_as_string = strconv.Itoa(number)
				break
			}
		}

		word_string.Reset()

		for _, ch := range reverse(scanner.Text()) {
			var found_word_string string

			if unicode.IsDigit(ch) {
				number_as_string += string(ch)
				break
			}

			prepend_string.WriteString(string(ch))
			prepend_string.WriteString(word_string.String())
			word_string = prepend_string

			prepend_string.Reset()

			found_word_string = findWordString(word_string.String(), word_map)
			if found_word_string != "" {
				number := word_map[found_word_string]
				number_as_string += strconv.Itoa(number)
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

func findWordString(find string, words map[string]int) string {
	for word := range words {
		if strings.Contains(find, word) {
			return word
		}
	}

	return ""
}
