
package main

import (
	"time"
	"fmt"
)

func main() {
	volumn := 200

	start := time.Now()

	container := order(volumn)
	for _, cup := range container {
		fmt.Println(cup)
	}
	end := time.Now()
	fmt.Println(end.Sub(start))

}

func order(volumn int) (container []string) {
	for i:=1; i<volumn; i++ {
		coffee := receiveOrder(i)
		coffee = brew(coffee)
		container = append(container, serve(coffee))
	}
	return
}

func receiveOrder(number int) string {
	time.Sleep(5 * time.Millisecond)
	return fmt.Sprintf("oder: %d", number)
}

func brew(order string) string {
	time.Sleep(100 * time.Millisecond)
	return fmt.Sprintf("%s %s", order, "espresso")
}

func serve(coffee string) string {
	time.Sleep(100 * time.Millisecond)
	return fmt.Sprintf("%s %s", coffee, "ready :)")
}