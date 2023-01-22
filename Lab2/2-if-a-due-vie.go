package main

import "fmt"

func main() {
	/*
	 Questo programma controlla se un numero inserito è pari
	 o dispari
	*/

	var n int
	fmt.Print("numero: ")
	fmt.Scan(&n)

	if n%2 == 0 {
		fmt.Println(n, "è pari")
	} else {
		fmt.Println(n, "è dispari")
	}
}
