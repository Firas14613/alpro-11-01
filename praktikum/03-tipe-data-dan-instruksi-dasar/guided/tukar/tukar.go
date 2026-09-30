package main

import "fmt"

func main() {
	var y, x, z int

	fmt.Print("Masukan nilai x: ")
	fmt.Scan(&x)
	fmt.Print("Masukan nilai y: ")
	fmt.Scan(&y)
	fmt.Print("Masukan nilai z: ")
	fmt.Scan(&z)

	// x, y, z = z, x, y
	temp := x
	x = y
	y = z
	z = temp

	fmt.Println(x," ",y," ",z)
}
