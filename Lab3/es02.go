package main

import "fmt"

/*Scrivere un programma es02.go che legge K = 5 numeri e di ciascuno stampa il doppio.*/

func main() {
	const K = 5
	var n int
	for i := 0; i < K; i++ {
		fmt.Print("un numero: ")
		fmt.Scan(&n)
		fmt.Println(n * 2)
	}
}
