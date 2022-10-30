package main

import "fmt"

func main24() {
	// Variadic Function
	// Parameter yang berbeda di posisi terakhir memiliki kemampuan sebuah varargs
	// Varargs artinya datanya bisa menerima lebih dari stu input atau anggap saja semacam array
	// apa bedanya dengan parameter biasa dengan tipe array?
	// - Jika parameter tipe array , kita wajib membuat array terlabih dahulu sebelum mengirimkan ke function
	// - Jika parameter menggunakan varargs kita bisa langsung mengirim datanya jika lebih dari satu cukup menggunakan tandakoma
	// contoh vaiadic function
	// func sumAll(number ...int) int{
	// yang menjadi kan indikasi adanya variadic function adalah ...
	// example 1
	total := sumAll(10, 10, 10, 10, 10, 10)
	fmt.Println(total)
	// example 2
	keterangan, total := newSumAll("", 10, 10, 10, 10, 10, 10)
	fmt.Println(keterangan, total)

	// Slice Parameter
	// kadang ada kasus dimana kita menggunakan variadic function namun memiliki variabel berupa slice
	// kita bisa menjadikan slice sebagai varargs parameter
	// example 3
	slice := []int{10, 10, 10, 10, 10, 2}
	// cara parsing data slice agar bisa digunakan di varargs menggunakan ...
	newTotal := sumAll(slice...)
	fmt.Println(newTotal)

}

// example 1
func sumAll(number ...int) int {
	total := 0
	for _, number := range number {
		total += number
	}
	return total
}

// example 2
func newSumAll(hasil string, newNumber ...int) (string, int) {
	total := 0
	keterangan := "Hasil dari variadic"
	for _, newNumber := range newNumber {
		total += newNumber
	}
	return keterangan, total
}
