package main

import "fmt"

func main() {
	var bilangan, bil1, bil2 int
	var ascending = true
	var descending = true

	fmt.Scan(&bilangan)

	if bilangan < 0 {
		bilangan = -bilangan
    }

	bil1 = bilangan % 10
	bilangan /= 10

	for bilangan > 0 {
		bil2 = bilangan % 10
		if bil2 < bil1 {
			descending = false
		} else if bil2 > bil1 {
			ascending = false
		}
		bil1 = bil2
		bilangan /= 10
	}

	if ascending && !descending {
		fmt.Println("ascending")
	} else if !ascending && descending {
		fmt.Println("descending")
	} else {
		fmt.Println("tidak terurut")
	}
}