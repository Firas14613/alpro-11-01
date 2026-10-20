package main

import "fmt"

func main() {
	var n, mobil int

	fmt.Scanln(&n)

	mobil = n/7
	if n % 7 != 0 {
		mobil++
	}

	fmt.Println(mobil)
}