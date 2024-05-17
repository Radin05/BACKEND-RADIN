package main

import(
	"fmt"
)

func main()  {
	
	// for1
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// for2
	i := 0
	for i < 10{
		fmt.Println(i)
		i++
	}

	// for with break expression
	for i:= 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
	
	// for with continue expression
	for i:= 0; i < 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}

	// for with range expression
	names := []string{"jhon","paul","george","ringo"}
	for i, name := range names {
		fmt.Println(i,"=",name)
	}
}