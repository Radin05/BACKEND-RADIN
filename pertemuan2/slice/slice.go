package main

import (
	"fmt"
)

func main()  {
	numbers := []int{1,2,3,4}
	fmt.Println(numbers)

	nums := make([]int, 2)
	fmt.Printf("%#v\n", nums)
	
	friends := []string{"abdu","zapran","zidni","wandi"}
	fmt.Printf("%v\n", friends)

	myBestFriend := friends[0:4]
	fmt.Println("my best friend is ", myBestFriend)
	fmt.Printf("%#v\n", myBestFriend)

	for index, value := range friends {
		fmt.Printf("index: %v, value: %v\n", index,value)		
	}
}