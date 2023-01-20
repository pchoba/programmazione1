package main

import "fmt"

func main() {
	var n, i int
	fmt.Print("Quante volte, Padrone?")
	fmt.Scan(&n)
	for i = 0; i < n; i++ {
		fmt.Println("Sono il tuo schiavo")
	}
}
