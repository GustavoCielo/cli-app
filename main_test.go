package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestGreetCommand(t *testing.T) {
	os.Args = []string{"", "greet"}
	main()
	// Add checks to verify output or any other side effects.
}

func TestGreetFormalCommand(t *testing.T) {
	os.Args = []string{"", "greet", "formal", "--name=World"}
	// Capture output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	main()
	// Stop capturing output
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	io.Copy(&buf, r)
	expected := "Hi there!\nHow do you do, World?"
	if buf.String() != expected {
		t.Fatalf("Expected \"%s\" but got \"%s\"", expected, buf.String())
	}
}
