package main

import "fmt"

func main () {
	var n, i int
	var str string

	fmt.Scan(&n)
	fmt.Scan(&str)

	for i = 0; i < n; i++ {
		fmt.Println(str)
	}
}