package main

import (
	"fmt"
	"strconv"
)

/*Scrivere un programma somma_cifre.go che calcola la somma delle cifre di un numero intero positivo fornito da standard input.*/

func main() {
	var n, somma int
	fmt.Print("Inserire un numero: ")
	fmt.Scan(&n)
	num_str := strconv.Itoa(n)
	for i := 0; i < len(num_str); i++ {
		cifra, _ := strconv.Atoi(string(num_str[i]))
		somma += cifra
	}
	fmt.Printf("La somma delle cifre di %d è %d\n", n, somma)

}
