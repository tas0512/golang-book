package main

import "fmt"
import "math/rand"
import "time"

func main() {
	
	var input int
	rand.Seed(time.Now().UTC().UnixNano())
	random := rand.Intn(100)
	fmt.Println("** random# is ",random)

	for i := 0; i < 10; i++ {
		if i>4 {
			fmt.Println("Maximum Count, Good Bye")
			return
		}
		fmt.Print("Input#",i+1,": ")
		fmt.Scanf("%d\n", &input)

		if input>random {
			fmt.Print("Your input is greater than random number\n")
		} else if input<random {
			fmt.Print("Your input is less than random number\n")
		} else if input == random {
			fmt.Print("Correct!")
			return
		}
	}
}