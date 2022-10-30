package main

import "fmt"

// func main() { // sementara di disable bentrok main
func main7() {
	// example convertion tipe data
	// notes : wajib di perhatikan isi varibel existing apa bila isi variabel melebihi variabel tujuan conversi maka value di turunkan paksa
	var nilai32 int32 = 1234
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai16, nilai32, nilai64)
	// notes : wajib di perhatikan isi varibel existing apa bila isi variabel melebihi variabel tujuan conversi maka value di turunkan paksa
	// example variabel over flow atau isi variabel melebihi batas

	var nilai8 int8 = int8(nilai32)
	fmt.Println(nilai8)
	// mengapa hasilnya bisa -46
	// karna ketika isi variabel melebihi batas akan kembali ke nilai awal
	// contoh int 8 yang range -128 sd 127 ketika nilai 128 makan nilai existing akan kembali ke -128
	// example
	var example int16 = 128
	var anomali int8 = int8(example)
	fmt.Println("sebelum di paksa conversi", example, "setelah di paksa conversi", anomali)
	example = 129
	anomali = int8(example)
	fmt.Println("sebelum di paksa conversi", example, "setelah di paksa conversi", anomali)
	example = 130
	anomali = int8(example)
	fmt.Println("sebelum di paksa conversi", example, "setelah di paksa conversi", anomali)

	// dimateri conversi string index sebelumnya sudah berhasil mendapatkan nilai byte dari index string
	// example merubah byte menjadi real carakter
	var testingString = "Bachtiar"
	var realKarakterDalamByte = testingString[0]
	var resultByteToRealKarakter = string(realKarakterDalamByte)
	fmt.Println("Berikut Byte Index Dari String", realKarakterDalamByte)
	fmt.Println("Berikut Hasil Konversi Dari Byte Index Menjadi Real Karakter", resultByteToRealKarakter)
}
