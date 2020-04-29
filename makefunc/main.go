package main

import "fmt"

func main() {
	fmt.Println("Hallo")
	fmt.Println("Hallo", sayPagi())
}

func sayPagi() map[string]int {
	var m = make(map[string]int)
	m["1001"] = 1001
	m["1002"] = 1002

	return m
}
