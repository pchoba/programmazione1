package main

import "fmt"

func main() {
	var num int
	fmt.Print("Inserire un numero: ")
	fmt.Scan(&num)
	if num%10 == 0 {
		if num > 0 {
			fmt.Println("positivo divisibile per 10")
		} else {
			fmt.Println("divisibile per 10")
		}
	}
	if num > 0 && num%10 != 0 {
		fmt.Println("positivo")
	}
}
