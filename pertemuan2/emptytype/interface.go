package main

import(
	"fmt"
)

func ups() interface{}{
	return "Ups"
}

func main()  {
	data := ups()
	fmt.Println(data)
}