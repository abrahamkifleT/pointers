package main

import "fmt"

func main() {
	age := 26 //regular varibale
	fmt.Println("Age: ", age)

	adultYears := getAdultYears(age)
	fmt.Println(adultYears)
}

func getAdultYears(age int) int {
	return age - 18
}
