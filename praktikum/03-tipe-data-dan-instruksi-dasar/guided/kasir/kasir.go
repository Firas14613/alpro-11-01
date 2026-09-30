package main

import "fmt"

func main() {
	var x int
	
	fmt.Print("Masukan Uang: ")
	fmt.Scan(&x)

	lembarSepuluh := x / 10000
	sisaSepuluh := x % 10000
	lembarLima := sisaSepuluh / 5000
	sisaLima := sisaSepuluh % 5000
	lembarSeribu := sisaLima / 1000

	fmt.Println(lembarSepuluh, " ", lembarLima, " ", lembarSeribu)
}
