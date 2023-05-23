package main

import "fmt"

type Room struct {
	number  int
	cleaned bool
}

func main() {
	// var arr [4]int
	// fmt.Println(arr)
	// //arrays can be printed
	// b := [3]int{1, 2}
	// fmt.Println(b)
	// c := [...]string{"Hello"}
	// fmt.Println(len(c), b[1])
	// b[2] = 3
	// fmt.Println(b[:2])
	rooms := [3]Room{{1, true}, {3, false}, {45, false}}
	//fmt.Println(rooms)
	var key int
	fmt.Scan(&key)
	for i := 0; i < len(rooms); i++ {
		//fmt.Println(rooms[i].number)
		if rooms[i].number == key {
			if rooms[i].cleaned {
				fmt.Println(rooms[i].number, "is Cleaned")
				break
			}
			fmt.Println("Not cleaned")
			break
		}
	}
}
