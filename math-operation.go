package main

import "fmt"

// func main() { // sementara di disable bentrok main
func main9() {
	// pada dasarnya di go sama dengan bahasa pemrograman lain
	// operator standar yang di gunakan + - * / %
	// example
	var a = 10
	var b = 10
	var c = a + b
	fmt.Println("n1 : ", a)
	fmt.Println("n2 : ", b)
	fmt.Println("hasil : ", c)

	// untuk augmented assignments pada dasarnya sama dengan bahasa pemrograman lain
	// example
	// a = a + 10 -> a += 10
	// a = a - 10 -> a -= 10
	// a = a * 10 -> a *= 10
	// a = a / 10 -> a /= 10
	// a = a % 10 -> a %= 10
	// example penerapan
	c += 10
	fmt.Println("Hasil Setelah Augmented Assignments : ", c)

	// unary merupakan operasi terhadap suatu variabel
	// untuk unary Operator pun pada dasarnya sama dengan bahasa pemrograman lain
	// example
	// ++ Keterangan dari -> a = a + 1
	// -- Keterangan dari -> a = a - 1
	// - Keterangan dari -> Negative
	// + Keterangan dari -> Positive
	// ! Keterangan dari -> Boolean Kebaikan
	a++
	fmt.Println("Hasil Setelah Unary Operato : ", a)

}
