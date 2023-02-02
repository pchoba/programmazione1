package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

/* Scrivere un programma bin2ten.go che converte un intero composto di soli 0 e 1 nel valore corrispondente al numero binario
rappresentato (es. 101 -> 5). Nel caso il numero contenga altre cifre, il programma stampa un messaggio di errore. */

func main() {
	var bin, esp, somma int
	fmt.Print("Inserire un numero binario: ")
	fmt.Scan(&bin)
	bin_str := strconv.Itoa(bin)
	for i := 0; i < len(bin_str); i++ {
		if string(bin_str[i]) != "1" && string(bin_str[i]) != "0" {
			fmt.Println("errore")
			os.Exit(1)
		}
		cifra_bin, _ := strconv.Atoi(string(bin_str[i]))
		if cifra_bin == 0 {
			esp++
			continue
		} else if cifra_bin == 1 {
			somma += int(math.Pow(2, float64(esp)))
			esp++
		}
	}
	fmt.Printf("Il numero decimale associato a %d è %d\n", bin, somma)
}
