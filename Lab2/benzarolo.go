package main

import "fmt"

func main() {
	const benzina = 0
	const diesel = 1
	const alcol = 2
	const cherosene = 3
	var litri float64
	var carb int
	fmt.Print("Inserire quantità del rifornimento: ")
	fmt.Scan(&litri)
	fmt.Print("Inserire tipo di carburante: ")
	fmt.Scan(&carb)
	if carb == 0 {
		fmt.Println("Prezzo:", benzina*litri, "€")
	} else if carb == 1 {
		fmt.Println("Prezzo:", diesel*litri, "€")
	} else if carb == 2 {
		fmt.Println("Prezzo:", alcol*litri, "€")
	} else if carb == 3 {
		fmt.Println("Prezzo:", cherosene*litri, "€")
	}
}
