package main

import "fmt"

func main() {
	/*
	 Dato un certo numero inserito dall'utente, stampa tanti asterischi quanto il numero messo
	*/
	var n int
	fmt.Print("un numero: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}
