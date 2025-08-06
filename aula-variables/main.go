package main

import "fmt"

func main() {

	var name, carPlate, isDriver = "Rafael", "ABD1234", true

	fmt.Println(name, carPlate, isDriver)

	fmt.Printf("Name: %s, Car Plate: %s, Is Driver: %t\n", name, carPlate, isDriver)

	var newVariable = "New Variable"

	fmt.Printf("New Variable: %s\n", newVariable)
}
