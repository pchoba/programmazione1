package main

import (
	"fmt"
	"os"
	"strconv"
)

/* Scrivere una programma quadrati.go che legge da linea di comando una sequenza di numeri interi non negativi
e stampa solo quelli che sono dei quadrati (1, 4, 9, ecc.). Il programma deve essere dotato di una funzione
isSquare(n int) bool
che restituisce true se n è un quadrato, false altrimenti.*/

func isSquare(n int) bool {
	for i := 1; i <= 50; i++ {
		if n == i*i {
			return true
		}
	}
	return false
}
func main() {
	if len(os.Args) == 1 {
		os.Exit(1)
	}
	for i := 1; i < len(os.Args); i++ {
		num, err := strconv.Atoi(os.Args[i])
		if err != nil {
			os.Exit(0)
		}
		if isSquare(num) {
			fmt.Println(os.Args[i])
		}
	}
}
