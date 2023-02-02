package main

import (
	"fmt"
)

/*Scrivere un programma es08.go che legge un numero n e stampa la tabellina di n (cioè n*1, n*2, ..., n*10)*/

func main() {
	var n int
	var moltiplicando int
	moltiplicando = 1
	fmt.Print("un numero: ")
	fmt.Scan(&n)
	for i := 1; i < 11; i++ {
		fmt.Println(n, "*", moltiplicando, "=", n*moltiplicando)
		moltiplicando++
	}
	fmt.Println("")
}
