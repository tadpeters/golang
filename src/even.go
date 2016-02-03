package main

import "fmt"

func main() {

	for counter := 0; counter < 100; counter++ {
		if counter%2 == 0 {
			fmt.Println(counter, "is even")
		}
	}
}
