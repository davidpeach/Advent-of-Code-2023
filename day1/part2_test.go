package main

import (
	"bytes"
	"testing"
)

func TestP2(t *testing.T) {
	var buffer2 bytes.Buffer
	buffer2.WriteString("aoneffthree4f")
	want := 14

	scanner := readFile(&buffer2)

	actual := P2(scanner)

	if actual != want {
		t.Fatalf("got %d, but wanted %d", actual, want)
	}
}
