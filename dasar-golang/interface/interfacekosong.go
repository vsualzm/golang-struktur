package main

// import (
// 	"fmt"
// 	"strings"
// )

// type person struct {
// 	name string
// 	age  int
// }

// func main() {

// 	var datas interface{} = &person{name: "wick", age: 27}
// 	var name = datas.(*person).name
// 	fmt.Println(name)

// 	var person = []map[string]interface{}{
// 		{"name": "Wick", "age": 23},
// 		{"name": "Ethan", "age": 23},
// 		{"name": "Bourne", "age": 22},
// 	}

// 	for no, each := range person {
// 		no++
// 		fmt.Println(no, each["name"], "age is", each["age"])
// 	}

// 	var secret interface{}

// 	secret = "iqbal navisah"
// 	fmt.Println(secret)

// 	secret = []string{"apple", "mangga", "banana"}
// 	fmt.Println(secret)

// 	secret = 12.4
// 	fmt.Println(secret)

// 	// var data map[string]any

// 	// data = map[string]any{
// 	// 	"name":      "ethan hunt",
// 	// 	"grade":     2,
// 	// 	"breakfast": []string{"apple", "manggo", "banana"},
// 	// }

// 	var rahasia interface{}

// 	rahasia = 2
// 	var number = rahasia.(int) * 10
// 	fmt.Println(rahasia, "multiplied by 10 is :", number)

// 	rahasia = []string{"apple", "manggo", "banana"}
// 	var gruits = strings.Join(rahasia.([]string), ", ")
// 	fmt.Println(gruits, "is my favorite fruits")

// }
