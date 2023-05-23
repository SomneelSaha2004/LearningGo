package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	fmt.Println("Welcome to Rock Paper Scissors with GO!")
	fmt.Println("R -> Rock		S -> Scissor		P -> Paper\nQ -> quit")
	comp := [3]string{"R", "S", "P"}

	for {
		rand.Seed(time.Now().UnixNano())
		fmt.Print("--> ")
		var input string
		fmt.Scanln(&input)
		input = strings.TrimRight(input, "\n")
		input = strings.TrimLeft(input, " ")
		fmt.Println((input))
		if input == "Q" {
			fmt.Println("Thanks for playing!")
			break
		}
		output := comp[rand.Intn(3)]
		switch input {
		case "R":
			switch output {
			case "R":
				fmt.Println("Computer threw", string(output), "\nIts a Draw!")
			case "S":
				fmt.Println("Computer threw", string(output), "\nYou Win!")
			case "P":
				fmt.Println("Computer threw", string(output), "\nComputer wins!")
			}
		case "S":
			switch output {
			case "S":
				fmt.Println("Computer threw", string(output), "\nIts a Draw!")
			case "P":
				fmt.Println("Computer threw", string(output), "\nYou Win!")
			case "R":
				fmt.Println("Computer threw", string(output), "\nComputer wins!")
			}
		case "P":
			switch output {
			case "P":
				fmt.Println("Computer threw", string(output), "\nIts a Draw!")
			case "R":
				fmt.Println("Computer threw", string(output), "\nYou Win!")
			case "S":
				fmt.Println("Computer threw", string(output), "\nComputer wins!")
			}
		}
	}
}
