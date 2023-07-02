package main

import (
	"errors"
	"fmt"
)

func main() {

	scores := []int{10, 5, 8, 9, 7}

	total := sum(scores)
	rata2 := total / len(scores)

	fmt.Println("isi jumlah dalam array", total)
	fmt.Println("rata2:", rata2)

	result, err := calculate(10, 20, "-")
	if err != nil {
		fmt.Println("terjadi kesalahan")
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}

	// result, err := calculate(10, 2, "+")
	// result, err := calculate(10, 2, "-")
	// result, err := calculate(10, 2, "*")
	// result, err := calculate(10, 2, "/")
	// result, err := calculate(10, 2, "=")
}

func sum(numbers []int) int {
	var total int
	for _, number := range numbers {
		total = total + number
	}

	return total
}

func calculate(number, numberTwo int, operation string) (int, error) {
	var result int
	var errorResult error

	switch operation {
	case "+":
		result = number + numberTwo
	case "-":
		result = number - numberTwo
	case "*":
		result = number * numberTwo
	case "/":
		result = number / numberTwo
	default:
		errorResult = errors.New("non operation")
	}

	return result, errorResult
}
