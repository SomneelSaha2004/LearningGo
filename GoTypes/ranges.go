package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	chars := strings.Split(s, "")
	rev := ""
	for i := len(chars) - 1; i >= 0; i-- {
		rev += chars[i]
	}
	fmt.Println(rev)
	return rev == s
}
func printPalindromes(s string) {
	words := strings.Split(s, " ")
	for _, word := range words {
		fmt.Println(word)
		if isPalindrome(word) {
			fmt.Println(word)
		}
	}
}
func main() {
	s := []int{1, 2, 2, 21}
	for i, e := range s {
		fmt.Println(i, e)
	}
	str := "Hello World!"
	for _, ch := range str {

		fmt.Println(string(ch))
	}
	var sent string
	fmt.Scan(&sent)
	printPalindromes(sent + " ")
}
