package main

import "fmt"

func main() {
	fmt.Println(add(1,2,3))

	xs := []int{1,2,3}
	fmt.Println(add(xs...))
}

func add(arg ...int) int {
	total := 0
	for _, v := range arg {
		total += v
	}
	return total
}