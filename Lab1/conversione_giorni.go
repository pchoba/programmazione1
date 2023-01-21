package main

import "fmt"

func main() {
	var gg int
	fmt.Print("Inserire un numero di giorni: ")
	fmt.Scan(&gg)
	gg1 := gg / 365
	gg2 := (gg % 365) / 7
	gg3 := ((gg % 365) % 7)
	fmt.Println(gg, "giorni =", gg1, "anni +", gg2, "settimane +", gg3, "giorni")
}
