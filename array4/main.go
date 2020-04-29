package main

import "fmt"

func main() {
	var a = [2]int{1, 2}

	//var b = [...]int{1, 2}

	var c = [2]int{1, 3}

	var opperand = a == c

	fmt.Println(opperand)

}
