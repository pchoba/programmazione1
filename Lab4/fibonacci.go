package main

import "fmt"

/* Scrivere un programma fibonacci.go che legge un intero positivo n e stampa i numeri di
fibonacci dal primo all'n-esimo, rappresentandoli come righe di asterischi, ciascuna lunga quanto il numero da rappresentare. */

func main() {
	var a0, a1, b, n int
	a0 = 1
	a1 = 1
	fmt.Print("Inserire un numero: ")
	fmt.Scan(&n)
	fmt.Println("*")
	fmt.Println("*")
	for i := 2; i < n; i++ {
		b = a0 + a1
		for i := 0; i < b; i++ {
			fmt.Print("*")
		}
		fmt.Println("")
		a0, a1 = a1, b
	}

}
