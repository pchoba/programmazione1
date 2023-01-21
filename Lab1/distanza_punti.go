package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, x2, y1, y2 int
	fmt.Print("Coordinate del primo punto: ")
	fmt.Scan(&x1, &y1)
	fmt.Print("Coordinate del secondo punto: ")
	fmt.Scan(&x2, &y2)
	distanza := math.Sqrt(math.Pow(float64(x2-x1), 2) + math.Pow(float64(y2-y1), 2))
	fmt.Println("La distanza tra i due punti è", distanza)
}
