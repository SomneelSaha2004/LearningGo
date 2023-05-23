//Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import "fmt"

type Vehicle interface {
	goToLift()
	String() string
}
type Motorcycle struct {
	name      string
	modelType string
}
type Car struct {
	name      string
	modelType string
}
type Truck struct {
	name      string
	modelType string
}

func (m Motorcycle) goToLift() { fmt.Println("Small lift") }
func (m Motorcycle) String() string {
	return fmt.Sprintf("Name : %v\nModel Type : %v", m.name, m.modelType)
}
func (m Car) goToLift() { fmt.Println("Standard lift") }
func (m Car) String() string {
	return fmt.Sprintf("Name : %v\nModel Type : %v", m.name, m.modelType)
}
func (m Truck) goToLift()      { fmt.Println("Large lift") }
func (m Truck) String() string { return fmt.Sprintf("Name : %v\nModel Type : %v", m.name, m.modelType) }
func whichLift(v Vehicle) {
	v.goToLift()
	fmt.Println(v)
}
func main() {
	v := []Vehicle{}
	var (
		ch        string
		name      string
		modelType string
		flag      bool = false
	)
	for {

		fmt.Print("Enter type of vehicle : ")
		fmt.Scan(&ch)
		fmt.Println(ch)
		switch ch {
		case "c":
			fmt.Println("Enter Car details")
			fmt.Print("name : ")
			fmt.Scan(&name)
			fmt.Print("Model Type : ")
			fmt.Scan(&modelType)
			v = append(v, Car{name, modelType})
		case "t":
			fmt.Println("Enter Truck details")
			fmt.Print("name : ")
			fmt.Scan(&name)
			fmt.Print("Model Type : ")
			fmt.Scan(&modelType)
			v = append(v, Truck{name, modelType})
		case "m":
			fmt.Println("Enter Motorcycle details")
			fmt.Print("name : ")
			fmt.Scan(&name)
			fmt.Print("Model Type : ")
			fmt.Scan(&modelType)
			v = append(v, Motorcycle{name, modelType})
		case "q":
			flag = true
		}
		if flag {
			break
		}
	}
}
