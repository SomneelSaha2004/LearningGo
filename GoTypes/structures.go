package main

import "fmt"

type Point struct {
	x, y float64
}
type Person struct {
	name, address string
	age           int
	position      Point
}

func main() {
	p := Point{2.3, 3.4}
	fmt.Println(p)
	sam := Person{
		address:  "somewhere",
		name:     "sam",
		age:      19,
		position: Point{1, 2},
	}

	fmt.Printf("sam: %v\n", sam)
	sam.age++
	sam.name = "Rick"
	fmt.Printf("sam: %v\n", sam)
	//if we pass no paramters
	// fields will have defualts
	mary := Person{name: "mary"}
	fmt.Printf("mary: %v\n", mary)
	// we can also create anonymous structures using var
	var dog struct {
		name, breed string
		age         int
	}
	dog.name = "Bruno"
	fmt.Println(dog)
	// we can do this in one line also
	cat := struct {
		name, breed string
		age         int
	}{
		"cecil", "smol",
		5,
	}
	// other fields will be defualt
	fmt.Println(cat)
}
