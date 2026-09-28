package main

import "fmt"

func main() {
	type angkaBang = int

	var nilai1 angkaBang = 100
	var nilai2 angkaBang = 200

	var contoh int = 300
	var contohAngkaBang angkaBang = angkaBang(contoh)

	fmt.Println(nilai1)
	fmt.Println(nilai2)
	fmt.Println(contohAngkaBang)
}