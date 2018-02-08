package main

import "fmt"

func main() {
	fmt.Print("Enter a number in Farenheit: ")
	var input float64
	fmt.Scanf("%f", &input)
	//output := input * 2
	//convert to Celsius
	output := ((input - 32)/9) * 5
	//convert to 2 decimal digits
	fmt.Printf("%.2f", output)
}