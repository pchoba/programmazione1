package main

import (
	"fmt"
	"os"
)

func eBisestile(n int) bool {
	if (n >= 4 && n <= 1580) && n%4 == 0 {
		return true
	} else if n > 1580 {
		if n%100 != 0 && n%4 == 0 {
			return true
		} else if n%100 == 0 && n%400 == 0 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func main() {
	var gg, mm, aaaa, temp int
	fmt.Print("Inserire giorno, mese e anno: ")
	fmt.Scan(&gg, &mm, &aaaa)
	if eBisestile(aaaa) {
		temp = 29
	} else {
		temp = 28
	}
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
		if gg >= 1 && gg <= temp {
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
