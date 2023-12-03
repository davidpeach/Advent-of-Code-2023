package main

import (
	"os"
	"testing"
)

func TestP1(t *testing.T) {
	file, _ := os.ReadFile("./test_input_p1")
	want := 925
	actual := P1(file)
	if actual != want {
		t.Fatalf("got %d, but wanted %d", actual, want)
	}
}
