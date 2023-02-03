package main

import (
	"fmt"
	"strconv"
)

/* Si vuole scrivere un programma max_somma_cifre.go che

legge da standard input una serie di numeri >= 0, terminata da 999,
trova il numero (escludendo 999) la cui somma delle cifre è la maggiore
e stampa tale somma. */

func main() {
	var num_str, max_num string
	var somma, max int
	for {
		fmt.Print("Inserire un numero: ")
		fmt.Scan(&num_str)
		if num_str == "999" {
			break
		}
		somma = 0
		for i := 0; i < len(num_str); i++ {
			cifra, _ := strconv.Atoi(string(num_str[i]))
			somma += cifra
		}
		if somma > max {
			max = somma
			max_num = num_str
		}
	}
	fmt.Printf("Il numero la cui somma delle cifre è la maggiore è %s\n", max_num)
}
