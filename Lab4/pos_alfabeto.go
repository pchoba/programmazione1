package main

import (
	"fmt"
	"strconv"
)

/* Scrivere un programma pos_alfabeto.go che

legge una sequenza di caratteri ASCII (byte) terminati da '.' (es: abc!54LMN.),
per ciascun carattere letto (tranne '.'),

stabilisce se è una lettera minuscola (tra 'a' e 'z'), una cifra o altro
se è una lettera minuscola, stabilisce anche la sua posizione
se è una cifra, stabilisce il suo valore numerico
stampa un messaggio costituito da:

il carattere letto
se minuscola, la sua posizione nell'alfabeto (es "f è la 6^a");
se cifra, il suo valore
altrimenti " - altro" (ad esempio "N - altro")




quando termina, stampa la somma delle cifre e  "bye" */

func main() {
	var s string
	var somma int
	fmt.Print("Inserire una sequenza di caratteri terminata da '.': ")
	fmt.Scanf("%s", &s)
	for i := 0; i < len(s)-1; i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			fmt.Println(string(s[i]), "è la", i+1, "^a")
		} else if s[i] >= '0' && s[i] <= '9' {
			cifra, _ := strconv.Atoi(string(s[i]))
			fmt.Println(string(s[i]), "-", cifra)
			somma += cifra
		} else {
			fmt.Println(string(s[i]), "- altro")
		}
	}
	fmt.Printf("somma: %d\n", somma)
	fmt.Println("bye")
}
