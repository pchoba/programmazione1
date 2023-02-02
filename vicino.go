package main

import (
	"fmt"
	"math"
)

/*
Scrivere un programma vicino.go che, data una serie di 5 valori interi tra 0 e 20, trova il valore piu' vicino a
TARGET, dove TARGET è una costante definita nel programma.
*/

func main() {
	var target, a, b, c, d, e int
	var vicino float64
	vicino = math.Pow(2.0, 32.0)
	fmt.Print("Inserire un target: ")
	fmt.Scan(&target)
	fmt.Print("Inserire cinque valori: ")
	fmt.Scan(&a, &b, &c, &d, &e)
	if math.Abs(float64(target-a)) < vicino {
		vicino = float64(a)
	}
	if math.Abs(float64(target-b)) < vicino {
		vicino = float64(b)
	}
	if math.Abs(float64(target-c)) < vicino {
		vicino = float64(c)
	}
	if math.Abs(float64(target-d)) < vicino {
		vicino = float64(d)
	}
	if math.Abs(float64(target-e)) < vicino {
		vicino = float64(e)
	}
	fmt.Printf("il valore più vicino è %d\n", int(vicino))

}
