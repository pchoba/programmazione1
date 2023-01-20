package main

import "fmt"

func SommaDivisoriPropri(n int) int {
	var somma int
	for i := 1; i < n; i++ {
		if n%i == 0 {
			somma += i
		}
	}
	if somma == 0 {
		return 0
	} else {
		return somma
	}
}

func eAbbondante(n int) bool {
	if n < SommaDivisoriPropri(n) {
		return true
	} else {
		return false
	}
}

func numeriAbbondanti(limite int) {
	if limite <= 0 {
		fmt.Println("La somma inserita non è positiva")
	} else {
		for i := 0; i < limite; i++ {
			if eAbbondante(i) {
				fmt.Print(i, " ")
			}
		}
	}
}
