package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

/*Scrivere un programma indovina_numero.go che chiede all'utente di indovinare un numero intero random x tra 1 e MAX,
(dove MAX è una costante definita nel programma) e ripete la richiesta fino a che l'utente indovina, oppure fino a un massimo di
MAX/2 tentativi.
Il programma stampa il numero di tentativi che sono stati necessari per indovinare oppure il messaggio "hai perso, il numero era x".
Se il numero digitato dall'utente è fuori dall'intervallo [1,MAX], il tentativo non viene considerato e viene visualizzato il
messaggio "fuori intervallo!", senza interrompere l'esecuzione.
Utilizzare la funzione rand.Intn del package "math/rand" per fissare il numero da indovinare.*/

func main() {
	var max, n, tentativo int
	rand.Seed(time.Now().Unix())
	fmt.Print("scegliere un max: ")
	fmt.Scan(&max)
	n = rand.Intn(max+1) + 1
	fmt.Printf("Indovina il numero! hai %d tentativi\n", max/2)
	for i := max / 2; i > 0; i-- {
		fmt.Scan(&tentativo)
		if tentativo == n {
			fmt.Println("Indovinato")
			os.Exit(0)
		} else if i == 1 {
			fmt.Println("Sbagliato! Sarà per la prossima volta")
			os.Exit(0)
		} else {
			fmt.Printf("Sbagliato! Ti rimangono %d tentativi\n", i-1)
		}
	}

}
