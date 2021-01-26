package main

import (
    "fmt"
)

func main() {
	fmt.Println("This is a Jenkins Demo")
	greetEmpty := greet("")
	fmt.Println(greetEmpty)
	greetNotEmpty := greet("World")
	fmt.Println(greetNotEmpty)
}