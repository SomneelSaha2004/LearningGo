//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

type item struct {
	name string
	tag  bool
}

func deactivate(i *item) {
	i.tag = false
}
func activate(i *item) {
	i.tag = true
}
func checkout(arr []item) {
	for i, _ := range arr {
		deactivate(&arr[i])
	}
}
func printItems(arr []item) {
	var status string
	for _, i := range arr {
		if i.tag {
			status = "Active"
		} else {
			status = "Inactive"
		}
		fmt.Println(i.name, status)
	}
	fmt.Println()
}
func main() {
	//  - Create at least 4 items, all with active security tags
	//  - Store them in a slice or array
	//  - Deactivate any one security tag in the array/slice
	//  - Call the checkout() function to deactivate all tags
	//  - Print out the array/slice after each change
	items := []item{{"s", true}, {"a", true}, {"b", true}, {"c", true}}
	deactivate(&items[1])
	printItems(items)
	checkout(items)
	printItems(items)
}
