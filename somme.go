package main

import (
	"fmt"
	"math/rand"
)

/*Scrivere un programma somme.go che genera 10 numeri interi casuali tra 0 e 10, li stampa, e stampa la somma dei valori generati.
nomefile: somme.go*/

func main() {
	const TARGET = 10
	var sum int
	for i := 0; i < TARGET; i++ {
		n := rand.Intn(11)
		fmt.Println(n)
		sum += n
	}
	fmt.Println("somma:", sum)
}
