package main

import (
	"fmt"
	"os"
)

/* Scrivere un programma giornoSettimana.go, dotato di un array giorniDellaSettimana
[7]string{"lun", "mart", "merc", "giov", "ven", "sab", "dom"}
che legge da linea di comando il nome del giorno della settimana del 1° gennaio di un dato anno.
Se il giorno non è fra "lun", "mart", "merc", "giov", "ven", "sab", "dom", il programma avvisa e termina.
Altrimenti poi accetta dall'utente dei numeri interi (tra 1 e 365), corrispondenti a giorni di quell'anno,
e per ciascuno stampa il giorno della settimana corrispondente. Il programma termina quando l'utente digita -1. */

func main() {
	var n int
	grn_sttm_ord := []string{}
	giorniDellaSettimana := [7]string{"lun", "mart", "merc", "giov", "ven", "sab", "dom"}
	if len(os.Args) == 1 {
		os.Exit(1)
	}
	// os.Args[1] è il primo giorno della settimana
	if os.Args[1] != "lun" && os.Args[1] != "mart" && os.Args[1] != "merc" && os.Args[1] != "giov" && os.Args[1] != "ven" && os.Args[1] != "sab" && os.Args[1] != "dom" {
		os.Exit(1)
	} else {
		for i := 0; i < len(giorniDellaSettimana); i++ {
			if giorniDellaSettimana[i] == os.Args[1] {
				grn_sttm_ord = append(giorniDellaSettimana[i:], giorniDellaSettimana[:i+1]...)
			}
		}
		for {
			fmt.Scan(&n)
			if n == -1 {
				break
			}
			for n > 7 {
				n -= 7
			}
			switch n {
			case 1:
				fmt.Println(grn_sttm_ord[0])
			case 2:
				fmt.Println(grn_sttm_ord[1])
			case 3:
				fmt.Println(grn_sttm_ord[2])
			case 4:
				fmt.Println(grn_sttm_ord[3])
			case 5:
				fmt.Println(grn_sttm_ord[4])
			case 6:
				fmt.Println(grn_sttm_ord[5])
			case 7:
				fmt.Println(grn_sttm_ord[6])
			}
		}
	}
}
