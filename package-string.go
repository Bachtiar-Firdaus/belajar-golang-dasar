package main

import (
	"fmt"
	"strings"
)

func main48() {
	// package string
	/*
		package string adalah package yang berisikan function function untuk memanipulasi tipedata string
		ada banyak sekali function yang kita bisa gunakan
		strings.Trim(string, cutset) 			-> memotong cutset diawal dan akhir string
		strings.ToLower(string)					-> membuat semua karakter menjadi lowercase
		strings.ToUpper(string)					-> membuat semua karakter menjadi uppercase
		strings.Split(string,separator)			-> memotong string berdasarkan separator
		strings.Contains(string, search)		-> mengecek apakah string mengandung string lain
		strings.ReplaceAll(string, from, to)	-> mengubah semua string dari from ke to
	*/
	fmt.Println(strings.Contains("bachtiar firdaus", "daus"))
	fmt.Println(strings.Split("bachtiar firdaus", " "))
	fmt.Println(strings.ToLower("BACHTIAR"))
	fmt.Println(strings.ToUpper("FIRDAUS"))
	fmt.Println(strings.Trim("       Bachtiar             ", " "))
	fmt.Println(strings.ReplaceAll("daus daus daus daus", "daus", "tampan"))

}
