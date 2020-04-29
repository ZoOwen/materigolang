package main

import "fmt"

func main() {
	status, ok := UserIsYoung(30)
	if !ok {
		fmt.Println(status)
	} else {
		fmt.Println(status)
	}
}

func UserIsYoung(age int) (string, bool) {
	if age < 20 {
		return "User berumur masih muda", true
	}

	return "User sudah tua", false
}
