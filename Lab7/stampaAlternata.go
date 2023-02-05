package main

import (
	"bufio"
	"fmt"
	"os"
)

/* Scrivere un programma stampaAlternata.go che legge da standard input del testo su
più righe (terminato da EOF) e stampa prima le righe pari e poi le righe dispari
(considerate la prima riga del testo la riga 1 (e non 0)). */

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	count := 1
	righe_pari := []string{}
	righe_dispari := []string{}
	for scanner.Scan() {
		text := scanner.Text()
		if count%2 == 0 {
			righe_pari = append(righe_pari, text)
		} else {
			righe_dispari = append(righe_dispari, text)
		}
		count++

	}
	for i := 0; i < len(righe_pari); i++ {
		fmt.Println(righe_pari[i])
	}
	for i := 0; i < len(righe_dispari); i++ {
		fmt.Println(righe_dispari[i])
	}
}
