package main

import (
	"fmt"
	"strconv"
)

/* In un file giorni_mese.go definire una funzione

	giorniInMese(mese int) int

che, dato come parametro il numero corrispondente a un mese, restituisce il numero di giorni di quel mese
(28 per febbraio; 30 per aprile, giugno, settembre, novembre; 31 per G,M,M,L,A,O,D).

Assumiamo che il numero passato come parametro sia in [1,12], quindi non facciamo controlli.

Usiamo uno switch. Quanti casi conterrà (incluso il caso default, se lo si usa)? 3 casi, 12 casi o quanti?

Scrivere inoltre una funzione

	main()

per invocare e testare la funzione giorniInMese.

La funzione main deve leggere da standard input una stringa nel formato gg-mm-aaaa (vedi funzione Atoi del pacchetto strconv)
e stampare "il mese <x> ha <tot> giorni", dove x e tot sono numeri, usando la funzione giorniInMese per determinare tot.
Chiamare il programma giorni_mese.go e caricarlo su upload (dopo aver lanciato i test) */

func giorniInMese(mese int) int {
	switch mese {
	case 11, 4, 6, 9:
		return 30
	case 2:
		return 28
	default:
		return 31
	}
}

func main() {
	var s string
	fmt.Print("Inserire una data (gg-mm-aaaa): ")
	fmt.Scan(&s)
	mese := s[3:5]
	mese_num, _ := strconv.Atoi(mese)
	fmt.Printf("Il mese %d ha %d giorni\n", mese_num, giorniInMese(mese_num))
}
