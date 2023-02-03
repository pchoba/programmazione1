package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/* Scrivere un programma ultima_pioggia.go che legge da standard input
una sequenza di valori interi (terminata da EOF, da tastiera prodotto con ctrl-D)
che indicano i mm di pioggia caduti (0 se non ha piovuto) ogni giorno in una sequenza
successiva di giorni e stampa (l'indice del)l'ultimo giorno in cui ha piovuto. */

func main() {
	slc := []int{}
	var text_num int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			text_num = 0
		}
		text_num, _ = strconv.Atoi(text)
		slc = append(slc, text_num)
	}
	fmt.Println(len(slc) - 1)
}
