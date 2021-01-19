package main

import (
    "fmt"
)

func greet(user string) string {
    if (len(user) == 0) {
        return "Hello Stranger"
    } else {
        return fmt.Sprintf("Hello %v!", user)
    }
}