package main

import "fmt"

func main() {
	var ch rune = '🐲' //rune is just an alias for int32
	var s string = "Hello World\n"
	fmt.Println(string(ch) + s) //concatenation is possible
	//to create a string with quotes use raw literals
	s = `he said "BRUH \n" `
	fmt.Println(s)
	fmt.Println(ch, s, "Hello")
	name := "Somneel"
	fmt.Println("Hello my name is ", name, "How are you")
}
