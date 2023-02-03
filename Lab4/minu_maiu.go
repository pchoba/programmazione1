package main

import "fmt"

/* Scrivere un programma minu_maiu.go che legge da standard input una stringa e
stabilisce se la stringa contiene solo minuscole, solo maiuscole o sia minuscole che
maiuscole, quindi stampa un messaggio opportuno (es. "solo minuscole","solo minuscole", "sia minuscole che maiuscole"). */

func main() {
	var s string
	var min, mai int
	fmt.Print("Inserire una stringa: ")
	fmt.Scan(&s)
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			mai++
		} else {
			min++
		}
	}
	if mai != 0 && min == 0 {
		fmt.Println("solo maiuscole")
	} else if min != 0 && mai == 0 {
		fmt.Println("solo minuscole")
	} else {
		fmt.Println("sia minuscole che maiuscole")
	}
}
