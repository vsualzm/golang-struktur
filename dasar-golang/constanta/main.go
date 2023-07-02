package main

import "fmt"

func main() {
	// constanta
	// nilai yang tidak bisa kita ubah
	// nilai PASTI
	const firstname string = "haloo"
	fmt.Println(firstname)

	const data = "fix"
	fmt.Println(data)

	fmt.Print(firstname, data)

	const (
		name     = "ilham maulana"
		usia     = 22
		alamat   = "Jl Cicadas Pasar II"
		isLulus  = true
		floatNum = 2.4
	)

	fmt.Println(name)
	fmt.Println(usia)
	fmt.Println(alamat)
	fmt.Println(isLulus)
	fmt.Println(floatNum)

}
