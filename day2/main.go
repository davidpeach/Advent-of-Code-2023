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

	sum := P21(scanner)

	fmt.Println(fmt.Sprintf("Answer: %d", sum))
}
