package main

import (
	"testing"
)

func TestGreet(t *testing.T) {
	emptyResult := greet("")
	if (emptyResult != "Hello Jenkins") {
		t.Errorf("func greet() failed: Expected %v got %v", "Hello Jenkins", emptyResult)
	}

	result := greet("World")
	if (result != "Hello World!") {
		t.Errorf("func greet() failed: Expected %v got %v", "Hello World!", result)
	}
}