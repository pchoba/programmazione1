package main

import "fmt"

/* Scrivere un programma crescente.go che legge da standard input una stringa di sole lettere minuscole e la stampa inserendo un '-'
ogni volta che una lettera viene prima in ordine alfabetico della lettera precedente.
Per esempio, data in input la parola ambire, il programma stampa
am-bir-e */

func main() {
	var s, s1 string
	fmt.Print("Inserire una stringa di sole lettere minuscole: ")
	fmt.Scan(&s)
	for i := 1; i < len(s); i++ {
		if s[i] < s[i-1] {
			s1 = s[:i] + "-" + s[i:]
		}
	}
	fmt.Println(s1)
}
