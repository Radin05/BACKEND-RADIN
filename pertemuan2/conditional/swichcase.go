package main

import (
	"fmt"
	"time"
)

func main()  {
	hour := time.Now().Hour()

	switch {
	case hour < 10:
		fmt.Println("good morning")
	case hour < 18:
		fmt.Println("good afternoon")
	default:
		fmt.Println("good evening")
	}
}