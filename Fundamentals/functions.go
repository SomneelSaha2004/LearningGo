package main

import "fmt"

func FizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		if i%5 == 0 && i%3 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

// we can return multipl values from functions(most likely via tuple unpacking)
func getFive() (int, int, float32, string) {
	return 1, 2, 3.4, "Yeet"

}
func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
func fibonacci(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	a, b, c := 0, 1, 1
	for i := 2; i <= n-1; i++ {
		c = a + b
		a = b
		b = c
	}
	return c
}
func sumOfPrimes(n int) int {
	sum := 0
	for i, count := 2, 1; count <= n; i++ {
		if isPrime(i) {
			fmt.Println(i)
			count += 1
			sum += i
		}
	}
	return sum
}
func main() {
	fmt.Println(fibonacci(6))
	FizzBuzz(100)
	fmt.Println(sumOfPrimes(5))
	a, b, _, s := getFive()
	fmt.Println((a + b), s)
}
