package main

import "fmt"

//Sum calls Check and returns the sum of two integers
func Sum(a, b int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovering")
		}
	}()
	Check(a)
	return a + b
}

//Check checks if a is greater than 10
func Check(a int) {
	if a > 10 {
		panic(fmt.Sprintf("%v", a))
	}
}

func main() {
	Sum(5, 10)
}
