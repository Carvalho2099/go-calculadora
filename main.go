package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Enter the operation (using this format: 2*2):")
	var input string
	fmt.Scan(&input)
	operation := strings.Split(input, "")
	result := getResult(operation)
	fmt.Println("The result is:", result)
}

func getResult(operation []string) int {
	num1, _ := strconv.Atoi(operation[0])
	num2, _ := strconv.Atoi(operation[2])

	switch operation[1] {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		return num1 / num2
	default:
		panic("Invalid operation")
	}
}
