package main

import "fmt"

func main() {
	amount := 5
	double(&amount)
	fmt.Printf("original %v\n", amount)
}

func double(number *int) {
	// this will effect the original value at func main()
	*number *=2
	fmt.Println(number)
}