package main

import (
	"fmt"
)

func main() {
	x, y := f()
	fmt.Println(x, y)

	z := f1(x,y)
	fmt.Println(z)
}

func f() (int, int) {
	return 5, 7
}

func f1(a int, b int) int {
	c := a+b
	return c
}