package main

import (
	"fmt"
	"math"
)

func main() {
	var num1, den1, num2, den2 int
	fmt.Print("num e den fraz 1: ")
	fmt.Scan(&num1, &den1)
	fmt.Print("num e den fraz 2: ")
	fmt.Scan(&num2, &den2)
	if math.Abs(float64((num1/den1)-(num2/den2))) <= 1e-6 {
		fmt.Println("equivalenti")
	} else {
		fmt.Println("non equivalenti")
	}
}
