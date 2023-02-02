package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*Stampa quante volte un numero randomico da [0, 20] esce diverso da 0 */
	rand.Seed(time.Now().Unix())
	const K = 20
	var n int

	c := 0
	for {
		n = rand.Intn(K + 1)
		fmt.Print(n, " ")
		if n == 0 {
			break
		}
		c++
	}
	fmt.Println(c)
}
