package simple

import "math"

// ищет все простые числа от 0 до N включительно

func SimpleCalc(n uint) []uint {

	var i, j uint
	simples := []uint{}

	for i = 2; i <= n; i++ {
		simple := true
		for j = 2; float64(j) <= math.Sqrt(float64(i)); j++ {
			if i%j == 0 {
				simple = false
			}
		}
		if simple {
			simples = append(simples, i)
		}
	}
	return simples
}
