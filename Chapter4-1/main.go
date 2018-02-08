package main

import "fmt"

func main(){
	v1, v2 := "first", "sec"
	//swapping
	v1, v2 = v2, v1

	fmt.Println(v1)
	fmt.Println(v2)

	// Long Declaration
	var x string = "Hello, World"
	fmt.Println(x)

	var y string
	y = "Hello, World"
	fmt.Println(y)

	// Short Declaration
	// Type Interface
	z := "Hello, World"
	fmt.Println(z)
	fmt.Printf("Type: %T\n", z)

	const xx string = "Hello, World"
	// cannot assign value to const 
	//xx = "Other string"

	var (
		a = 5
		b = 10
		c = 15
	)
	println(a,b,c)
}