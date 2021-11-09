package main

import (
	"fmt"
	"math"
)

func main() {

	var numPoint int
	fmt.Println("Приветсвуем Тебя!")
	fmt.Println("1. Вычислить площадь прямоугольника")
	fmt.Println("2. Вычислить диаметр и длину окружности по заданной площади круга")
	fmt.Println("3. Разбить число трехзначное число на разряды")
	fmt.Println(" Выбери пункт(1, 2, 3)?")
	fmt.Scanf("%d\n", &numPoint)

	switch numPoint {
	case 1:
		var a, b float32
		fmt.Println("Вычисление площадь прямоугольника!")
		fmt.Println("Введите сторону a?")
		fmt.Scanf("%f\n", &a)
		fmt.Println("Введите сторону b?")
		fmt.Scanf("%f\n", &b)
		fmt.Printf("Площадь прямоугольника равна - %f\n", a*b)

	case 2:
		var S float64
		fmt.Println("Вычислиние диаметр и длину окружности по заданной площади круга!")
		fmt.Println("Введите площадь круга S?")
		fmt.Scanf("%f\n", &S)
		fmt.Printf("Диаметр круга равен - %f\n", 2 * math.Sqrt(S / math.Pi))
		fmt.Printf("Длина окружности равна - %f\n", 2 * math.Sqrt(S / math.Pi) * math.Pi)

	case 3:
		var num int
		fmt.Println("Разбиение числа трехзначное число на разряды!")
		fmt.Println("Введите трехзначное число?")
		fmt.Scanf("%d\n", &num)
		if num > 0 || num < 1000 {
			fmt.Printf("Количество сотен - %d\n", num / 100)
			fmt.Printf("Количество десятков - %d\n", num / 10 % 10)
			fmt.Printf("Количество единиц - %d\n", num - ((num / 10) * 10))
		} else {
			fmt.Println("Введено не корректное число!")
		}

	default:
		fmt.Println("Выбран не корректный пункт!")
	}
}