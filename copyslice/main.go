package main

import "fmt"

func main() {
	src := []string{"Sublime", "VSCode", "IntelliJ", "Eclipse"}
	dest := make([]string, 2)
	//arr1 := []int{1, 2, 3, 4, 5}

	//rr2 := []int{6, 7, 8, 9, 10}

	numElementsCopied := copy(dest, src)

	//concat := (arr1,arr2)

	fmt.Println("src = ", src)
	fmt.Println("dest = ", dest)
	fmt.Println("Number of elements copied from src to dest = ", numElementsCopied)
	fmt.Println("hasil 2 array jadi 1")
}
