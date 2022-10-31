package main

import "fmt"

func main28() {
	// Recursive funtion
	// recursive function adalah function yang memanggil function dirinya sendiri
	// kadang dalam pekerjaan kita sering menemui kasus dimana menggunakan recursive function lebih mudah dibandingkan tidak menggunakan recursive function
	// contoh kasus yang lebih mudah diselesaikan menggunakan recursive adalah factorial
	// example tanpa menggunakan recursive function
	// menggunakan kode program factorial for loop
	fmt.Println(factorialLoop(10))

	// example menggunakan Recursive funtion
	fmt.Println(factorialRecursive(10))
}
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i // result = result * i
	}
	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}
