package main

import "fmt"

func main() {
	var (
		a int
		b int
	)

	fmt.Println(a)
	fmt.Println(b)

	fmt.Scan(&a)
	fmt.Scan(&b)

	a, b = b, a

	fmt.Println(a)
	fmt.Println(b)
}
	
