package main

import "fmt"

func main() {
	var celcius float64

	fmt.Print("Masukan suhu: ")
	fmt.Scan(&celcius)

	r := 4.0 / 5.0 * celcius
	fmt.Println("Reamur: ", r)
}
