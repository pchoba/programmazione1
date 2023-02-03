package main

import "fmt"

/* Scrivere un programma lunghezza_tot.go che legge da standard input un int totLen e una sequenza di stringhe (una per riga)
sommandone le lunghezze fino a raggiungere (o superare) totLen.
Raggiunto totLen, il programma stampa la somma delle lunghezze e la concatenazione delle stringhe lette. */

func main() {
	var totLen int
	var concat_strings, str string
	fmt.Print("inserire un numero: ")
	fmt.Scanf("%d", &totLen)
	for len(concat_strings) < totLen {
		fmt.Scan(&str)
		concat_strings += str
	}
	fmt.Printf("somma delle lunghezze: %d\nconcatenazione delle stringhe lette: %s\n", len(concat_strings), concat_strings)
}
