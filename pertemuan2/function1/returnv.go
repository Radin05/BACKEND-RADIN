package main

import(
	"fmt"
)

func getHello() (string,string) {
	return "hello","hello"
}

func main()  {
	x, y := getHello()
	fmt.Println(x,y)
}