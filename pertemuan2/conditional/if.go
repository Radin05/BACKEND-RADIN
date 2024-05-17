package main 

import(
	"fmt"
)

func main()  {
	age := 18

	if age >= 0 && age < 16 {
		fmt.Printf("you cannot vote! please return in %d years!\n", 18-age)
	} else if age == 18 {
		fmt.Printf("this is your first vote!")
	}else if age > 18 && age <= 100 {
		fmt.Printf("vote! it's important!")
	}else{
		fmt.Printf("invalid age")
	}
}