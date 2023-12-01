package main

import (
	"bytes"
	"testing"
)

func TestP1(t *testing.T) {
	var buffer bytes.Buffer
	buffer.WriteString("f1a2ke\nc4sv\nd5at5a")
	// 12, 44, 55
	want := 111

	scanner := readFile(&buffer)

	actual := P2(scanner)

	if actual != want {
		t.Fatalf("got %d, but wanted %d", actual, want)
	}
}
