package main

import (
	"fmt"
	"os"
)

/*Scrivere un programma differenze.go che legge una serie di valori float64 da tastiera e stampa le differenze, cioè la
differenza tra il secondo e il primo, tra il terzo e il secondo, e così via.
Il programma termina quando riceve in input il valore 0.*/

func main() {
	var a, b, r float64
	fmt.Print("due numeri: ")
	fmt.Scan(&a, &b)
	if a != 0 && b != 0 {
		r = a - b
		fmt.Println("differenza: ", r)
	} else {
		os.Exit(1)
	}
	for b != 0 {
		fmt.Print("un numero: ")
		fmt.Scan(&b)
		r -= b
		fmt.Println("differenza: ", r)
	}
}
