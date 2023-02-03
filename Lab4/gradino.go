package main

import "fmt"

/* Definiamo "gradino" una sequenza di (uno o più) interi non negativi uguali seguita da un'altra
sequenza di (uno o più) interi più grandi di 1 (es. 1 1 2 2 2).
Scrivere un programma gradino.go che, data in input una sequenza di interi tali che ogni intero è
>= del precedente, stampa la lunghezza (il numero di interi) del gradino più lungo. (Si noti che i gradini si sovrappongono).
Il programma termina quando legge un numero minore di quello appena letto. */

func main() {
	var n, n1, count, count_max int
	for {
		fmt.Scan(&n)
		if n < n1 {
			break
		} else if n == n1 {
			count++
			n1 = n
			if count > count_max {
				count_max = count
			}
		} else if n > n1 {
			n1 = n
			count = 1
			continue
		}
	}
	fmt.Printf("\n%d\n", count_max)
}
