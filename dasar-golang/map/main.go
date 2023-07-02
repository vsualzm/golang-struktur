package main

import "fmt"

func main() {

	// map
	// key dan value

	barang := map[string]int{}

	barang["kunci"] = 1
	barang["laptop"] = 2
	barang["mouse"] = 3

	fmt.Println(barang) // map[kunci:1 laptop:2 mouse:3]

	item := map[string]int{
		"laptop":   20,
		"keyboard": 30,
		"charger":  43,
		"terminal": 40,
	}

	// cek
	for i, items := range item {
		if i == "laptop" {
			fmt.Println("see")
			continue
		}

		if i == "charger" && items == 43 {
			delete(item, "charger")
			continue
		}

		fmt.Println(i, items)
	}

	bebek := map[string]int{"januari": 50, "februari": 40}
	value, isExist := bebek["januari"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}

	fmt.Println("================================")

	keranjang := map[string]int{
		"BCA":     10,
		"BNI":     23,
		"MANDIRI": 43,
	}

	for k, v := range keranjang {
		fmt.Println(k, v)
	}

	hasil, iskeluar := keranjang["BA"]

	if iskeluar {
		fmt.Println(hasil)
	} else {
		fmt.Println("item is not exists")
	}

}
