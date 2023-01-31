package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	var allStudents = []person{
		{name: "ilham", age: 21},
		{name: "ucok", age: 22},
		{name: "repal", age: 23},
	}

	for _, student := range allStudents {
		fmt.Println(student.name, "age is", student.age)
	}

	// anon

	// var allStudents = []struct {
	// 	person
	// 	grade int
	// }{
	// 	{person: person{"wick", 21}, grade: 2},
	// 	{person: person{"ethan", 22}, grade: 3},
	// 	{person: person{"bond", 21}, grade: 3},
	// }

	// for _, student := range allStudents {
	// 	fmt.Println(student)
	// }

}
