package main

import "fmt"

func main() {
	/*
	 Questo programma controlla se la variabile voto è compresa
	 negli estremi possibili [0,30]
	*/

	var voto int

	fmt.Print("voto: ")
	fmt.Scan(&voto)

	if voto < 0 || voto > 30 {
		fmt.Println("voto non valido")
	}
}
