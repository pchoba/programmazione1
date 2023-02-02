package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*Scrivere un programma massimo.go che genera 10 numeri interi casuali tra 10 e 30, li stampa, e stampa il massimo valore generato.*/

func main() {
	var max int
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		n := rand.Intn(21) + 10
		fmt.Println(n)
		if n > max {
			max = n
		}
	}
	fmt.Println("massimo: ", max)

}
