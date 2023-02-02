package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*Scrivere un programma es10.go che genera numeri casuali tra 1 e 10, stampandoli, fino a quando la
loro somma raggiunge un obiettivo fissato (TARGET), ad esempio 50. Poi stampa la somma.*/

func main() {
	const TARGET = 50
	var somma int
	rand.Seed(time.Now().Unix())
	for somma < TARGET {
		n := rand.Intn(10) + 1
		somma += n
	}
	fmt.Println(somma)
}
