package main

import (
	"os"
	"log"
	"testing"
)

func TestGreet(t *testing.T) {

	file, err := os.OpenFile("logs.txt", os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0666)
	if err != nil {
		//Fatal is equivalent to Print() followed by a call to os.Exit(1).
		log.Fatal(err)
	}

	log.SetOutput(file)

	emptyResult := greet("")
	if (emptyResult != "Hello Stranger") {
		t.Errorf("func greet() failed: Expected %v got %v", "Hello Stranger", emptyResult)
		log.Printf("func greet() failed: Expected %v got %v", "Hello Stranger", emptyResult)
	} else {
		log.Printf("func greet() passes when input is empty")
	}

	result := greet("World")
	if (result != "Hello World!") {
		t.Errorf("func greet() failed: Expected %v got %v", "Hello World!", result)
		log.Printf("func greet() failed: Expected %v got %v", "Hello World!", result)
	} else {
		log.Printf("func greet() passes when input is not empty")
	}
	defer file.Close()

    log.Print("-----------------------------------------------------------")
}