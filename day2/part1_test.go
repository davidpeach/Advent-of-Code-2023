package main

import (
	"bufio"
	"bytes"
	"testing"
)

func TestP1(t *testing.T) {
	var buffer bytes.Buffer
	buffer.WriteString("Game 1: 10 green, 5 blue; 1 red, 9 green, 10 blue; 5 blue, 6 green, 2 red; 7 green, 9 blue, 1 red; 2 red, 10 blue, 10 green; 7 blue, 1 red\n") // Enough balls
	buffer.WriteString("Game 10: 99 green, 5 red, 3 blue; 4 blue, 7 green, 8 red; 9 blue, 4 green; 6 green, 3 red, 4 blue\n")                                          // Not Enough
	buffer.WriteString("Game 25: 2 green, 4 blue, 12 red; 14 blue, 9 red, 3 green; 3 red, 1 blue, 3 green; 6 red, 4 green, 2 blue; 6 blue, 12 red")                    // Enough balls

	want := 26

	scanner := bufio.NewScanner(&buffer)

	balls_available := make(map[string]int)
	balls_available["red"] = 12
	balls_available["green"] = 13
	balls_available["blue"] = 14

	actual := P1(scanner, balls_available)

	if actual != want {
		t.Fatalf("got %d, but wanted %d", actual, want)
	}

}
