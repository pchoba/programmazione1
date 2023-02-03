package main

import "fmt"

/*
	Scrivere un programma disegna_slash.go che legge un intero positivo n

e stampa un backslash (\) di altezza n composto da asterischi.
*/
func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}
		fmt.Print("*")
		fmt.Println()
	}
}
