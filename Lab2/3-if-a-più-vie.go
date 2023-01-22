package main

import "fmt"

func main() {
	/*
	 Dato un numero inserito dall'utente, il programma controlla
	 se il numero sia negativo, nullo o positivo
	*/

	var num int

	fmt.Print("un int: ")
	fmt.Scan(&num)

	if num > 0 {
		fmt.Println("positivo")
	} else if num == 0 {
		fmt.Println("nullo")
	} else { // < 0
		fmt.Println("negativo")
	}
}
