package main

import "fmt"

func operazioni(a, b int) (int, int, int) {
	return a + b, a * b, a - b
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(operazioni(a, b))
}
