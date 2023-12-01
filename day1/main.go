package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./input")

	if err != nil {
		panic(err)
	}

	scanner := readFile(file)

	// sum := P1(scanner)
	sum := P2(scanner)

	// fmt.Println(fmt.Sprintf("Part 1 answer: %d", sum))
	fmt.Println(fmt.Sprintf("Part 2 answer: %d", sum))
}

func readFile(reader io.Reader) *bufio.Scanner {
	scanner := bufio.NewScanner(reader)
	return scanner
}

func reverse(str string) string {
	rune_array := []rune(str)
	var reversed []rune
	for i := len(rune_array) - 1; i >= 0; i-- {
		reversed = append(reversed, rune_array[i])
	}
	return string(reversed)
}
