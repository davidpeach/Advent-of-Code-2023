package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("./input")

	scanner := bufio.NewScanner(file)
	sum := P1(scanner)

	fmt.Println("Answer is ", sum)
}
