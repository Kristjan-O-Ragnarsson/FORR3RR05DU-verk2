package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Summa verk efni 3")
	var n int32
	fmt.Scanln(&n)
	x := summa(n)
	fmt.Println(x)
}

func summa(n int32) int32{
	if n == 1{
		return 1
	} else {
		return summa(n -1) + int32(math.Pow(float64(n),2))
	}
}