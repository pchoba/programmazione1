package main

import "fmt"

func main() {
	var n, max int
	for i := 0; i < 10; i++ {
		fmt.Print("Inserire numero: ")
		fmt.Scan(&n)
		if n > max {
			max = n
		}
	}
	fmt.Println("Il numero max è", max)
}
