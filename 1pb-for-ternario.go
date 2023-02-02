package main

import "fmt"

func main() {
	/*Dato un numero n, stampa tutti i numeri da n a 1 in ordine decresente*/
	var n int
	fmt.Print("un numero: ")
	fmt.Scan(&n)
	for i := n; i >= 1; i-- {
		fmt.Print(i, " ")
	}
	fmt.Println()
}
