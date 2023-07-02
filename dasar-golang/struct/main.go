package main

import "fmt"

type Student struct {
	Name string
	Umur int
}

func main() {

	// cara pemanggilann

	// var s1 Student

	// s1.Name = "ilham"
	// s1.Umur = 12

	// fmt.Println("NAMA:", s1.Name)
	// fmt.Println("UMUR:", s1.Umur)

	// memakai pointer
	var s1 = Student{Name: "wick", Umur: 2}

	var s2 *Student = &s1
	fmt.Println("student 1, name :", s1.Name)
	fmt.Println("student 4, name :", s2.Name)

	fmt.Println("============================")

	s2.Name = "ethan"
	fmt.Println("student 1, name :", s1.Name)
	fmt.Println("student 4, name :", s2.Name)

}
