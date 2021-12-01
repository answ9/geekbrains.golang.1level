package main

import (
	"fmt"
)

var mapFibbonachi map[uint]uint

func inputNum() {
	var num uint
	fmt.Print("Введите значение для расчета числа Фибоначии: ")
	fmt.Scan(&num)

	_, aExists := mapFibbonachi[num]

	// Проверяем есть ли запись карте
	if(aExists != true) {
		mapFibbonachi[num] = fibbonachi(num)
	}
	fmt.Println(mapFibbonachi[num])

	// рекурсивный вызов
	inputNum()
}

func fibbonachi(n uint) uint{

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibbonachi(n - 1) + fibbonachi(n - 2)
}

func main() {
	fmt.Println("Приветсвуем Тебя!")
	mapFibbonachi = make(map[uint]uint)
	inputNum()
}
