package main

import "fmt"

func main12() {
	// pada dasarnya array di golang sama dengan bahasa pemrograman lain
	// sedikit perbedaan pada tipe data array di golang perlu mendeklarasikan jumlah kemampuan variabel array untuk menampung datanya
	// example ketika di deklarasikan kemampuannya 10 maka data yang sanggup di tampung hanya 10 dst
	// untuk indexing array di go sama dengan bahasa pemrograman lain dibulai dari 0

	// example case
	var location [3]string
	location[0] = "palas"
	location[1] = "kalianda"
	location[2] = "lampung-selatan"

	fmt.Println(location)
	// example pemanggilan pada index array
	fmt.Println(location[2])

	// example untuk parsing dari dalam index keluar ke varibel baru
	locationDistric := location[0]
	fmt.Println(locationDistric)

	// example untuk mempersingkat dalam pembuatan
	var exampleNumber = [3]int{
		90,
		80,
	}
	fmt.Println(exampleNumber)

	// example function array
	// len(array) Untuk mendapatkan panjang Array
	// array[index] Mendapatkan data index
	// array[index] Mengubah data di posisi index
	// example case
	fmt.Println(len(location))
	fmt.Println(len(exampleNumber))

	// unix case ketika variabel array tersebut belum terisi values maka ketika di lens akan tetap sama dalam pengecekan jumlah panjang array dari variabel tersebut
	// example
	var jumahBaju [20]int
	fmt.Println("meski belum terisi data untuk variabel julah baju tetap terdefine 20 karna sesuai dengan deklarasi")
	fmt.Println("Resul var : ", jumahBaju)
	fmt.Println("Resul var length : ", len(jumahBaju))

}
