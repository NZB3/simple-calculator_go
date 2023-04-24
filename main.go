package main

import (
	"fmt"
)

func calculate(number1 float64, number2 float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return number1 + number2, nil

	case "-":
		return number1 - number2, nil

	case "*":
		return number1 * number2, nil

	case "/":
		if number2 == 0.0 {
			return 0, fmt.Errorf("can't divide '%f' by zero", number1)
		}
		return number1 / number2, nil
	default:
		return 0, fmt.Errorf("operator was not recognized")
	}
}

func main() {
	var number1, number2 float64
	var operator string

	fmt.Print("Enter the first number: ")
	fmt.Scanln(&number1)

	fmt.Print("Enter the second number: ")
	fmt.Scanln(&number2)

	fmt.Print("Enter the operator ( + - * /): ")
	fmt.Scanln(&operator)

	result, err := calculate(number1, number2, operator)

	fmt.Printf("%f %s %f = %v", number1, operator, number2, result)
	if err != nil {
		fmt.Println(err)
	}
}
