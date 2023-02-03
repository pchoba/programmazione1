package main

import "fmt"

/*  Scrivere un programma max_char.go che legge da standard input una sequenza di 5 caratteri ASCII (byte)
e stampa il maggiore in ordine lessicografico (cioè con il codice ASCII più alto). */

func main() {
	var s string
	var max int
	var max_byte byte
	fmt.Print("Inserire una sequenza di cinque caratteri ASCII: ")
	fmt.Scanf("%s", &s)
	for i := 0; i < len(s); i++ {
		if int(s[i]) > max {
			max = int(s[i])
			max_byte = s[i]
		}
	}
	fmt.Println("il maggiore in ordine lessicografico è", string(max_byte))
}
