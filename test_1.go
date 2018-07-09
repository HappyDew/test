package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter a number:")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * (1 / 3.2808)

	fmt.Println(output)
}
