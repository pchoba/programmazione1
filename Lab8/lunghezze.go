package main

/*Scrivere un programma lunghezze.go che legge riga per riga un testo da standard input (potete usare la ridirezione),
terminato da EOF, e stampa quante parole ci sono di lunghezza 1, quante di lunghezza 2, ecc.
Il programma è dotato di una funzione*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func aggiornaConteggio(m map[int]int, riga string) {
	lenght := len(riga)
	m[lenght] += 1
}

func main() {
	var lunghezze map[int]int
	var max int
	lunghezze = make(map[int]int)
	var lara []string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()
		lara = strings.Split(text, " ")
		if len(lara) == 0 {
			continue
		}
		for i := 0; i < len(lara); i++ {
			aggiornaConteggio(lunghezze, lara[i])
			if len(lara[i]) > max {
				max = len(lara[i])
			}

		}
	}
	/*fmt.Println(lunghezze)
	fmt.Println(len(lunghezze))*/
	for j := 1; j <= max; j++ {
		if lunghezze[j] != 0 {
			fmt.Println("parole di lunghezza", j, ":", lunghezze[j])
		}
	}
	/*for k, v := range lunghezze {
		fmt.Printf("Le parole di lunghezza %v sono %v\n", k, v)
	}*/
}
