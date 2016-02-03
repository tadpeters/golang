package main

import (
	"fmt"
)

func main() {
	greet("first", "last")
	greet("game", "chess")

}

func greet(firstName, lastName string) string {
	var fullname string
	fullname = (firstName + " " + lastName)
	fmt.Println(fullname)
	return fullname
}

// func recivers name (params) returns {}
