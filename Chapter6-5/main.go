package main

import "fmt"
import "strconv"

func main() {
	for number :=1; number <=100 ; number++ {
		modnumber(number)
	}
}

func modnumber(input int) {
	if input%15 == 0 {
		a := "a"
		printoutput(a)
	}
	if input%3 == 0 {
		a := "b"
		printoutput(a)
	}
	if input%5 ==0 {
		a := "c"
		printoutput(a)
	} else {
		a := strconv.Itoa(input)
		printoutput(a)
	}
}

func printoutput(n string) {
	switch n {
	case "a":
		fmt.Println("FizzBuzz")
	case "b":
		fmt.Println("Fizz")
	case "c":
		fmt.Println("Buzz")
	default:
		fmt.Println(n)
	}
}