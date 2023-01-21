package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("val1 e val2 (int): ")
	fmt.Scan(&a, &b)
	fmt.Println("Prima dello scambio:", a, b)
	a, b = b, a
	fmt.Println("Dopo lo scambio:", a, b)

}
