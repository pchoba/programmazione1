package main

import (
	"fmt"
	"strconv"
)

/* Scrivere un programma num_sequenze.go che legge da standard input una sequenza di uni (1) e zeri (0) (terminata da un 2),
che inizia e finisce con 1, e stampa il numero di sottosequenze di zeri.
Ad esempio per input 1 1 0 0 1 0 1 0 0 0 1 1 1 0 1 2, l'output è 4. (si considerano anche le sottosequenze di lunghezza 1) */

func main() {
	var num int
	var num_string string
	for {
		fmt.Scan(&num)
		if num == 2 {
			break
		}
		num_string += strconv.Itoa(num)
	}
	num_string = num_string[1 : len(num_string)-2]
	fmt.Println(num_string)
}
