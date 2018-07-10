package main

import "fmt"

func main() {

	var text = "Admin"

	switch  {
	case text == "Admin" || text == "admin":
		fmt.Println("ok")
	default:
		fmt.Println("not ok")

	}

	switch text {
	case "word":
		fmt.Println("not admin")
	case "Admin":
		fmt.Println("Yeah - Admin")
	//fallthrough
	case "User":
		fmt.Println("not admin")
	default:
		fmt.Println("Some text")
	}
}
