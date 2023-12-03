package main

import (
	"regexp"
	"strconv"
	"strings"
)

func P1(file []byte) int {

	sum := 0
	file_lines := strings.Split(string(file), "\n")

	for line_number, line := range file_lines {
		line_as_bytes := []byte(line)
		re := regexp.MustCompile("[0-9]+")
		indexes_found := re.FindAllIndex(line_as_bytes, -1)

		if indexes_found == nil {
			continue
		}

		for _, ind := range indexes_found {
			number := line[ind[0]:ind[1]]
			position := ind[0]
			size := len(number)
			is_part_number := determineIfIsPartNumber(position, size, line_number, file_lines)

			if is_part_number {
				num, _ := strconv.Atoi(number)
				sum += num
			}
		}
	}
	return sum
}

func determineIfIsPartNumber(position int, size int, line_number int, file_lines []string) bool {
	line_length := len(file_lines[line_number])

	// Check before
	if position > 0 {
		check := strings.Split(file_lines[line_number], "")[position-1]
		if check != "." {
			return true
		}
	}

	// Check after
	if position+size < line_length {
		check := strings.Split(file_lines[line_number], "")[position+size]
		if check != "." {
			return true
		}
	}

	start_at := 0
	if position > 0 {
		start_at = position - 1
	}

	finish_at := position + size + 1
	if position+size == line_length {
		finish_at = position + size
	}

	// Check above
	if line_number > 0 {
		previous_line := file_lines[line_number-1]
		splice := previous_line[start_at:finish_at]
		if findPart(splice) {
			return true
		}
	}

	// Check below
	if line_number < len(file_lines)-2 {
		next_line := file_lines[line_number+1]
		splice := next_line[start_at:finish_at]
		if findPart(splice) {
			return true
		}
	}

	return false
}

func findPart(splice string) bool {
	replaced := strings.ReplaceAll(splice, ".", "")
	re := regexp.MustCompile("[0-9]+")
	replaced = re.ReplaceAllString(replaced, "")

	return len(replaced) > 0
}
