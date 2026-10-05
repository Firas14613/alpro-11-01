package main

import "fmt"

func main() {
	var a, b int
	var hasilLebih, sama, hasilKurang bool

	fmt.Scan(&a, &b)
	hasilLebih = a > b
	sama = a == b
	hasilKurang = a < b

	fmt.Println(hasilLebih, " ", sama, " ", hasilKurang)
}