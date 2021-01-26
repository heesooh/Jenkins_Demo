package main

import (
    "fmt"
	"flag"
)

func main() {
	var filename string
	flag.StringVar(&filename, "filename", "default.txt", "default txt file")
	flag.Parse()
	fmt.Printf("filename is: %v", filename)
	fmt.Println("This is a Jenkins Demo")
	greetEmpty := greet("")
	fmt.Println(greetEmpty)
	greetNotEmpty := greet("World")
	fmt.Println(greetNotEmpty)
	sendEmail(filename)
}