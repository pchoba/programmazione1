package main

import "fmt"

/* Scrivere un programma es0.go che fa le seguenti cose:

legge un byte
lo stampa
stampa il byte precedente, il byte stesso, e quello successivo in ordine lessicografico (ASCII).
Ad esempio, con 'd' come input, deve stampare: c d e

stabilisce se è una lettera tra A e L, o altro (stampa "A-L" o "altro")
poi legge una stringa e la stampa in verticale, una runa per riga. */

func main() {
	var b byte
	var s string
	fmt.Print("Inserire un carattere byte: ")
	fmt.Scanf("%c", &b)
	fmt.Println(string(b - 1))
	fmt.Println(string(b))
	fmt.Println(string(b + 1))
	if b >= 'A' && b <= 'L' {
		fmt.Println("A-L")
	} else {
		fmt.Println("altro")
	}
	fmt.Print("Inserire una stringa: ")
	fmt.Scan(&s)
	for i := 0; i < len(s); i++ {
		fmt.Println(string(s[i]))
	}
}
