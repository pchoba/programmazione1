package main

import "fmt"

/*
	Problema. Scrivere un programma Go conversione_secondi.go che converta un numero dato di secondi

(fornito dall’utente) in giorni, ore, minuti, secondi senza mai usare l’operazione di sottrazione.
*/
func main() {
	var sc int
	fmt.Print("Inserire un numero di secondi: ")
	fmt.Scan(&sc)
	sc_giorni := sc / (3600 * 24)
	sc_ore := (sc % (3600 * 24)) / 3600
	sc_minuti := ((sc % (3600 * 24)) % 3600) / 60
	sc_secondi := (((sc % (3600 * 24)) % 3600) % 60)
	fmt.Printf("%d:%d:%d:%d\n", sc_giorni, sc_ore, sc_minuti, sc_secondi)
}
