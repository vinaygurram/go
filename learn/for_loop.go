package main

import "fmt"

func main() {
	fmt.Println("For loop example")
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("Total sum is", sum)

	fmt.Println("For loop example 2")
	for ; sum > 0; {
		sum = sum -10
	}
	fmt.Println("Total sum is", sum)

	sum = 342
	for sum > 0 {
		sum = sum -10
	}
	fmt.Println("Sum is ", sum)

	/***
	 Infinite loop syntax
	 for {}
	 */
}
