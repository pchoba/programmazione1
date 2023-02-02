package main

import "fmt"

/* Scrivere un programma capitale.go che legge da stardard input tre valori float64 che rappresentano un capitale
iniziale (es. 1000.50 euro), un tasso di interesse annuale (es. 1.3%) e un obiettivo (es. 2000 euro), e calcola
quanti anni occorrono per arrivare a (o superare) l'obiettivo. */

func main() {
	var cap_iniz, tas_int_ann, obb float64
	var anni int
	fmt.Print("Capitale iniziale (€): ")
	fmt.Scan(&cap_iniz)
	fmt.Print("Tasso di interesse annuo (%): ")
	fmt.Scan(&tas_int_ann)
	fmt.Print("Obiettivo finale (€): ")
	fmt.Scan(&obb)
	for cap_iniz < obb {
		cap_iniz += cap_iniz * (tas_int_ann / 100)
		anni++
	}
	fmt.Printf("Per raggiungere l'obiettivo ci vorranno almeno %d anni\n", anni)
}
