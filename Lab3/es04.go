package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*Scrivere un programma es04.go che genera K = 10 numeri casuali in [0,10] e li stampa su una riga, separati da spazi.*/

func main() {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(11), " ")
	}
}
