package main

import "fmt"

func main() {
	var personMobileNo = map[string]string{
		"John":  "+33-8273658526",
		"Steve": "+1-8579822345",
		"David": "+44-9462834443",
	}

	var mobileNo = personMobileNo["Steve"]
	fmt.Println("Steve's Mobile No : ", mobileNo)

	// If a key doesn't exist in the map, we get the zero value of the value type
	mobileNo = personMobileNo["Jack"]
	fmt.Println("Jack's Mobile No : ", mobileNo)

	//key is not exist

	/*var employees = map[int]string{
		1001: "Timmy",
		1002: "Ben",
		1003: "James",
	}
	name, ok := employees[1001]  // "Timmy", true
	name, ok := employees[1010]  // "", false
	_, ok := employees[1005]*/
}
