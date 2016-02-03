package main

import "fmt"

func main() {
	var f = 100.0
	var c = 0.0
	var result = 0.00

	fmt.Print("Please enter Fahrenheit temperature? ")
	fmt.Scan(&f)

	c = (((f - 32) * 5) / 9)

	result = float64(int(c*100)) / 100

	fmt.Println("The temperature in Celsius is", result)
	// fmt.Println(result)
	// not very elegant, but im looking up the stupid stuff still
}
