package main

import (
	"testing"
)

func TestGreet(t *testing.T) {
	emptyResult := greet("")
	if (emptyResult != "Hello Stranger") {
		t.Errorf("func greet() failed")
	}

	result := greet("World")
	if (result != "Hello World!") {
		t.Errorf("func greet() failed")
	}
}