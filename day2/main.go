package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./input")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	balls_available := make(map[string]int)
	balls_available["red"] = 12
	balls_available["green"] = 13
	balls_available["blue"] = 14

	sum := P1(scanner, balls_available)

	fmt.Println(fmt.Sprintf("Answer: %d", sum))
}
