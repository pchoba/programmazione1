package main

import "fmt"

func main() {
	var eta int
	var studente string
	var temp bool
	fmt.Print("età: ")
	fmt.Scan(&eta)
	fmt.Print("studente? (t/f): ")
	fmt.Scan(&studente)
	if studente == "f" {
		temp = false
	} else if studente == "t" {
		temp = true
	}
	if eta >= 0 && eta < 9 {
		fmt.Println("ingresso gratis")
	} else if eta >= 9 && eta < 18 {
		fmt.Println("ingresso 5 euro")
	} else if eta >= 10 && eta < 26 && temp == false {
		fmt.Println("ingresso 7.5 euro")
	} else if eta >= 26 && eta < 65 && temp == false {
		fmt.Println("ingresso 10 euro")
	} else if eta >= 65 && temp == false {
		fmt.Println("ingresso 7.5 euro")
	} else if eta >= 18 && temp == true {
		fmt.Println("ingresso 5 euro")
	}
}
