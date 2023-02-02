package main

import (
	"fmt"
)

/*Scrivere un programma quadrati.go che legge da standard input un intero n positivo e calcola,
utilizzando solo variabili di tipo int, il massimo quadrato (k^2) <= n.
Per calcolare il quadrato di un numero n usate n*n.*/

func main() {
	var n int
	var a int
	fmt.Print("un numero: ")
	fmt.Scan(&n)
	for i := n; i > 1; i-- {
		if i*i <= n {
			a = i
			break
		}
	}
	fmt.Printf("Il massimo quadrato minore di %d è %d\n", n, a*a)
}
