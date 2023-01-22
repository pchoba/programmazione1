package main

import "fmt"

func main() {
	var reddito int
	var coniug string
	var coniugato bool
	fmt.Print("Inserire il reddito e se si è coniugati (y/n): ")
	fmt.Scan(&reddito, &coniug)
	if coniug == "y" {
		coniugato = true
	} else if coniug == "n" {
		coniugato = false
	}
	if ((reddito >= 0 && reddito <= 32000) && coniugato == false) || ((reddito >= 0 && reddito <= 64000) && coniugato == true) {
		fmt.Println("tasse: 10%")
	} else if (reddito > 32000 && coniugato == false) || (reddito > 64000 && coniugato == true) {
		fmt.Println("tasse: 24%")
	}
}
