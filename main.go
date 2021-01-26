package main

import (
    "fmt"
	"flag"
	"strings"
)

func main() {
	var filename string
	var errorMessage string
	flag.StringVar(&filename, "filename", "default.txt", "default txt file")
	flag.StringVar(&errorMessage, "error", "FAIL", "default error message")
	flag.Parse()
	fmt.Printf("filename is: %v", filename)
	fmt.Println("This is a Jenkins Demo")
	greetEmpty := greet("")
	fmt.Println(greetEmpty)
	greetNotEmpty := greet("World")
	fmt.Println(greetNotEmpty)
	if (strings.Contains(errorMessage, "FAIL")) {
		sendEmail(filename, errorMessage)
	}
}