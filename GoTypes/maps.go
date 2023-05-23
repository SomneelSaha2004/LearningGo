package main

import "fmt"

func main() {
	//2 ways to make a map
	m := make(map[string]int)
	m["Hello"] = 1
	m["Worlds"] = 2
	fmt.Println(m)
	myMap := map[string]int{
		"nice": 1,
		"wow":  2,
	}
	fmt.Println(myMap)
	// if we acess a key that does not exits default value is assigned there
	//mao remains unchanged
	d := m["defualt"]
	fmt.Println(d)
	// how to delete items
	delete(m, "Hello")
	fmt.Println(m)
	ma := map[string]int(m)
	fmt.Println(ma)
	// to determine if a key exists
	element, found := m["nice"]
	if !found {
		fmt.Println(element, "not found", m)
	}
	m["woog"] = 122
	//iterating through maps
	l := []int{}
	for key, value := range m {
		fmt.Println(key, value)
		l = append(l, value)
	}
	for i, val := range l {
		fmt.Println(i, val)
	}

}
