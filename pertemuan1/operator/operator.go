package main

import (
	"fmt"
)

// operator aritmatika

func main(){
	a, b := 5, 10

	// operator aritmatika 
	c := a + b
	d := a - b
	e := a * b
	f := a / b
	g := a % b
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	
	// augmented assignments
	a += b
	b -= a
	b /= 5
	a *= b
	a %= b
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(b)
	fmt.Println(a)
	fmt.Println(a)

	// unary operator
	a++
	fmt.Println(a)
	a--
	fmt.Println(a)

	// comparison operator
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a < b)
	fmt.Println(a > b)

	// logical operator
	fmt.Println(a < 1 && b == 10)
	fmt.Println(a > 1 || b == 100)
	fmt.Println(!(a > 0))
	fmt.Println(!(a == 1) || (b == 100))
}

// <-------------------------------------------->

// func main() {
// 	var numbers [4]string

// 	fmt.Printf("%v/n", numbers)
// 	fmt.Printf("%#v/n", numbers)
// }
