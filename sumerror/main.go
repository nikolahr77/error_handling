package main

import (
	"fmt"
	"strconv"
)

//Sum converts two strings to int and returns their sum
func Sum(a, b string) (int64, error) {
	i, err := strconv.ParseInt(a, 10, 64)
	if err != nil {
		return 0, err
	}

	v, err := strconv.ParseInt(b, 10, 64)
	if err != nil {
		return 0, err
	}

	return i + v, nil
}

func main() {
	i, err := Sum("4", "5")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(i)
}
