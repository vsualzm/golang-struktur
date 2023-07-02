package main

import (
	"fmt"
	"regexp"
)

func main() {

	// FindAllString
	var text = "banana burger soup"
	var regex, err = regexp.Compile(`[a-z]+`)

	if err != nil {
		fmt.Println(err.Error())
	}

	var res1 = regex.FindAllString(text, 2)
	fmt.Printf("%#v \n", res1)
	// []string{"banana", "burger"}

	var res2 = regex.FindAllString(text, -1)
	fmt.Printf("%#v \n", res2)
	// []string{"banana", "burger", "soup"}

	// match string
	var dog = "banana burger soup"
	var cok, _ = regexp.Compile(`[a-z]+`)

	var isMatch = cok.MatchString(dog)
	fmt.Println(isMatch) // true

	// FindString
	var cariKata = "banana burger soup"
	var gek, _ = regexp.Compile(`[a-z]+`)

	var str = gek.FindString(cariKata)
	fmt.Println(str) // "banana"

	// findStringIndex
	var kata = "banana burger soup"
	var gekk, _ = regexp.Compile(`[a-z]+`)

	var idx = gekk.FindStringIndex(kata)
	fmt.Println(idx) // [0, 6]

	var stringg = text[0:6]
	fmt.Println(stringg) // "banana"

	//replaceAllString
	var namaTulisan = "banana burger soup"
	var gek1, _ = regexp.Compile(`[a-z]+`)

	var coy = gek1.ReplaceAllString(namaTulisan, "potato")
	fmt.Println(coy)
	// "potato potato potato"

	// replaceAllStringFunc
	var namanama = "banana burger soup"
	var koss, _ = regexp.Compile(`[a-z]+`)

	var data1 = koss.ReplaceAllStringFunc(namanama, func(each string) string {
		if each == "burger" {
			return "potato"
		}
		return each
	})
	fmt.Println(data1)
	// "banana potato soup"

	// dengan metode split
	var textSplit = "banana burger soup"
	var geks2, _ = regexp.Compile(`[a-b]+`) // split dengan separator adalah karakter "a" dan/atau "b"

	var con = geks2.Split(textSplit, -1)
	fmt.Printf("%#v \n", con)
	// []string{"", "n", "n", " ", "urger soup"}

}
