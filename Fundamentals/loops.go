package main

import "fmt"

func main() {
	//Go only has for loops
	var n int
	fmt.Scan(&n)
	//basic
	for i := 1; i <= 10; i++ {
		fmt.Println(i * n)
	}
	//while loop
	i := 2
	for i < 10 {
		if i%2 == 0 {
			fmt.Println(i)
		}
		i++
	}
	//infinite loop
	for {
		break
	}
}
