package main

import "fmt"

func main() {
	var peso, altezza float64
	fmt.Print("Inserire peso: ")
	fmt.Scanf("%f", &peso)
	fmt.Print("Inserire altezza (in cm): ")
	fmt.Scanf("%f", &altezza)
	if peso >= 10 && peso <= 50 {
		if altezza >= 100 && altezza <= 150 {
			fmt.Println("normopeso")
		} else if altezza > 150 && altezza <= 200 {
			fmt.Println("sottopeso")
		}
	} else if peso > 50 && peso <= 100 {
		if altezza >= 100 && altezza <= 150 {
			fmt.Println("sovrappeso")
		} else if altezza > 150 && altezza <= 200 {
			fmt.Println("normopeso")
		}
	} else if peso > 100 {
		if altezza >= 100 && altezza <= 150 {
			fmt.Println("molto sovrappeso")
		} else if altezza > 150 && altezza <= 200 {
			fmt.Println("sovrappeso")
		}
	}
}
