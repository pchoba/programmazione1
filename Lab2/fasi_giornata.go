package main

import "fmt"

func main() {
	var ora int
	fmt.Print("ora: ")
	fmt.Scan(&ora)
	if ora < 0 || ora > 23 {
		fmt.Println("orario non valido")
	} else {
		if ora >= 0 && ora <= 6 {
			fmt.Println("ora")
		} else if ora >= 7 && ora <= 13 {
			fmt.Println("mattino")
		} else if ora >= 14 && ora <= 18 {
			fmt.Println("pomeriggio")
		} else if ora >= 19 && ora <= 23 {
			fmt.Println("sera")
		}
	}
}
