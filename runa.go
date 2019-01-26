package main

import (
	"fmt"
)

func main() {
	fmt.Println("runa verk efni 4")
	var n int32
	fmt.Scanln(&n)
	summa(n)
}

func summa(n int32) int32{
	if n == 1{
		fmt.Printf("1 ")
		return 1
	} else {
		su := summa(n -1) +  n
		fmt.Printf("%v ", su)
		return su
	}
}
