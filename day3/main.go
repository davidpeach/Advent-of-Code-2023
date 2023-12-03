package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadFile("./input")

	if err != nil {
		panic(err)
	}

	// scanner := bufio.NewScanner(file)

	sum := P1(file)

	fmt.Println(sum)
}
