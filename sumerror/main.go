package main

import (
	"fmt"
	"strconv"
)

//Sum converts two strings to int and returns their sum
func Sum(a, b string) (int64, error) {
	if _, err := strconv.ParseInt(a, 10, 64); err != nil {
		return 0, err
	}
	i, _ := strconv.ParseInt(a, 10, 64)

	if _, err := strconv.ParseInt(b, 10, 64); err != nil {
		return 0, err
	}
	v, _ := strconv.ParseInt(b, 10, 64)
	return i + v, nil
}

func main() {
	i, err := Sum("aaaava", "5")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(i)
}
