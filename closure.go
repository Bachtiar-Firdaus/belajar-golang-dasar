package main

import "fmt"

func main29() {
	// closure atau berinteraksi dengan data di sebelahnya
	// closure adalah kemampuan sebuah function berinteraksi dengan data data disekitarnya dalam scope yang sama
	// harap gunakan fitu closure ini dengan bijak saat kita membuat aplikasi
	counter := 0
	increment := func() {
		fmt.Println("Increment")
		// jika tidak ingin variabel yang di deklarasikan di luar ber impact disarankan mendeklarasikan variabel ulang
		counter++
	}
	increment()
	increment()
	fmt.Println(counter)
}
