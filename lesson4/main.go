package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	ar :=  []int64{}
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		num, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		ar = append(ar, num)
	}

	InsertionSort(ar)
	fmt.Println(ar)
}

func InsertionSort(ar []int64) {
	for i := 1; i < len(ar); i++ {
		x := ar[i]
		j := i
		for ; j >= 1 && ar[j-1] > x; j-- {
			ar[j] = ar[j-1]
		}
		ar[j] = x
	}
}
