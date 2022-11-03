package main

import (
	"fmt"
	"math"
)

func main50() {
	// package math
	/*
		package math merupakan package yang berisikan constant dan fungsi matematika
	*/
	// math.Round(float64) membulatkan float64 ke atas atau kebawah sesuai dengan yang paling dekat
	// math.Floor(float64) membulatkan float64 kebawah
	// math.Ceil(float64) membulatkan float64 ke atas
	// math.Max(float64) mengembalikan nilai64 paling besar
	// math.Min(float64) mengembalikan nilai float64 paling kecil

	fmt.Println(math.Round(1.7))
	fmt.Println(math.Round(1.3))

	fmt.Println(math.Floor(1.7))
	fmt.Println(math.Floor(1.3))

	fmt.Println(math.Ceil(1.7))
	fmt.Println(math.Ceil(1.3))

	fmt.Println(math.Max(10, 20))
	fmt.Println(math.Max(10, 20))

	fmt.Println(math.Min(10, 20))
	fmt.Println(math.Min(10, 20))
}
