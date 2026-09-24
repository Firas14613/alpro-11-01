package main

import "fmt"

func main() {
	var (
		a int
		b int
	)

	for {
		fmt.Scanln(&a)
		fmt.Scanln(&b)

		if b == 0 {
			fmt.Println("Input Ulang")
			continue
		}

		fmt.Println(a+b, " ", a-b, " ", a*b, " ", a/b, " ", a%b)
		break
	}
	// tambah := a+b
	// kurang := a-b
	// kali := a*b
	// bagi := a/b
	// sisa := a%b

	// if b==0 {
	// 	fmt.Println("Input Ulang")
	// }

	// fmt.Println(a+b)
	// fmt.Println(a-b)
	// fmt.Println(a*b)
	// fmt.Println(a/b)
	// fmt.Println(a%b)
	// fmt.Println(a+b, " ", a-b, " ", a*b, " ", a/b, " ", a%b)

}
