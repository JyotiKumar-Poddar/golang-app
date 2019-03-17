package main

// function with variable argument

import (
	"fmt"
)

func main() {
	sum := variableArgument(1, 2, 3, 4, 5, 6)
	println(sum)
	fmt.Println("The sum of number is ",
		parameterWithVariableArgument("add", 1, 2, 3, 4, 5))
}

func variableArgument(numbers ...int) int {

	fmt.Println(numbers)
	total := 0
	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}
	return total
}

func parameterWithVariableArgument(operation string, numbers ...int) int {
	total := 0
	if operation == "add" {
		for i := 0; i < len(numbers); i++ {
			total += numbers[i]
		}
	}
	return total
}
