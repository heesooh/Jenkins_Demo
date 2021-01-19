package main

import (
    "fmt"
)

func main() {
	greetEmpty := greet("")
	fmt.Println(greetEmpty)
	greetNotEmpty := greet("World")
	fmt.Println(greetNotEmpty)
}