package main

import (
	"fmt"
	"lesson10/calc"
	"lesson10/fibonachi"
	"lesson10/simple"
)

func main() {

	// test work task1
	fmt.Println("Simples 0-100: ", simple.SimpleCalc(100))

	// test work task2
	fmt.Println("Fibonachi for 10 without cache: ", fibonachi.Fib(10, false))
	fmt.Println("Fibonachi for 10 with cache: ", fibonachi.Fib(10, true))

	// test work task3
	fmt.Println(calc.Calc("+", 2, 2))
}
