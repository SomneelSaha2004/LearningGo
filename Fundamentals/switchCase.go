package main

import "fmt"

func main() {
	var i int
	fmt.Scan(&i)
	switch result := i % 7; {
	case result > 4:
		fmt.Println("Greater than 4")
	case result == 1:
		fmt.Println(1)
		fallthrough
	case result < 4 && result > 1:
		fmt.Println("Nice")
	}
	var result rune = 'a'

	switch result {
	case 'a', 'e', 'i', 'o', 'u':
		fmt.Println("A vowel")
		fallthrough
	case 'A', 'E', 'I', 'O', 'U':
		fmt.Println("Vowels are cool")
	default:
		fmt.Println("Consonant")
	}

}
