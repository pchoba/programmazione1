package main

import "fmt"

/* Scrivere un programma cesare.go che legge da standard input un valore intero non negativo k
(la chiave di cifratura) e una sequenza di lettere minuscole consecutive (sulla stessa riga e senza spazi) terminate da invio.
Il programma stampa la sequenza letta cifrata secondo il cifrario di Cesare, usando come chiave k (quella fornita dall'utente):
ogni lettera del testo in chiaro è sostituita nel testo cifrato dalla lettera che si trova k posizioni dopo nell'alfabeto,
ritornando alla lettera a dopo la zeta. */

func main() {
	var k int
	var s string
	var active bool
	active = true
	fmt.Print("chiave: ")
	fmt.Scan(&k)
	fmt.Print("caratteri da cifrare: ")
	fmt.Scan(&s)
	for i := 0; i < len(s); i++ {
		if s[i]+byte(k) >= 'a' && s[i]+byte(k) <= 'z' {
			fmt.Printf("%c", s[i]+byte(k))
		} else {
			for active {
				if s[i]+byte(k) >= 'a' && s[i]+byte(k) <= 'z' {
					fmt.Printf("%c", s[i]+byte(k))
					active = false
				} else {
					k -= 26
				}
			}
		}
	}
	fmt.Println("")
}
