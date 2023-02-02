package main

import (
	"fmt"
	"os"
	"strconv"
)

/* Scrivere un programma cifra_pari.go che, dato un intero da standard input, determina e stampa in che posizione
(procedendo da destra a sinistra) si trova la prima cifra pari del numero. Se il numero non contiene cifre pari,
il programma stampa -1.
nomefile: cifra_pari.go */

func main() {
	var n int
	fmt.Print("Inserire un numero: ")
	fmt.Scan(&n)
	n_string := strconv.Itoa(n)
	for i := len(n_string) - 1; i >= 0; i-- {
		cifra_n, _ := strconv.Atoi(string(n_string[i]))
		if cifra_n%2 == 0 {
			fmt.Println(i + 1)
			os.Exit(0)
		}
	}
	fmt.Println("-1")
}
