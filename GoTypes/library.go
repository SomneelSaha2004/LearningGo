//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type Book struct {
	name       string
	checkedOut bool
	t          float64
	owner      string
}
type Member struct {
	name string
	book Book
}

func checkin(m *Member, b *Book) {
	m.book = Book{}
	b.checkedOut = false
	b.t = float64(time.Now().UnixMilli())
	b.owner = ""
}
func checkout(m *Member, b *Book) {
	if b.checkedOut {
		fmt.Println(b.name, "has already been checked out by", b.owner)
		return
	}
	m.book = *b
	b.checkedOut = true
	b.t = float64(time.Now().UnixMilli())
	b.owner = m.name
}
func printBooks(arr []Book) {
	status := ""
	for _, b := range arr {
		if b.checkedOut {
			status = "Checked out"
			fmt.Println(b.name, status, "at", b.t, "by", b.owner)
		} else {
			fmt.Println(b.name, "is in stock")
		}

	}
	fmt.Println()
}
func printMembers(m []Member) {
	status := ""
	for _, m2 := range m {
		if m2.book.name == "" {
			status = "No books issued"
		} else {
			status = m2.book.name
		}
		fmt.Println(m2.name, "-->", status)
	}
	fmt.Println()
}
func main() {
	books := []Book{{"a", false, 0, ""}, {"b", false, 0, ""}, {"c", false, 0, ""}, {"d", false, 0, ""}}
	members := []Member{{name: "Mary"}, {name: "Sam"}, {name: "Ken"}}
	checkout(&members[1], &books[3])
	printBooks(books)
	printMembers(members)
	checkout(&members[0], &books[3])
	checkout(&members[2], &books[0])
	checkin(&members[1], &books[3])
	printBooks(books)
	printMembers(members)
}
