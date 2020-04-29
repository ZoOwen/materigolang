package main

import "fmt"

func main() {
	// Initializing a map
	var tinderMatch = make(map[string]string)

	// Adding keys to a map
	tinderMatch["Tommy"] = "Angelina" // Assigns the value "Angelina" to the key "Tommy"
	tinderMatch["James"] = "Sophia"
	tinderMatch["David"] = "Emma"

	fmt.Println(tinderMatch)

	/*
	   Adding a key that already exists will simply override
	   the existing key with the new value
	*/
	tinderMatch["Tommy"] = "Jennifer"
	fmt.Println(tinderMatch)
}
