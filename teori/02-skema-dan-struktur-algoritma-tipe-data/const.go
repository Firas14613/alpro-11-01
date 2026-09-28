package main

import "fmt"

func main() {
	// const firstName string = "Firas"
	// const lastName = "Abdurrahman Sandro"
	// fmt.Println(firstName, lastName)
	// const tidak bisa di ubah nilainya, jadi jika sudah di deklarasikan maka nilainya tidak bisa di ubah lagi. jadi const ini bersifat immutable
	// beda dengan variable, const nilainya tidak bisa di ubah, dan const tidak akan error jika tidak di gunakan

	//akan error jika di ubah nilainya
	// firstName = "Firas Abdurrahman Sandro"

	const (
		firstName = "Firas"
		lastName  = "Abdurrahman Sandro"
	)
	fmt.Println(firstName, lastName)
	// const bisa di tulis langsung banyak seperti ini. jadi tidak perlu mengerik satu persatu
}