package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
	   Data la somma di numeri randomici da 1 a 10, stampare la somma appena raggiunge almeno 20
	*/
	const TARGET = 20
	const MAX = 10
	rand.Seed(time.Now().Unix())
	var n int

	t := 0
	for t < TARGET {
		n = rand.Intn(MAX) + 1
		fmt.Print(n, " ")
		t += n
	}
	fmt.Println(t)
}
