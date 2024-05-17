package main

import (
	"fmt"
	"math/rand"
)

func main()  {
	var pemain1, pemain2 string
	var pilihanPemain1, pilihanPemain2 int

	fmt.Println("pilih apa?")
	fmt.Println("1 gunting")
	fmt.Println("2 batu")
	fmt.Println("3 kertas")

	fmt.Scanln(&pilihanPemain1)

	switch pilihanPemain1 {
	case 1:
		pemain1 = "gunting"
	case 2:
		pemain1 = "batu"
	case 3:
		pemain1 = "kertas"
	default:
		fmt.Println( "ga milih")
	}
	
	pilihanPemain2 = rand.Intn(3)

	switch pilihanPemain2 {
	case 1:
		pemain2 = "gunting"
	case 2:
		pemain2 = "batu"
	case 3:
		pemain2 = "kertas"
	default:
		fmt.Println("ga milih")
	}

	fmt.Printf("pilihan pemain 1 = %s \n", pemain1)
	fmt.Printf("pilihan pemain 2 = %s \n", pemain2) 

	if pemain1 == "kertas" && pemain2 == "gunting"{
		fmt.Println("pemain 2 menang")
	}else if pemain1 == "gunting" && pemain2 == "batu"{
		fmt.Println("pemain 2 menang")
	}else if pemain1 == "batu" && pemain2 == "kertas"{
		fmt.Println("pemain 2 menang")
	}else if pemain1 == "batu" && pemain2 == "gunting"{
		fmt.Println("pemain 1 menang")
	}else if pemain1 == "kertas" && pemain2 == "batu"{
		fmt.Println("pemain 1 menang")
	}else if pemain1 == "gunting" && pemain2 == "kertas"{
		fmt.Println("pemain 1 menang")
	}else if pemain1 == pemain2{
		fmt.Println("draw")
	}else{
		fmt.Println("gagal")
	}
}
