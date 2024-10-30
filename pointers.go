package main

import "fmt"

func main() {
	age := 26 //regular varibale

	agePointer := &age
	fmt.Println("Age: ", *agePointer)

	editAgeToAdultYears(agePointer)
	fmt.Println(age)
}

func editAgeToAdultYears(age *int) {
	*age = *age - 18
}
