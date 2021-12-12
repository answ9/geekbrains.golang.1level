package calc

import (
	"fmt"
	"os"
)

func isInt(a float64) (bool) {
	return a == float64(int(a))
}

func Calc(operation string, operand1, operand2 float64) (result float64) {	
	switch operation {
	case "+":
		result = operand1 + operand2
	case "-":
		result = operand1 - operand2
	case "*":
		result = operand1 * operand2
	case "/":
		if operand2 == 0 {
			fmt.Print("Деление на 0 запрещено!\n")
			os.Exit(1)
		} else {
			result = operand1 / operand2
		}
	case "%":
		if isInt(operand1) && isInt(operand2) {
			result = float64(int(operand1) % int(operand2))
		} else {
			fmt.Print("Вы ввели не целые числа!\n")
			os.Exit(1)
		}
	default:
		fmt.Print("Введена неизвестная операция\n")
		os.Exit(1)
	}
	return result
}
