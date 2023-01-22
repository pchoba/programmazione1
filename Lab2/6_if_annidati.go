package main

import "fmt"

func main() {
	/*
	 Inserire un numero che rappresenti una figura, e poi un'operazione geometrica sempre rappresentata
	 da un numero. Successivamente stampa la formula dell'operazione scelta per la figura scelta
	*/
	const (
		RETTANGOLO = 1
		TRIANGOLO  = 2
		PERIMETRO  = 1
		AREA       = 2
	)
	var figura, operazione int

	fmt.Print("rettangolo (1) o triangolo (2)? ")
	fmt.Scan(&figura)
	fmt.Print("perimetro (1) o area (2)? ")
	fmt.Scan(&operazione)

	if figura == RETTANGOLO {
		if operazione == PERIMETRO {
			fmt.Println("formula = (base + altezza) * 2")
		} else { //operazione == AREA
			fmt.Println("formula = base * altezza")
		}
	} else { //figura == TRIANGOLO
		if operazione == PERIMETRO {
			fmt.Println("formula = lato1 + lato2 + lato3 ")
		} else { //operazione == AREA
			fmt.Println("formula = (base * altezza)/2")
		}
	}
}
