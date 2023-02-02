package main

import (
	"fmt"
)

/* Scrivere un programma andamento.go che legge da tastiera una serie (di almeno un elemento) di numeri interi > -1 e
stampa "+" ogni volta che il nuovo valore è maggiore o uguale al precedente e "-" altrimenti.
Si ferma quando il numero in input è -1 e stampa la somma di tutti i numeri letti. */

func main() {
	var n, temp, diff int
	for {
		fmt.Scan(&n)
		if n == -1 {
			break
		}
		diff = n - temp
		if diff > 0 {
			fmt.Print("+")
		} else if diff == 0 {
			fmt.Print("=")
		} else {
			fmt.Print("-")
		}
		temp = n
	}
	fmt.Println("")
}
