package main

import "fmt"

func main() {
	var n int
	fmt.Print("inserire un anno: ")
	fmt.Scan(&n)
	if (n >= 4 && n <= 1580) && n%4 == 0 {
		fmt.Println("bisestile")
	} else if n > 1580 {
		if n%100 != 0 && n%4 == 0 {
			fmt.Println("bisestile")
		} else if n%100 == 0 && n%400 == 0 {
			fmt.Println("bisestile")
		} else {
			fmt.Println("non bisestile")
		}
	} else {
		fmt.Println("non bisestile")
	}
}
