package main

import (
	"fmt"
	"os"
)

func main() {
	var gg, mm int
	fmt.Print("Inserire giorno e mese: ")
	fmt.Scan(&gg, &mm)
	if mm < 1 || mm > 12 {
		fmt.Println("data non valida")
		os.Exit(1)
	}
	switch mm {
	case 1, 3, 5, 7, 8, 10, 12:
		if gg >= 1 && gg <= 31 {
			fmt.Println("data valida")
		} else {
			fmt.Println("data non valida")
		}
	case 2:
		if gg >= 1 && gg <= 28 {
			fmt.Println("data valida")
		} else {
			fmt.Println("data non valida")
		}
	default:
		if gg >= 1 && gg <= 30 {
			fmt.Println("data valida")
		} else {
			fmt.Println("data non valida")
		}

	}
}
