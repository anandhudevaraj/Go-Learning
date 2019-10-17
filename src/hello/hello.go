package main

import "fmt"

func main() {
	// Doing all type of var initialisation
	var age int                   // default to zero
	var name = "anandhu"          // if initial value, we can skip data type
	var height, weight = 345, 455 // multiple initialisation same type is supported
	var (
		house = "mepp"
		km    = 3333
	)
	var place, date = "clt", 12344 // multiple is also supported
	first, second := 1, "second"
	fmt.Println(age, name, height, date, weight, house, km, place, first, second)
	fmt.Println("Hello world")
}
