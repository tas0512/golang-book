//Channel - Direction (Pipeline#3)
package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(naturals, squares)
	print(squares)
}

func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(in <-chan int, out chan<-int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func print(out <-chan int) {
	for x := range out {
		fmt.Println(x)
	}
}