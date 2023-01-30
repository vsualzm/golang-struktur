package main

import (
	"fmt"
	"strings"
)

func main() {
	names := []string{"ilham", "maulana"}
	printMessage("Helllo", names)

	fmt.Println("======")

	// rand.Seed(time.Now().Unix())
	// var randomValue int

	// randomValue = randomWithRange(2, 10)
	// fmt.Println("random number:", randomValue)
	// randomValue = randomWithRange(2, 10)
	// fmt.Println("random number:", randomValue)
	// randomValue = randomWithRange(2, 10)
	// fmt.Println("random number:", randomValue)
}

func printMessage(message string, arr []string) {
	nameString := strings.Join(arr, " ")
	fmt.Println(message, nameString)
}
