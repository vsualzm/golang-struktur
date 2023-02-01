package main

import "fmt"

// day 3
func main() {

	jumlah := 8 + 3
	fmt.Println(jumlah)

	kurang := 10 - 3
	fmt.Println(kurang)

	kali := 10 * 6
	fmt.Println(kali)

	bagi := 20 / 2
	fmt.Println(bagi)

	modulus := 8 % 3
	fmt.Println(modulus)

	// perbandingan

	angka := 8
	fmt.Println(angka > 5) // true

	fmt.Println(angka < 5) //false

	fmt.Println(angka >= 5) // false

	fmt.Println(angka <= 5) // false

	fmt.Println(angka == 5) // false

	fmt.Println(angka != 5) // true

	// operator logic

	a := true
	b := false
	c := true
	d := false

	fmt.Println(a && c)

	fmt.Println(a && b)

	fmt.Println(a || b)

	fmt.Println(b || d)

	fmt.Println(!b && !d)

	fmt.Println(!a || b)

	// kondisional if else

	tokoKelontong := "buka"

	if tokoKelontong == "buka" {
		fmt.Println("toko sudah buka")
	} else if tokoKelontong == "tutup" {
		fmt.Println("toko tutup")
	} else {
		fmt.Println("teuing tkoo kamana ajig")
	}

	var minimarketStatus = "open"
	var telur = "soldout"
	var buah = "soldout"
	if minimarketStatus == "open" {
		fmt.Println("saya akan membeli telur dan buah")
		if telur == "soldout" || buah == "soldout" {
			fmt.Println("belanjaan saya tidak lengkap")
		} else if telur == "soldout" {
			fmt.Println("telur habis")
		} else if buah == "soldout" {
			fmt.Println("buah habis")
		}
	} else {
		fmt.Println("minimarket tutup, saya pulang lagi")
	}

	var buttonPushed = 1
	switch {
	case buttonPushed == 1:
		fmt.Println("matikan TV!")
	case buttonPushed == 2 && buttonPushed == 3:
		fmt.Println("turunkan volume TV!")
	case buttonPushed == 4:
		fmt.Println("matikan suara TV!")
	default:
		fmt.Println("Tidak terjadi apa-apa")
	}

	var point = 6

	switch {
	case point == 8:
		fmt.Println("perfect")
	case (point < 8) && (point > 3):
		fmt.Println("awesome")
		fallthrough
	case point < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

}
