package main

import "fmt"

type product struct {
	name  string
	price float64
}

func main() {
	s := make([]product, 0)
	fmt.Println("a -> add	d --> delete p--> print List	q--> Quit\n")
	for {
		fmt.Print(" --> ")
		var op string
		fmt.Scan(&op)
		if op == "q" {
			break
		} else if op == "p" {
			fmt.Println(s)
		} else if op == "a" {
			var p product
			fmt.Print("Name : ")
			fmt.Scan(&p.name)
			fmt.Print("Price : ")
			fmt.Scan(&p.price)
			s = append(s, p)
		} else if op == "d" {
			i := 0
			fmt.Scan(&i)
			i--
			next := i + 1
			fmt.Println(i, next)
			if next >= len(s) {
				s = append(s[:i])
				continue
			}
			s = append(s[:i], s[next:]...)
		} else {
			fmt.Println("Uknown command :", op)
		}

	}
}
