package main

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestChangeDir(t *testing.T) {
	path := "..\\"
	newPath, err := filepath.Abs(path)
	want, _ := filepath.Abs("D:\\Go")

	state.setState(newPath)

	if state.getState() != want || err != nil {
		t.Fatalf(`changeDir("..\") = %q, %v, want match for %#q, nil`, state.getState(), err, want)
	}
}

func TestEcho(t *testing.T) {
	arg := []string{"Hello", "World"}
	want := "Hello World"

	res := strings.Join(arg, " ")

	if res != want {
		t.Fatalf("echo(['Hello', 'World'] = %q, want match for %#q", res, want)
	}
}
