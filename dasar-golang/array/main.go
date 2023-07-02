package main

import "fmt"

func main() {

	// array

	var names [3]string
	names[0] = "ilham"
	names[1] = "ivanka"
	names[2] = "rauzi"

	fmt.Println(names) // [ilham ivanka rauzi]

	var data [3]int

	data[0] = 1
	data[1] = 2
	data[2] = 3

	fmt.Println(data) // [1 2 3]

	// Berapa isi dalam array menggunakan len()
	fmt.Println("isi dalam array:", len(data))

	// array dengan variabel ke bawah
	// jika array Ada angka di dalam kurung nya
	// slice tidak ada di dalam kruung nya
	kota := [3]string{
		"Bandung",
		"jakarta",
		"Surabaya",
	}

	fmt.Println(kota)

	// array Multi dimensi

	// var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}
	var numbers3 = [2][4]int{{2, 1, 2, 3}, {3, 4, 5, 6}}

	// fmt.Println("numbers1", numbers1)
	fmt.Println("numbers3", numbers3)
	fmt.Println("numbers2", numbers2)

	// array with for
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	for i := 0; i < len(fruits); i++ {
		fmt.Printf("elemen %d : %s\n", i, fruits[i])
	}

	barang := [3]string{"hp", "laptop", "mouse"}
	for i := 0; i < len(barang); i++ {
		fmt.Printf("barang %d : %s\n", i, fruits[i])
	}

	// 	// memakai for range
	indera := [3]string{"melihat", "merasa", "mendengar"}
	for i, indra := range indera {
		fmt.Printf("indera %d : %s\n", i, indra)
	}

	// menggunakan make()
	var dor = make([]string, 2)
	dor[0] = "apple"
	dor[1] = "manggo"

	fmt.Println(dor) // [apple manggo]

	fmt.Println("================================================")

	// menggunakan array slice

	nama := []string{
		"abah juang",
		"Rene Albert",
		"Luis Milla",
		"aldof ridel",
	}

	for _, namaPelatih := range nama {
		fmt.Printf("nama Pelatih: %s\n", namaPelatih)
	}

	// menggunakan slice

	cokes := []string{"ilham", "alam", "aslan", "nafi"}

	fmt.Println("isi names \t:", len(cokes))
	fmt.Println(cokes)

	crot := append(cokes, "rudi")
	fmt.Println(crot)

	dst := make([]string, 3)
	src := []string{"watermelon", "pinnaple", "apple", "orange"}
	n := copy(dst, src)

	fmt.Println(dst) // watermelon pinnaple apple
	fmt.Println(src) // watermelon pinnaple apple orange
	fmt.Println(n)   // 3

	cuk := []string{"potato", "potato", "potato"}
	cak := []string{"watermelon", "pinnaple"}
	na := copy(cuk, cak)

	fmt.Println(cuk) // watermelon pinnaple potato
	fmt.Println(cak) // watermelon pinnaple
	fmt.Println(na)  // 2

}
