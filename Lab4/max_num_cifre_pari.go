package main

import (
	"fmt"
	"strconv"
)

/* Scrivere un programma max_num_cifre_pari.go che, data una sequenza di numeri (da leggere come stringhe),
terminata dalla stringa "000", stampa il numero di cifre pari contenute nel numero che ne contiene più di tutti.
E` possibile risolvere il problema senza tenere in memoria sequenze di numeri. Che tipo di composizione occorrerà
per mettere insieme il conteggio delle cifre pari e il calcolo del massimo di tali conteggi? */

func main() {
	var num_str string
	var cont_even, max_even int
	for {
		fmt.Print("Inserire un numero: ")
		fmt.Scan(&num_str)
		if num_str == "000" {
			break
		}
		cont_even = 0
		for i := 0; i < len(num_str); i++ {
			cifra, _ := strconv.Atoi(string(num_str[i]))
			if cifra%2 == 0 {
				cont_even++
			}
		}
		if cont_even > max_even {
			max_even = cont_even
		}
	}
	fmt.Println("Il numero di cifre pari maggiore è", max_even)
}
