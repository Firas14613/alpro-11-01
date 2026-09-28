package main

import "fmt"

func main() {
	var uang int
	fmt.Scanln(&uang)

	fmt.Println(uang / 10000, " Lembar")
	fmt.Println((uang % 10000) / 5000, " Lembar")
	fmt.Println((uang % 5000) / 1000, " Lembar")
}
