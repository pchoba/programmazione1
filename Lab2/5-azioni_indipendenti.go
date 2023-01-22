package main

import "fmt"

func main() {
	/*
		Il programma leggere un numero da input. Se questo è divisibile
		per 3, stampa Fizz, se invece è divisibile per 5, stampa Buzz
	*/

	var num int
	fmt.Print("numero? ")
	fmt.Scan(&num)

	if num%3 == 0 {
		fmt.Println("Fizz")
	}
	if num%5 == 0 {
		fmt.Println("Buzz")
	}
}
