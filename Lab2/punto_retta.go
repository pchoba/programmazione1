package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, m, q float64
	fmt.Print("Inserire m e q di una retta: ")
	fmt.Scan(&m, &q)
	fmt.Print("Inserire le coordinate di un punto: ")
	fmt.Scan(&x1, &y1)
	y_r := m*x1 + q
	if y_r > y1 {
		fmt.Println("sopra")
	} else if y_r < y1 {
		fmt.Println("sotto")
	} else if math.Abs(y_r-y1) <= 1e-6 {
		fmt.Println("sulla retta")
	}
}
