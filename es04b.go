package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*Scrivere un programma es04b.go che genera K = 10 numeri casuali in [0,10], conta quanti sono pari e stampa questo risultato.*/

func main() {
	const K = 10
	var counter int
	rand.Seed(time.Now().Unix())
	for i := 0; i < K; i++ {
		n := rand.Intn(11)
		if n%2 == 0 {
			counter++
		}
	}
	fmt.Println(counter)
}
