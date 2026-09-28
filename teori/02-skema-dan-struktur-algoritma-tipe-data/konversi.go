package main

import "fmt"

func main() {
	// var nilai32 int32 = 3278
	// var nilai64 int64 = int64(nilai32)
	// var nilai8 int8 = int8(nilai32)

	// fmt.Println(nilai32)
	// fmt.Println(nilai64)
	// fmt.Println(nilai8)
	//pada int8, karena data yang di konversi lebih besar dari kapasitas int8, maka nilainya akan berubah menjadi -2. karena kapasitas int8 hanya sampai 127, jadi jika melebihi kapasitasnya maka nilainya akan berubah menjadi negatif

	var nama string = "Firas Abdurrahman Sandro"
	var e = nama[0]
	var eString = string(e)

	fmt.Println(nama)
	fmt.Println(e)
	fmt.Println(eString)
	//pada konversi string ke byte, maka akan mengembalikan nilai byte dari karakter pertama pada string. jadi jika ingin mengubahnya menjadi string lagi, maka harus di konversi lagi menjadi string
}