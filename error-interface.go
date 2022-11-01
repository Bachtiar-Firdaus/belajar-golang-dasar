package main

// kita tidak perlu membuat type dibawah ini karna by default sudah  dibuatkan oleh golang
// sebagai default error
// type error interface {
// 	Error() string
// }

// example untuk memanggil package yang defaul dibuatkan oleh golang
import (
	"errors"
	"fmt"
)

func main37() {
	// error interface
	// golang memiliki interface yang digunakan sebagai kontrak untuk membuat error nama interfacenya adalah eror

	// Membuat Error
	/*
		Untuk membuat error kita tidak perlu manual
		golang sudah menyediakan library untuk membuat helper secara mudah yang terdapat di package errors (Package akan kita bahsa secara mendetail di materi tersendiri)
	*/
	// example 1
	var contohError = errors.New("sample errors")
	fmt.Println(contohError)

	// example 2
	hasil, err := Pembagian(10, 10)
	fmt.Println(hasil, err)
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Hasil", err.Error())

	}
}

// example penggunaan
func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi Dengan NOL dilarang")
	} else {
		return nilai / pembagi, nil
	}
}
