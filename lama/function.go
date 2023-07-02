package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// // func main() {
// // 	names := []string{"ilham", "maulana"}
// // 	printMessage("hello", names)
// // }

// // func printMessage(message string, arr []string) {
// // 	nameString := strings.Join(arr, " ")
// // 	fmt.Println(message, nameString)
// // }

// func main() {
// 	rand.Seed(time.Now().Unix())
// 	var randomValue int

// 	randomValue = randomWithRange(2, 1000)
// 	fmt.Println("random number :", randomValue)
// 	randomValue = randomWithRange(2, 1000)
// 	fmt.Println("random number :", randomValue)
// 	randomValue = randomWithRange(2, 1000)
// 	fmt.Println("random number :", randomValue)
// }

// func randomWithRange(min, max int) int {
// 	var value = rand.Int()%(max-min+1) + min
// 	return value
// }

// // func main() {
// //     divideNumber(10, 2)
// //     divideNumber(4, 0)
// //     divideNumber(8, -4)
// // }

// // func divideNumber(m, n int) {
// //     if n == 0 {
// //         fmt.Printf("invalid divider. %d cannot divided by %d\n", m, n)
// //         return
// //     }

// //     var res = m / n
// //     fmt.Printf("%d / %d = %d\n", m, n, res)
// // }
