package main

import (
	"fmt"
	"time"
)

func main54() {
	// package time
	/*
		package time adalah package yang berisikan funsionalitas untuk management waktu di golang
	*/
	// time.Now() untuk mendapatkan waktu saat ini
	// time.Date(...) untuk membuat waktu
	// time.Parse(layout, string) untuk memparsing waktu dari string
	now := time.Now()
	fmt.Println(now.Local())

	utc := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(utc.Local())

	parse, _ := time.Parse(time.RFC3339, "2006-01-03T15:04:05Z")
	fmt.Println(parse)

}
