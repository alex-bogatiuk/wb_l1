package main

import "fmt"

func main() {

	num1, num2 := 38, 62
	num1, num2 = num2, num1
	fmt.Printf("num1: %d, num2: %d", num1, num2)

}
