package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	var numPoint int
	var res string
	fmt.Println("Приветсвуем Тебя!")
	fmt.Println("1. Вычислить площадь прямоугольника")
	fmt.Println("2. Вычислить диаметр и длину окружности по заданной площади круга")
	fmt.Println("3. Разбить число трехзначное число на разряды")
	fmt.Println("4. Поиск простого числа")
	fmt.Println("5. Простые арифметические операции двух чисел (+, -, *, /)")
	fmt.Print(" Выбери пункт(1, 2, 3, 4, 5): ")
	fmt.Scanln(&numPoint)

	switch numPoint {
	case 1:
		var a, b float32
		fmt.Println("Вычисление площадь прямоугольника!")
		fmt.Print("Введите сторону a: ")
		fmt.Scanln(&a)
		fmt.Println("Введите сторону b: ")
		fmt.Scanln(&b)
		fmt.Printf("Площадь прямоугольника равна - %f\n", a*b)

	case 2:
		var S float64
		fmt.Println("Вычислиние диаметр и длину окружности по заданной площади круга!")
		fmt.Print("Введите площадь круга S: ")
		fmt.Scanf("%f\n", &S)
		fmt.Printf("Диаметр круга равен - %f\n", 2 * math.Sqrt(S / math.Pi))
		fmt.Printf("Длина окружности равна - %f\n", 2 * math.Sqrt(S / math.Pi) * math.Pi)

	case 3:
		var num int
		fmt.Println("Разбиение числа трехзначное число на разряды!")
		fmt.Print("Введите трехзначное число?")
		fmt.Scan(&num)
		if num > 0 || num < 1000 {
			fmt.Printf("Количество сотен - %d\n", num / 100)
			fmt.Printf("Количество десятков - %d\n", num / 10 % 10)
			fmt.Printf("Количество единиц - %d\n", num - ((num / 10) * 10))
		} else {
			fmt.Println("Введено не корректное число!")
		}

	case 4:
		var num int
		var count int
		fmt.Println("Поиск простого числа!")
		fmt.Print("Введите значение верхнего предела диапазона поиска N: ")
		fmt.Scan(&num)
		for i := 2; i < num; i++ {
			count = 0
			for j := 1; j < num; j++ {
				if(i % j == 0) {
					count++
				}
			}

			if(count < 3) {
				res += strconv.Itoa(i) + " "
			}
		}

		fmt.Print(res)

	case 5:
		var a, b, res float32
		var op string

		fmt.Print("Введите первое число: ")
		fmt.Scanln(&a)

		fmt.Print("Введите второе число: ")
		fmt.Scanln(&b)

		fmt.Print("Введите арифметическую операцию (+, -, *, /): ")
		fmt.Scanln(&op)

		switch op {
		case "+":
			res = a + b
		case "-":
			res = a - b
		case "*":
			res = a * b
		case "/":
			res = a / b
		default:
			fmt.Println("Операция выбрана неверно")
			os.Exit(1)
		}

		fmt.Printf("Результат выполнения операции: %f\n", res)

	default:
		fmt.Println("Выбран не корректный пункт!")
		os.Exit(1)
	}
}