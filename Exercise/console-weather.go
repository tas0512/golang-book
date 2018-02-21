
package main

import "fmt"

func main() {

	weatherCelsius(25,"Bangkok few cloud")
	weatherCelsius(34,"Tak sunny")
	weatherCelsius(17,"Phuket rainy")
}

func weatherCelsius(number int, message string) {
	slice := make([]string,16)
	slice2 := make([]string,16)
	slice3 := make([]string,16)

	if number == 25 {
		//number 2 (first digit)
		slice[1] = "_"
		slice2[1] = "_"
		slice2[2] = "|"
		slice3[0] = "|"
		slice3[1] = "_"

		//number 5 (second digit)
		slice[6] = "_"
		slice2[3] = "|"
		slice2[4] = "_"
		slice3[4] = "_"
		slice3[5] = "|"

		slice3[7] = "C"

		fmt.Println(slice)
		fmt.Println(slice2)
		fmt.Println(slice3)

		fmt.Println(message)
	}

	if number==34 {
		//display 3
		fmt.Printf("_\n_|\n_|")

		//display 4
		fmt.Printf("|_|\n |")
	}
}