package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/* Scrivere un programma ordini.go che legge da standard input una serie di stringhe che descrivono ordini nel formato
prezzo#quantità#sconto
e stampa il totale finale da pagare.
Prezzo, quantità e sconto sono float; prezzo indica il prezzo unitario del prodotto, quantità indica la quantità acquistata
e sconto è lo sconto applicato per quel prodotto, espresso come float tra 0 e 1 (ad esempio 0.2 indica uno sconto del 20%).
Il programma termina la lettura quando incontra EOF (ctrl D su riga nuova). */

func main() {
	var totale float64
	var prodotto []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		prodotto = strings.Split(text, "#")
		prezzo, _ := strconv.ParseFloat(string(prodotto[0]), 64)
		quantità, _ := strconv.ParseFloat(string(prodotto[1]), 64)
		sconto, _ := strconv.ParseFloat(string(prodotto[2]), 64)
		totale += (prezzo * quantità) - (prezzo*quantità)*sconto
	}
	fmt.Println(totale)
}
