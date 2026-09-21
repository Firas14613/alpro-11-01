package main

import "fmt"

func main() {
	var umur int8
	var anake int8
	var saldo int64

	umur = 18
	anake = 2
	saldo = 1000000000

	fmt.Println("Umur saya: ", &umur)
	fmt.Println("Saya Anak ke: ", anake)
	fmt.Println("Saya Memiliki Saldo: ", saldo)

}
