package main

import "fmt"

func main() {
	var name string
	fmt.Print("Whats your name? ")
	fmt.Scan(&name)
	fmt.Println("}{ello " + name)
}
