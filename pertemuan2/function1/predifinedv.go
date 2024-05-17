package main

import(
	"fmt"
)

func getName() (firstName, lastName string) {
	return "hello","hello"
}

func main()  {
	x, y := getName()
	fmt.Println(x,y)
}