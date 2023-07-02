package main

import (
	"fmt"
	"math"
)

type Hitung interface {
	luas() float64
	keliling() float64
}

type Lingkarang struct {
	diameter float64
}

type Persegi struct {
	sisi float64
}

func (p Persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p Persegi) keliling() float64 {
	return p.sisi * 4
}

func (l Lingkarang) jariJari() float64 {
	return l.diameter / 2
}

func (l Lingkarang) luas() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}

func (l Lingkarang) keliling() float64 {
	return math.Pi * l.diameter
}

func main() {

	var bangunDatar Hitung

	bangunDatar = Persegi{10.0}

	fmt.Println("===== persegi")
	fmt.Println("luas      :", bangunDatar.luas())
	fmt.Println("keliling  :", bangunDatar.keliling())

	bangunDatar = Lingkarang{14.0}
	fmt.Println("===== lingkaran")
	fmt.Println("luas      :", bangunDatar.luas())
	fmt.Println("keliling  :", bangunDatar.keliling())
	fmt.Println("jari-jari :", bangunDatar.(Lingkarang).jariJari())

}
