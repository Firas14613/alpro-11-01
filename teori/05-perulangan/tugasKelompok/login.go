package main

import "fmt"

func main() {
	var pass, user string

	fmt.Scan(&user, &pass)

	for user != "admin" || pass != "admin" {
		fmt.Scan(&user, &pass)
	}

	fmt.Println("Login Berhasil")
}

// program login
// kamus
// 	pass, user : string
// algoritma
// 	input(user, pass)

// 	while user != "admin" || pass != "admin" do
// 		input(user, pass)
// 	endwhile

// 	output("Login Berhasil")
// endprogram