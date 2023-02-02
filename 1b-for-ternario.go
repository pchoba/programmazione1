package main

import "fmt"

func main() {
	/*
	   dati cinque numeri scelti dall'utente, stampa la somma dei cinque numeri

	*/
	const K = 5
	var n int
	somma := 0
	for i := 1; i <= K; i++ {
		fmt.Print("un numero: ")
		fmt.Scan(&n)
		somma = somma + n
	}
	fmt.Println(somma)
}
