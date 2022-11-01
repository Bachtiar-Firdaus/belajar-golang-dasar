package main

import "fmt"

// interface kosong mengignore semua tipe data yang di return

func main35() {
	// interface kosong
	/* Golang bukanlah bahasa pemrograman yang berorientasi object
	biasanya dalam pemrograman berorientasi object ada satu data parent di puncak yang bisa
	dianggap sebagai semua implementasi data yang ada di bahsa pemrograman tersebur
	contoh di java ada java.lang.Object
	untuk menangani kasus seperti ini di goalng kita bisa menggunakan interface kosong
	interface kosong adalah interface yang tidak memiliki deklarasi method satupun hal ini
	membuat secara otomatis semua tipe data akan menjadi implementasinya
	*/

	// example penggunaan inteface kosong di golang
	/*
		fmt.Println(a ...interface{})
		panic(v interface{})
		recover() interface{}
		dan lain lain
	*/
	// example 1
	kosong := Ups()
	fmt.Println(kosong)
	// example 2

	newTesting := testingAllReturnNotBlocking(0)
	fmt.Println(newTesting)

}

// example
func Ups() interface{} {
	return 1
	// return true
	// return "Ups"
}

// example 2
func testingAllReturnNotBlocking(number int) interface{} {
	if number < 1 {
		return "Bachtiar nama saya"
	} else if number == 1 {
		return 25
	} else {
		return true
	}
}
