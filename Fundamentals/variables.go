package main

import "fmt"

func main() {
	// var <name> <type>
	var num int
	// types have default values
	num += 1
	fmt.Println(string(num))
	var n = 2
	n -= 1
	// we can declare multipl variables in the same line
	var a, b, c = 3, 4, "hello"
	a = b + a + 32
	fmt.Println(string(a)) //string of a int yields a rune
	c += string(num)
	fmt.Println(c)
	// can also declare using blocks
	var (
		x float32 = 1212221
		y rune    = 3
		z         = "Hello"
	)
	x += float32(n)
	fmt.Println(string(int(x)+int(y)) + z)
	// special syntac
	p, q, _ := 1, 2, 3
	//_ will be ignored
	p = q
	fmt.Println(p)
	//:= is the create and assign operator
	//to make it easier to assign errors following is allowed
	s, err := "Somneel", num
	d, err := 4, num
	err, e := 4, b
	println(s)
	println(d + err + e)
	//declaring constants
	const Pi float32 = (22 / 7)
	fmt.Println(Pi * x)
	//we can ignore variables also

}
