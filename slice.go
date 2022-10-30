package main

import "fmt"

func main13() {
	// untuk tipe data slice jarang di temukan di bahasa pemrogrman lain
	// slice merupakan ptongan dari data array
	// slice dan array selalu terkoneksi dimana slice adalah data yang mengakses sebagian atau seluruh data dari array
	// tipe data slice meiliki 3 data yaitu pointer, length dan capasity

	// - pointer adalah penunjuk data pertama di array para slice
	// - length alah panjang dari slice dan
	// - capacity adalah kapasitas dari slice, dimana tidak boleh lebih dari capacity

	// example
	// slice membutuhkan existing array
	// array[low:high] membuat slice dari array dimulai index low sampai index sebelum high
	// exampl array[1:10] maka array index rang yang consume dari 1-9 karna di posisi high mengambil nilai high -1
	// array[low:] membuat slice dari array dimulai index low sampai index akhir di array
	// array[:high] membuat slice dari array dimulai dari index 0 sampai 1 index sebelum high
	// array[:] membuat slice dari array di muluai index 0 sampai index akhir di array

	// warning ketika merubah data di array semua data yang ada di slice akan ikut berubah
	//example case

	// [...] adalah auto generate max index
	name := [...]string{
		"daus",
		"tyo",
		"ragil",
		"bachtiar",
		"miqdad",
		"elfin",
		"ega",
	}
	nameSlice := name[1:3]
	// maka yang akan ditampilkan adalah tyo dan ragil karna data yang di tampilkan index ke 1 dan 2

	fmt.Println(nameSlice)
	// exapmle penerapan semua case
	fmt.Println(name[3:])
	fmt.Println(name[:4])
	fmt.Println(name[1:3])
	fmt.Println(name[:])

	// example case pada function slice
	// len(slice) untuk mendapatkan panjang
	// cap(slice Untuk mendapatkan kapasitas)
	// append(slice, data) Membuat slice baru dengan menambahkan data ke posisi terakhir slice, jika kapasitas sudah penuh, maka akan membuat array baru
	// make([]TypeData, length, capacity) Membuat slice baru
	// copy(destination, source) Menyalin slice dari source ke destination
	// example
	fmt.Println("data yang di slice", name[2:3])
	fmt.Println("panjang data yang di slice", len(name[2:3]))
	fmt.Println("kapasitas maksimal slice : ", cap(name[2:3]))

	// warning untuk tidak asal merubah slice karna merefer langsung ke data induk
	fmt.Println("=======================")
	fmt.Println("master data", name)
	exampleSlice := name[2:3]
	exampleSlice[0] = "Head-Phone"
	fmt.Println("Exaple Slice change data menjadi : Head-Phone", exampleSlice)
	fmt.Println("master data after change in slice", name)

	// maka ketika master data di panggil lagi datanya telah berubah

	// example pada case hari
	fmt.Println("=======================")
	days := [...]string{
		"senin,",
		"selasa,",
		"rabu,",
		"kamis,",
		"jumat,",
		"sabtu,",
		"minggu,",
	}
	daySlice1 := days[5:]
	daySlice1[0] = "Sabtu Baru,"
	daySlice1[1] = "Minggu Baru,"
	fmt.Println("Master data", days)

	// ketika dilakukan append maka tidak merubah data master
	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Ups,"

	fmt.Println("Data Menggunakan Append", daySlice2)
	fmt.Println("Master data", days)
	// notes : jika menggunakan append tidak merubah data master jika index terdeteksi full

	fmt.Println("=======================")
	tes := [...]int8{
		0,
		1,
		2,
	}
	fmt.Println(tes)
	dataSlice := tes[1:3]
	dataAppend := append(dataSlice, 10)
	dataAppend[0] = 1
	fmt.Println(dataAppend)
	fmt.Println(tes)

	fmt.Println("=======================")

	// example case make slice
	// ini case paling aman ketika penggunaan slice
	// karna array berada didalam slice dan tidak terlihat dalam sebuah variabel
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Daus"
	newSlice[1] = "Daus"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// example case copy slice
	fromSlice := newSlice[:]
	toSlice := make([]string, len(newSlice), cap(newSlice))
	copy(toSlice, fromSlice)
	fmt.Println(toSlice)

	// warning ketika declarasi jangan sampai salah mendeklarasikan mana new array mana slice
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
