package main

import (
	"fmt"
	"math"
)

func main() {
	const EPSILON = 1e-5
	var x, y float64
	fmt.Print("Inserire x e y di un punto: ")
	fmt.Scan(&x, &y)
	dist := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
	if dist < EPSILON {
		fmt.Println("vicino all'origine")
	} else {
		fmt.Println("non vicino all'origine")
	}
}
