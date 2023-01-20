package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

type Carta struct {
	Valore string
	Seme   string
}

const semi = 4
const valori = 13
const carte = 52

var num int
var carta_gen Carta
var mazzo []Carta

// funzione carta
func carta(n int) (Carta, bool) {
	carta_gen = Carta{"0", "0"}
	if n >= 0 && n <= 12 {
		carta_gen.Seme = "♡"
	} else if n >= 13 && n <= 25 {
		n -= 13
		carta_gen.Seme = "♢"
	} else if n >= 26 && n <= 38 {
		n -= 26
		carta_gen.Seme = "♤"
	} else if n >= 39 && n <= 51 {
		n -= 39
		carta_gen.Seme = "♧"
	} else {
		return carta_gen, false
	}
	switch n {
	case 0:
		carta_gen.Valore = "A"
	case 1:
		carta_gen.Valore = "2"
	case 2:
		carta_gen.Valore = "3"
	case 3:
		carta_gen.Valore = "4"
	case 4:
		carta_gen.Valore = "5"
	case 5:
		carta_gen.Valore = "6"
	case 6:
		carta_gen.Valore = "7"
	case 7:
		carta_gen.Valore = "8"
	case 8:
		carta_gen.Valore = "9"
	case 9:
		carta_gen.Valore = "10"
	case 10:
		carta_gen.Valore = "J"
	case 11:
		carta_gen.Valore = "Q"
	case 12:
		carta_gen.Valore = "K"
	}
	return carta_gen, true

}

// funzione estrai carta
func estraiCarta(mazzo []Carta) Carta {
	rand.Seed(time.Now().UTC().UnixNano())
	random_num := rand.Intn(52)
	carta_gen = mazzo[random_num]
	mazzo = append(mazzo[:random_num], mazzo[random_num+1:]...)
	return carta_gen

}

/* funzione dai4Carte() che restituisce un array di 4 carte estratte con estraiCarta.
func dai4Carte() [4]Carta {
	for i := 0; i < 4; i++ {
		mano[i] = estraiCarta()
	}
	return mano
}
*/

func getvaloreBJ(carta_gen Carta) int {
	var valorebj int
	switch carta_gen.Valore {
	case "A":
		valorebj = 11
	case "2":
		valorebj = 2
	case "3":
		valorebj = 3
	case "4":
		valorebj = 4
	case "5":
		valorebj = 5
	case "6":
		valorebj = 6
	case "7":
		valorebj = 7
	case "8":
		valorebj = 8
	case "9":
		valorebj = 9
	case "10":
		valorebj = 10
	case "J", "Q", "K":
		valorebj = 10
	}
	return valorebj
}

func MazzodaPoker(mazzo []Carta) []Carta {
	for i := 0; i < 52; i++ {
		if i >= 0 && i <= 12 {
			carta_gen.Seme = "♡"
		} else if i >= 13 && i <= 25 {
			carta_gen.Seme = "♢"
		} else if i >= 26 && i <= 38 {
			carta_gen.Seme = "♤"
		} else if i >= 39 && i <= 51 {
			carta_gen.Seme = "♧"
		}
		switch i {
		case 0, 13, 26, 39:
			carta_gen.Valore = "A"
		case 1, 14, 27, 40:
			carta_gen.Valore = "2"
		case 2, 15, 28, 41:
			carta_gen.Valore = "3"
		case 3, 16, 29, 42:
			carta_gen.Valore = "4"
		case 4, 17, 30, 43:
			carta_gen.Valore = "5"
		case 5, 18, 31, 44:
			carta_gen.Valore = "6"
		case 6, 19, 32, 45:
			carta_gen.Valore = "7"
		case 7, 20, 33, 46:
			carta_gen.Valore = "8"
		case 8, 21, 34, 47:
			carta_gen.Valore = "9"
		case 9, 22, 35, 48:
			carta_gen.Valore = "10"
		case 10, 23, 36, 49:
			carta_gen.Valore = "J"
		case 11, 24, 37, 50:
			carta_gen.Valore = "Q"
		case 12, 25, 38, 51:
			carta_gen.Valore = "K"
		}
		mazzo = append(mazzo, carta_gen)
	}
	return mazzo
}

func mischia(mazzo []Carta) []Carta {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 52; i++ {
		mazzo[i], mazzo[rand.Intn(52)] = mazzo[rand.Intn(52)], mazzo[i]
	}
	return mazzo
}
func cancella() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func main() {
	var mano_giocatore, mano_dealer []Carta
	var decisione string
	var mazzo_gen []Carta
	var active, decision, card_choice bool
	active = true
	decision = true
	card_choice = true
	var punti_giocatore, punti_dealer int
	for active {
		cancella()
		mazzo_gen = nil
		mazzo_gen = MazzodaPoker(mazzo_gen)
		mano_giocatore = nil
		mano_dealer = nil
		rand.Seed(time.Now().UTC().UnixNano())
		mischia(mazzo_gen)
		fmt.Println("Benvenuto al tavolo!")
		card_choice = true
		decision = true
		time.Sleep(time.Duration(1) * time.Second)
		rand.Seed(time.Now().UTC().UnixNano())
		fmt.Println("Queste sono le tue carte: ")
		// do 2 carte al giocatore
		mano_giocatore = append(mano_giocatore, estraiCarta(mazzo_gen))
		mano_giocatore = append(mano_giocatore, estraiCarta(mazzo_gen))
		for i := 0; i < 2; i++ {
			fmt.Print(mano_giocatore[i], " ")
			fmt.Println(getvaloreBJ(mano_giocatore[i]))
		}
		punti_giocatore = getvaloreBJ(mano_giocatore[0]) + getvaloreBJ(mano_giocatore[1])
		fmt.Println("Punteggio attuale giocatore: ", punti_giocatore)
		fmt.Println("-------------------------------")
		time.Sleep(time.Duration(1) * time.Second)
		mano_dealer = append(mano_dealer, estraiCarta(mazzo_gen))
		mano_dealer = append(mano_dealer, estraiCarta(mazzo_gen))
		punti_dealer = getvaloreBJ(mano_dealer[0]) + getvaloreBJ(mano_dealer[1])
		fmt.Println("Questa è la carta del dealer: ")
		fmt.Print(mano_dealer[0], " ")
		fmt.Println(getvaloreBJ(mano_dealer[0]))
		fmt.Println("Punteggio attuale dealer: ", getvaloreBJ(mano_dealer[0]))
		fmt.Println("-------------------------------")
		if punti_giocatore == 21 && punti_dealer != 21 {
			fmt.Println("Blackjack! Hai vinto!")
			for decision {
				fmt.Print("Vuoi fare un'altra partita? (y/n): ")
				fmt.Scan(&decisione)
				if decisione == "y" {
					cancella()
					decision = false
					continue
				} else if decisione == "n" {
					decision = false
					active = false
				} else {
					fmt.Println("Uh-oh: hai scelto male")
					continue
				}
			}
		} else if (mano_dealer[0].Valore == "J" || mano_dealer[0].Valore == "Q" || mano_dealer[0].Valore == "K") && (mano_dealer[1].Valore == "A") {
			fmt.Print(mano_dealer[0], " ", mano_dealer[1])
			fmt.Println("Il dealer ha fatto 21 con ", mano_dealer[0], " e", mano_dealer[1], ": hai perso!")
			for decision == true {
				fmt.Println("Vuoi fare un'altra partita? (y/n): ")
				fmt.Scan(&decisione)
				if decisione == "y" {
					cancella()
					decision = false
				} else if decisione == "n" {
					decision = false
					active = false
				} else {
					fmt.Println("Uh-oh: hai scelto male")
					continue
				}
			}
			continue

		} else if (mano_dealer[0].Valore == "A") && (mano_dealer[1].Valore == "J" || mano_dealer[1].Valore == "Q" || mano_dealer[1].Valore == "K") {
			//fmt.Print(mano_dealer[0], " ", mano_dealer[1])
			fmt.Println("Il dealer ha fatto 21 con ", mano_dealer[0], " e", mano_dealer[1], ": hai perso!")
			for decision {
				fmt.Println("Vuoi fare un'altra partita? (y/n): ")
				fmt.Scan(&decisione)
				if decisione == "y" {
					cancella()
					decision = false
					continue
				} else if decisione == "n" {
					decision = false
					active = false
				} else {
					fmt.Println("Uh-oh: hai scelto male")
					continue
				}
			}
		} else if punti_dealer == 21 && punti_giocatore == 21 {
			fmt.Println("Pareggio! Entrambi avete fatto blackjack")
			for decision {
				fmt.Print("Vuoi fare un'altra partita? (y/n): ")
				fmt.Scan(&decisione)
				if decisione == "y" {
					cancella()
					decision = false
					continue
				} else if decisione == "n" {
					decision = false
					active = false
				} else {
					fmt.Println("Uh-oh: hai scelto male")
					continue
				}
			}
		}
		for card_choice {
			fmt.Print("Vuoi pescare una carta o vuoi fermarti? (carta/stop): ")
			fmt.Scan(&decisione)
			if decisione == "stop" {
				cancella()
				card_choice = false
				fmt.Println("Queste sono le tue carte: ")
				for i := 0; i < len(mano_giocatore); i++ {
					fmt.Print(mano_giocatore[i], " ")
					fmt.Println(getvaloreBJ(mano_giocatore[i]))
				}
				fmt.Print("Punteggio attuale giocatore: ")
				fmt.Println(punti_giocatore)

				fmt.Println("-------------------------------")

				punti_dealer = getvaloreBJ(mano_dealer[0]) + getvaloreBJ(mano_dealer[1])
				fmt.Println("Ecco le carte del dealer: ")
				fmt.Println(mano_dealer[0], " ", getvaloreBJ(mano_dealer[0]))
				time.Sleep(time.Duration(1) * time.Second)
				fmt.Println(mano_dealer[1], " ", getvaloreBJ(mano_dealer[1]))
				for punti_dealer < 17 {
					mano_dealer = append(mano_dealer, estraiCarta(mazzo_gen))
					punti_dealer += getvaloreBJ(mano_dealer[len(mano_dealer)-1])
					time.Sleep(time.Duration(1) * time.Second)
					fmt.Println(mano_dealer[len(mano_dealer)-1], " ", getvaloreBJ(mano_dealer[len(mano_dealer)-1]))
				}
				fmt.Println("-------------------------------")
				if punti_dealer > 21 {
					fmt.Println("Il dealer ha superato 21, hai vinto!")
					for decision {
						fmt.Print("Vuoi fare un'altra partita? (y/n): ")
						fmt.Scan(&decisione)
						if decisione == "y" {
							cancella()
							decision = false
						} else if decisione == "n" {
							decision = false
							active = false
						} else {
							fmt.Println("Uh-oh: hai scelto male")
							continue
						}
					}

				} else if punti_dealer == punti_giocatore {
					fmt.Println("Pareggio!")
					for decision {
						fmt.Print("Vuoi fare un'altra partita? (y/n): ")
						fmt.Scan(&decisione)
						if decisione == "y" {
							cancella()
							decision = false
						} else if decisione == "n" {
							decision = false
							active = false
						} else {
							fmt.Println("Uh-oh: hai scelto male")
							continue
						}
					}
				} else if punti_dealer > punti_giocatore && punti_dealer <= 21 {
					fmt.Println("Il dealer ha fatto più punti di te, hai perso!")
					for decision {
						fmt.Print("Vuoi fare un'altra partita? (y/n): ")
						fmt.Scan(&decisione)
						if decisione == "y" {
							cancella()
							decision = false
						} else if decisione == "n" {
							decision = false
							active = false
						} else {
							fmt.Println("Uh-oh: hai scelto male")
							continue
						}
					}
				} else if punti_giocatore > punti_dealer && punti_giocatore <= 21 {
					fmt.Println("Hai fatto più punti del dealer, hai vinto!")
					for decision {
						fmt.Print("Vuoi fare un'altra partita? (y/n): ")
						fmt.Scan(&decisione)
						if decisione == "y" {
							cancella()
							decision = false
						} else if decisione == "n" {
							decision = false
							active = false
						} else {
							fmt.Println("Uh-oh: hai scelto male")
							continue
						}
					}
				} else {
					fmt.Println("Errore")
					active = false
					decision = false
					break
					//STACCAH STACCAH
				}

			} else if decisione == "carta" {
				cancella()
				fmt.Println("Ecco le tue carte: ")
				for i := 0; i < len(mano_giocatore); i++ {
					fmt.Println(mano_giocatore[i], " ", getvaloreBJ(mano_giocatore[i]))
				}
				mano_giocatore = append(mano_giocatore, estraiCarta(mazzo_gen))
				time.Sleep(time.Duration(1) * time.Second)
				fmt.Println(mano_giocatore[len(mano_giocatore)-1], " ", getvaloreBJ(mano_giocatore[len(mano_giocatore)-1]))
				fmt.Print("Punteggio attuale giocatore: ")
				punti_giocatore += getvaloreBJ(mano_giocatore[len(mano_giocatore)-1])
				fmt.Println(punti_giocatore)
				fmt.Println("-------------------------------")
				time.Sleep(time.Duration(1) * time.Second)
				fmt.Println("Ecco la carte del dealer: ")
				fmt.Println(mano_dealer[0], " ", getvaloreBJ(mano_dealer[0]))
				fmt.Println("Punteggio attuale dealer: ", getvaloreBJ(mano_dealer[0]))
				fmt.Println("-------------------------------")
				if punti_giocatore == 21 {
					card_choice = false
					fmt.Println("Blackjack! Hai vinto!")
					for decision {
						fmt.Print("Vuoi fare un'altra partita? (y/n): ")
						fmt.Scan(&decisione)
						if decisione == "y" {
							cancella()
							decision = false
						} else if decisione == "n" {
							decision = false
							active = false
						} else {
							fmt.Println("Uh-oh: hai scelto male")
							continue
						}
					}
				} else if punti_giocatore > 21 {
					card_choice = false
					fmt.Println("Hai superato 21, hai perso!")
					for decision {
						fmt.Print("Vuoi fare un'altra partita? (y/n): ")
						fmt.Scan(&decisione)
						if decisione == "y" {
							cancella()
							decision = false
						} else if decisione == "n" {
							decision = false
							active = false
						} else {
							fmt.Println("Uh-oh: hai scelto male")
							continue
						}
					}
				} else {
					continue
				}
			} else {
				fmt.Println("Devi scegliere una delle due opzioni (carta/stop)")
				fmt.Println("")
			}
		}
	}
}
