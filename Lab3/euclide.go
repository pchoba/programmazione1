package main

import "fmt"

/*Scrivere un programma euclide.go che legge da standard input due interi a e b, con a >= b, e calcola il MCD tra i due numeri
con l'algoritmo di Euclide.
Algoritmo di Euclide:
Dati due numeri naturali a e b:

controlla se b è zero.
se lo è, a è il MCD.
se non lo è, assegna ad r il resto della divisione a / b,
poni a = b e b = r e ripeti da 1.
*/

func main() {
	var a, b, r, mcd int
	fmt.Print("inserire due numeri (il primo deve essere maggiore del secondo); ")
	fmt.Scan(&a, &b)
	fmt.Printf("L'MCD tra %d e %d", a, b)
	for {
		if b == 0 {
			mcd = a
			break
		} else {
			r = a / b
			a = b
			b = r
		}
	}
	fmt.Printf(" è %d\n", mcd)
}
