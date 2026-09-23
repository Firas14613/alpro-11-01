package main

import "fmt"

func main() {
	var (
		a int
		b int
	)

	fmt.Scanln(&a)
	fmt.Scanln(&b)

	tambah := a+b
	kurang := a-b
	kali := a*b
	bagi := a/b
	sisa := a%b

	fmt.Print(tambah, " ", kurang, " ", kali, " ", bagi, " ", sisa)

}
