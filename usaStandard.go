package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	const VOCALI = "aeiouAEIOU"
	const CIFRE = "0123456789"
	const s1 = "a"
	const s2 = "bogdan"
	var parola string
	//const s3 = "culo"
	var voc int
	var cont int
	var simo string

	/*
	   Il programma deve leggere da standard input una parola e utilizzare le funzioni dei pacchetti strings e strconv per svolgere i seguenti compiti:
	   Deve verificare

	   se la stringa letta contiene S2 come sottostringa (nel caso stampare un messaggio)
	   se la stringa letta contiene almeno una vocale (nel caso stampare un messaggio)

	   Deve poi determinare
	   3. quante occorrenze di S1 ha la stringa letta (e stampare un messaggio)
	   4. in che posizione si trova la prima vocale della stringa letta e in che posizione si trova l'ultima vocale della stringa letta
	   e stampare i risultati ottenuti

	   5. Deve stampare la stringa ottenuta concatenando S2 3 volte con se stessa.
	   6. Deve stampare la stringa ottenuta concatenando S1 5 volte con se stessa.

	   Deve infine estrarre le eventuali cifre contenute nella parola letta, concatenarle in una stringa s nello stesso ordine in cui le incontra e
	   7. stampare l'int n corrispondente a s usando fmt.Printf("%d\n", n)
	   8. stampare il float64 f corrispondente a "0.s" usando fmt.Printf("%f\n", f).
	*/
	fmt.Print("Inserire una parola: ")
	fmt.Scan(&parola)

	// 1. controllare una sottostringa
	if strings.Contains(parola, s2) {
		fmt.Println("s2 è una sottostringa")
	} else {
		fmt.Println("s2 non è una sottostringa")
	}
	// 2.controllare una vocale
	for i := 0; i < len(parola); i++ {
		if strings.Contains(VOCALI, string(parola[i])) == true {
			fmt.Println("La stringa contiene almeno una vocale")
			voc = 1
			break
		}
	}
	if voc == 0 {
		fmt.Println("La stringa non ha vocali")
	}

	// 3.controllare quante occorrenze ha s1 in parola

	for i := 0; i < len(parola); i++ {
		if strings.Contains(string(parola[i]), s1) == true {
			cont++
		}
	}
	fmt.Println("La sottostringa è contenuta", cont, "volte")
	// 4.controllo posizione prima e ultima vocale
	for i := 0; i < len(parola); i++ {
		if strings.Contains(VOCALI, string(parola[i])) {
			fmt.Println("La prima vocale è in", i+1, "posizione")
			break
		}
	}

	for i := len(parola) - 1; i >= 0; i-- {
		if strings.Contains(VOCALI, string(parola[i])) {
			fmt.Println("L'ultima vocale è in", i+1, "posizione")
			break
		}
	}
	// 5.stringa s2 concatenata con se stessa 3 volte
	fmt.Println(s2 + s2 + s2)
	// 6.stringa s1 concatenata con se stessa 5 volte
	fmt.Println(s1 + s1 + s1 + s1 + s1)

	for i := 0; i < len(parola); i++ {
		if unicode.IsNumber(rune(parola[i])) {
			simo += string(parola[i])
		}
	}
	// 7.stampare l'int corrispondente alla var simo con printf
	simoint, _ := strconv.Atoi(simo)
	fmt.Printf("%d\n", simoint)

	// 8.stampare il float64 corrispondente a 0.simoint

	simo = "0." + simo
	simofloat, _ := strconv.ParseFloat(simo, 64)
	fmt.Printf("%f\n", simofloat)

}
