package main

import "fmt"

func ePrimo(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func NumeriPrimi(limite int) {
	for n := 2; n < limite; n++ {
		if ePrimo(n) {
			fmt.Print(n, " ")
		}
	}
}

func main() {
	var limite int
	fmt.Print("Inserire un numero limite: ")
	fmt.Scanf("%d\n", &limite)
	NumeriPrimi(limite)
	fmt.Print("\n")
}
