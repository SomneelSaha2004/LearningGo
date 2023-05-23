//--Summary:
//  Create a program that can perform dice rolls using various configurations
//  of number of dice, number of rolls, and number of sides on the dice.
//  The program should report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these circumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must use variables to configure:
//  - number of times to roll the dice
//  - number of dice used in the rolls
//  - number of sides of all the dice (6-sided, 10-sided, etc determined
//    with a variable). All dice must use the same variable for number
//    of sides, so they all have the same number of sides.
//
//--Notes:
//* Use packages from the standard library to complete the project
//* Try using different values for your variables to check the results

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func roll(noOfDice int, sides int) int {
	sum := 0
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= noOfDice; i++ {
		sum += rand.Intn(sides) + 1
	}
	return sum
}
func main() {
	var noOfDice, noOfRolls, sides int
	for {
		fmt.Print("No of dice : ")
		fmt.Scan(&noOfDice)
		fmt.Print("No of rolls : ")
		fmt.Scan(&noOfRolls)
		fmt.Print("No of sides of each dice : ")
		fmt.Scan(&sides)
		for i := 1; i <= noOfRolls; i++ {
			result := roll(noOfDice, sides)
			fmt.Println(result)
			if noOfDice == 2 && result == 2 {
				fmt.Println("Snake eyes")
			} else if result == 7 {
				fmt.Println("Lucky Seven")
			} else if result%2 == 0 {
				fmt.Println("Even")
			} else {
				fmt.Println("Odd")
			}
		}
	}
}
