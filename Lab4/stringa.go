package main

import "fmt"

/* Scrivere un programma stringa.go che legge da standard input una stringa e la analizza carattere
per carattere: stampa "+" quando il carattere considerato è maggiore del precedente, stampa "-" quando
è minore e stampa "=" quando è uguale. */

func main() {
	var s string
	fmt.Print("Inserire una stringa: ")
	fmt.Scan(&s)
	for i := 1; i < len(s); i++ {
		if s[i-1] < s[i] {
			fmt.Print("+")
		} else if s[i-1] == s[i] {
			fmt.Print("=")
		} else if s[i-1] > s[i] {
			fmt.Print("-")
		}
	}
	fmt.Println("")
}
