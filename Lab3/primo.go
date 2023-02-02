package main

import (
	"fmt"
	"os"
)

/* Scrivere un programma primo.go che, dato un numero intero su standard input, determina se il numero è primo.
Suggerimento: occorre determinare il primo numero che è un divisore (se c'è).
Domanda: dato in input n, quante iterazioni dovrò fare al massimo?*/

func main() {
	var n int
	fmt.Print("Inserire un numero: ")
	fmt.Scan(&n)
	for i := 2; i < n; i++ {
		if n%i == 0 {
			fmt.Println("non è primo")
			os.Exit(0)
		}
	}
	fmt.Println("è primo")
}
