package main

import "fmt"

func main() {
	var (
		suhu float64
		umur int8
	)
	suhu = 36.3
	umur = 10

	fmt.Println("Umur: ", umur)
	fmt.Println("Suhu: ", suhu)
	fmt.Println("Alamat memori suhu: ", &suhu)
	fmt.Println("Alamat memori umur: ", &umur)
}
