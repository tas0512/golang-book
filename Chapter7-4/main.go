package main

//copy slice
import "fmt"

func main() {
	slice := []int{1,2,3}
	fmt.Println(slice)
	newSlice := make([]int,2)
	
	fmt.Println(newSlice)
	// copy newSlice to slice (replace first 2[])
	copy(slice, newSlice)

	fmt.Printf("slice: %v\n", slice)
	fmt.Printf("slice: %v\n", newSlice)
}