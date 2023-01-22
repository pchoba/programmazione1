package main

import (
	"fmt"
)

func main() {
	var num_r, num int
	num_r = 8
	fmt.Print("Inserire un numero: ")
	fmt.Scan(&num)
	if num < 1 || num > 10 {
		fmt.Println("Valore non valido")
	} else if num == num_r {
		fmt.Println("Hai indovinato!")
	} else {
		fmt.Println("Non hai indovinato")
	}
}
