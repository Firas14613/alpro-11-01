package main

import "fmt"


func main() {
	const phi float64 = 3.14
	var jari float64
	

	fmt.Scan(&jari)

	luas := phi * jari * jari
	fmt.Println(luas)
}
