package main

import (
	"fmt"
	"os"
	"strconv"
)

func multiple(number int, multipleNumber int) bool {
	return (number % multipleNumber) == 0
}

func sumOfMultiplesOF3And5BelowNumber(number int) int {
	sum := 0

	for i := 1; i < number; i++ {
		if multiple(i, 3) || multiple(i, 5) {
			sum += i
		}
	}

	return sum
}

func main() {
	arg, _ := strconv.Atoi(os.Args[1])

	fmt.Println(sumOfMultiplesOF3And5BelowNumber(arg))
}
