package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const VOCALI = "aeiouAEIOU"

func contaVocali(s string, vocali map[rune]int) {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(VOCALI); j++ {
			if s[i] == VOCALI[j] {
				vocali[rune(s[i])]++
			}
		}
	}
}
func main() {
	var riga []string
	vocali := make(map[rune]int)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		riga = strings.Split(text, " ")
		for i := 0; i < len(riga); i++ {
			contaVocali(riga[i], vocali)
		}
	}
	for i := 0; i < len(VOCALI); i++ {
		if vocali[rune(VOCALI[i])] != 0 {
			fmt.Println(string(VOCALI[i]), "presenti: ", vocali[rune(VOCALI[i])])
		}
	}
}
