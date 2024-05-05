package main

import (
	"strings"
	"testing"
)

func TestEcho(t *testing.T) {
	arg := []string{"Hello", "World"}
	want := "Hello World"

	res := strings.Join(arg, " ")

	if res != want {
		t.Fatalf("echo(['Hello', 'World'] = %q, want match for %#q", res, want)
	}
}
