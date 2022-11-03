package main

import (
	"fmt"
	"strconv"
)

func main49() {
	// package strconv
	/*
		sebelumnya kita sudah belajar cara konversi tipe data misal dari int 32 ke int 34
		bagaimana jika kita butuh melakukan konversi dari tipe data yang berbeda? misalnya int ke string, atau sebaliknya
		hal tersebut kita lakukan dengan menggunakan strconv (string conversion)
	*/
	// beberapa function di package strconv
	/*
		strconv.ParseBool(string) mengubah string ke bool
		strconv.ParseFloat(string) mengubah string ke float
		strconv.ParseInt(string) mengubah string ke int64
		strconv.FormatBool(bool) mengubah bool ke string
		strconv.FormatInt(int,...) mengubah int64 ke string
		strconv.FormatFloat(float32,...) mengubah float64 ke string
	*/

	bolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(bolean)
	} else {
		fmt.Println("Error", err.Error())
	}

	number, err := strconv.ParseInt("28982", 10, 32)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println("Error", err.Error())
	}

	newNumber := strconv.FormatInt(10000, 10)
	if err == nil {
		fmt.Println(newNumber)
	} else {
		fmt.Println("Error", err.Error())
	}

	ck, err := strconv.Atoi("999999")
	fmt.Println(ck)
	fmt.Println(strconv.Itoa(1111))
}
