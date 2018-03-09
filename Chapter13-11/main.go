//Buffer Channel
package main

import (
	"fmt"
)

func main() {
	// 3 = number of receiver
	ch := make(chan int,3)
	go func() {ch<-1} ()
	ch <- 2
	// capacity of Channel
	fmt.Println("cap:", cap(ch))
	// lenge of channel (used..)
	fmt.Println("len:", len(ch))
}