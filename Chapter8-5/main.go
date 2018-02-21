//pass by slice value
package main

import "fmt"

func main() {
	slice := []int{1,2,3}
	double(slice)
	fmt.Printf("origin add %p\n", slice)
	fmt.Printf("original %v\n", slice)
}

func double(nums []int) {
	fmt.Printf("double addr %p\n", nums)

	// result in changing original value
	for i := 0; i < len(nums); i++ {
		nums[i] *= 2
	}
	fmt.Println(nums)
}