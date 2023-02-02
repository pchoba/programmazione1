package main

import "fmt"

func main() {
	/*
	   Dato un numero inserito in standard input, stampare tutti i numeri pari da 0 (compreso), fino al numero (compreso)
	*/
	var n int
	fmt.Print("un numero: ")
	fmt.Scan(&n)
	for i := 0; i <= n; i = i + 2 {
		fmt.Print(i, " ")
	}
	fmt.Println()
}
