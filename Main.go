package main

import (
	"fmt"
)

//Sqrt root function
func Sqrt(x float64) float64 {
	z := 1.0
	for z*z != x {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func main() {
	var x float64
	fmt.Println("Enter the number")
	fmt.Scanln(&x)
	go saturday.main()
	fmt.Println("The root of ", x, "is", Sqrt(x))
}
