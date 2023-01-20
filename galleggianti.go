package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const LETTERE = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const NUMERI = "123456789"

func main() {
	var r, c, ltrc, stodiventandopazzo int //ltrc sta per letter counter
	var ltrs string                        // ltrs sta per letter string
	matrice := make(map[string]string)
	fmt.Print("Inserire numero righe della matrice: ")
	fmt.Scan(&r)
	fmt.Print("Inserire il numero di colonne della matrice: ")
	fmt.Scan(&c)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ltrs = string(LETTERE[ltrc])
		text := scanner.Text()
		text_slice := strings.Split(text, " ")
		for i := 0; i < len(text_slice); i++ {
			matrice[ltrs+string(NUMERI[i])] = text_slice[i]
		}
		ltrc++
	}
	//fmt.Println(matrice)
	for {
		stodiventandopazzo = 0
		for i := r; i > 1; i-- {
			for j := 0; j < c; j++ {
				if matrice[string(LETTERE[i-1])+string(NUMERI[j])] == "*" && matrice[string(LETTERE[i-2])+string(NUMERI[j])] != "*" {
					matrice[string(LETTERE[i-1])+string(NUMERI[j])], matrice[string(LETTERE[i-2])+string(NUMERI[j])] = matrice[string(LETTERE[i-2])+string(NUMERI[j])], matrice[string(LETTERE[i-1])+string(NUMERI[j])]
					stodiventandopazzo++
				}
			}
		}
		if stodiventandopazzo == 0 {
			break
		}
	}
	//fmt.Println(matrice)

	fmt.Println("")
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			fmt.Print(matrice[string(LETTERE[i])+string(NUMERI[j])], " ")
		}
		fmt.Println("")
	}
}
