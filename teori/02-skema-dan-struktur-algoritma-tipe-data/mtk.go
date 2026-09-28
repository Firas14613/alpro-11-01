package main

import "fmt"

func main() {
	var a = 10
	var b = 15
	var c = b-a

	fmt.Println(c)

	var i = 10
	i += 5 // i = i + 5
	fmt.Println(i)

	i += 10
	fmt.Println(i)


	j := 20
	j++ // j = j + 1
	fmt.Println(j)
	j-- // j = j - 1
	fmt.Println(j)
}