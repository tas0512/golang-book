//Leak
package main

import (
	"fmt"
)

func main() {
	c:=make(chan int,2)
	go func() { c <- 1}()
	go func() { c <- 2}()
	//go func() { c <- 3}()
	c <- 5
	fmt.Println(<-c)
	fmt.Println(<-c)
}