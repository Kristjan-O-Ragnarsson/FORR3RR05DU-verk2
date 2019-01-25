package main

import (
	"fmt"
)

func main() {
	fmt.Println("runa verk efni 4")
	var n int32
	fmt.Scanln(&n)
	x := summa(n)
	fmt.Println(x)
}

func summa(n int32) int32{
	if n == 1{
		return 1
	} else {
		return summa(n -1) +  n
	}
}
