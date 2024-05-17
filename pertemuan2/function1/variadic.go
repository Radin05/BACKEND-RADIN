package main

import(
	"fmt"
)

func sumAll(numbers ... int) (sum int) {
	for _,numbers := range numbers {
		sum += numbers
	}
	return
}

func main()  {
	total := sumAll(20,2,2,2)
	fmt.Println(total)
}