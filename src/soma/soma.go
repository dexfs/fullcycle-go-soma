package main

import (
	"fmt"
)

func soma(num1, num2 int) int {
	return num1 + num2
}

func main() {
	result := soma(5, 5)
	fmt.Println(result)
}
