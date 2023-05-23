package main

import "fmt"

func printList(s map[string]int) {
	if len(s) == 0 {
		fmt.Println("Empty")
		return
	}
	for key, value := range s {
		fmt.Println(value, key)
	}
}
func main() {
	s := make(map[string]int)
	fmt.Println("a -> add	r-> remove 1  	d --> delete item		 p--> print List	q--> Quit\n")
	for {
		fmt.Print(" --> ")
		var op string
		fmt.Scan(&op)
		if op == "q" {
			break
		} else if op == "p" {
			printList(s)
		} else if op == "a" {
			var item string
			fmt.Scan(&item)
			s[item] += 1
		} else if op == "d" {
			var item string
			fmt.Scan(&item)
			_, found := s[item]
			if !found {
				fmt.Println(item, "is not in the list currently, please add using a command")
			} else {
				delete(s, item)
			}
		} else if op == "r" {
			var item string
			fmt.Scan(&item)
			_, found := s[item]
			if !found {
				fmt.Println(item, "is not in the list currently, please add using a command")
			} else {
				s[item] -= 1
				if s[item] <= 0 {
					delete(s, item)
				}
			}

		} else {
			fmt.Println("Uknown command :", op)
		}

	}
}
