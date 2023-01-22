package main

import "fmt"

func main() {
	/*
		Il programma controlla se un voto sia valido e se lo è
		restituisce una lettera dipendente dal voto
	*/

	var voto int
	fmt.Print("voto? ")
	fmt.Scan(&voto)

	if voto < 0 || voto > 100 {
		fmt.Println("voto non valido")
		return //interrompe l'esecuzione del programma
	}

	if voto >= 90 {
		fmt.Println("A")
	} else if voto >= 80 {
		fmt.Println("B")
	} else if voto >= 70 {
		fmt.Println("C")
	} else if voto >= 60 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}
}
