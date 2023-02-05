package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

/* Scrivere un programma polinomio.go che legge da standard input una riga che contiene dei numeri a, b, ....
Il programma calcola il valore in x del polinomio
a + bx + cx^2 + dx^3 ....
corrispondente alla sequenza dei numeri letti tranne l'ultimo, dove il valore di x è l'ultimo valore della sequenza.
Ad esempio,
3 2 0 7 5
corrisponde al polinomio 3 + 2x + 7x^3 da valutare per x = 5 */

func main() {
	var risultato int
	valori := []int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		valori_str := strings.Split(text, " ")
		for i := 0; i < len(valori_str); i++ {
			num, _ := strconv.Atoi(valori_str[i])
			valori = append(valori, num)
		}
	}
	risultato = valori[0]
	for i := 1; i < len(valori)-1; i++ {
		risultato += valori[i] * int(math.Pow(float64(valori[len(valori)-1]), float64(i)))
	}
	fmt.Println(risultato)
}
