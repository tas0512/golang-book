//Buffer Channel - overfilled
package main

import (
	"fmt"
)

func main() {
	// 2 = number of receiver
	c := make(chan int,2)
	c <-1
	c <-2
	//c <-3 // overfilled the buffer
	
	fmt.Println(<-c)
	fmt.Println(<-c)

	c <-4
	fmt.Println(<-c)
}