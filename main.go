package main

import (
    "fmt"
	"flag"
)

func main() {
	var filename string
	flag.StringVar(&filename, "filename", "log.txt", "default txt file")
	fmt.Println("This is a Jenkins Demo")
	greetEmpty := greet("")
	fmt.Println(greetEmpty)
	greetNotEmpty := greet("World")
	fmt.Println(greetNotEmpty)
	sendEmail(filename)
}