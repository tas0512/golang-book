// go-routine
//Using Add / Defer / Waiting

package main

import (
	"sync"
	"fmt"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	
	for i:=0; i<3; i++ {
		go func(n int) {
			defer wg.Done()
			for i:=0 ; i<10 ;i++ {
				fmt.Println(n,":",i)
			}
		} (i)
	}
	wg.Wait()
	fmt.Println("Finished")
}