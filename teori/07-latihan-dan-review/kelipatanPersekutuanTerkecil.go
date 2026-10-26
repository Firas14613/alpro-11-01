package main

import "fmt"

func main() {
	var n, m int

	fmt.Scan(&n, &m)

	kpk := n
	if m > n {
		kpk = m
	}

	for {
		if kpk % n == 0 && kpk % m == 0 {
			break
		} else {
			kpk++
		}
	}

	fmt.Println(kpk)
}