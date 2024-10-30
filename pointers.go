package main

import "fmt"

func main() {
	age := 26 //regular varibale

	agePointer := &age
	fmt.Println("Age: ", *agePointer)

	adultYears := getAdultYears(agePointer)
	fmt.Println(adultYears)
}

func getAdultYears(age *int) int {
	return *age - 18
}
