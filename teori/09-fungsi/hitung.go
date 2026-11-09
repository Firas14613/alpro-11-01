package main

import "fmt"

func hitung(x, y int) int {
	return 2 * x + y - 5
}

func main() {
	hasil := hitung(10, 3)
	fmt.Println(hasil)
	fmt.Println(hasil * 2)
	fmt.Println(hitung(2, 1))
}