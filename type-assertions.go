package main

import "fmt"

func main38() {
	// type assertions
	/*
		Type Assertion merupakan kemampuan merubah tipe data menjadi tipe data yang diinginkan
		fitur ini sering sekali digunakan ketika bertemu dengan data interface yang kosong
	*/

	// notes : ketika data tidak sesuai saat di lakukan assertions akan terjadi panic
	result := random()

	// example 1
	// resultString := result.(string) // proses conversi tipe data
	// fmt.Println(resultString)
	// fmt.Println(resultInt)
	// example 2
	// resultInt := result.(int) // proses conversi tipe data
	// fmt.Println(resultInt)

	// type assertions menggunakan switch
	/*
		saat salah menggunakan type assertions maka bisa berakibat terjadi panic di aplikasi
		jika panic dan tidak ter recover maka otomatis program kita akan mati
		agar lebih aman sebaiknya kita menggunakan switch expression untuk melakukan type assertions
	*/
	// example 3
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("int", value)
	case bool:
		fmt.Println("boolean", value)
	default:
		fmt.Println("Unknow")
	}
}

// example
func random() interface{} {
	return 10
}
