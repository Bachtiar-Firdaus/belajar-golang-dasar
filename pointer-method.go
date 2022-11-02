package main

import "fmt"

// example tanpa pointer
type Man struct {
	Name string
}

// example pointer
type Ladies struct {
	Name string
}

func main41() {
	// pointer di method atau lebih tepatnya di stuct funtion
	/*
		walaupun method akan menempel di stuct tapi sebenarnya data struct yang diakses di dalam method adalah pass by value
		sangat di rekomendasikan menggunakan pointer di method sehingga tidak boros memory karna harus diduplikasi ketika memanggil method
	*/
	// example tanpa pointer
	daus := Man{"daus"}
	daus.Married()
	fmt.Println(daus)
	// example pointer
	nur := Ladies{"nur"}
	nur.Married()
	fmt.Println(nur)

}

// example tanpa pointer
func (man Man) Married() {
	man.Name = "Mr. " + man.Name
}

// example pointer
func (ladies *Ladies) Married() {
	ladies.Name = "Ms. " + ladies.Name
}
