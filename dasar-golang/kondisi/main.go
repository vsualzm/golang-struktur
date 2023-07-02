package main

import "fmt"

func main() {

	// case 1
	age := 10

	// nilaiKkm := 70
	// nilaiSiswa := 90

	if age < 21 {
		fmt.Println("data benar") // data benar
	} else {
		fmt.Println("data belum benar")
	}

	// case 2
	value := 7
	switch value {
	case 10:
		fmt.Println("perfect")
	case 8:
		fmt.Println("Great")
	case 7:
		fmt.Println("Good")
	case 5:
		fmt.Println("Bad")
	default:
		fmt.Println("No value")
	}

	// case 3
	var cek = 8
	switch {
	case cek == 8:
		fmt.Println("perfect")
	case (cek < 8) && (cek > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	// case 4
	var point = 8840.0
	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	// case 5
	count := 6
	if count == 10 {
		fmt.Println("hello")
	} else if count >= 7 {
		fmt.Println("cukup")
	} else if count >= 5 {
		fmt.Println("di bawah kkm")
	} else {
		fmt.Println("input salah")
	}
}
