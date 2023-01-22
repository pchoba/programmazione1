package main

import "fmt"

func main() {
	var gg, mm int
	fmt.Print("Inserire giorno e mese: ")
	fmt.Scan(&gg, &mm)
	if gg >= 1 && gg <= 31 {
		if mm >= 1 && mm <= 12 {
			fmt.Println("data valida")
		} else {
			fmt.Println("data non valida")
		}
	} else {
		fmt.Println("data non valida")
	}
}
