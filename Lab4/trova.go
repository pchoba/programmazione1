package main

import (
	"fmt"
	"os"
)

/* Scrivere un programma trova.go che legge da standard input un carattere (runa) e
una stringa e stampa la posizione del carattere nella stringa (la prima volta che appare), o -1 se il carattere non c'è. */

func main() {
	var r rune
	var s string
	fmt.Print("Inserire un carattere 'rune': ")
	fmt.Scanf("%c", &r)
	fmt.Print("Inserire una stringa: ")
	fmt.Scan(&s)
	for i := 0; i < len(s); i++ {
		if rune(s[i]) == r {
			fmt.Println(i)
			os.Exit(0)
		}
	}
	fmt.Println(-1)
}
