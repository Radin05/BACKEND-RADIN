package main

import(
	"fmt"
)

func helloParameter(firstName, lastName string)  {
	fmt.Println("hello", firstName, lastName)
}
func main()  {
	helloParameter("jhon","ganteng")
}