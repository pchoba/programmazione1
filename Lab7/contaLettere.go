package main

import (
	"bufio"
	"fmt"
	"os"
)

const LEN_ALFABETO = 26
const lett = "abcdefghijklmnopqrstuvwxyz"

func aggiornaConteggi(s string, contaMinu *[LEN_ALFABETO]int) {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(lett); j++ {
			if s[i] == lett[j] {
				contaMinu[j]++
			}
		}
	}
}

func main() {
	var contaMinu [LEN_ALFABETO]int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		aggiornaConteggi(scanner.Text(), &contaMinu)
	}
	for i := 0; i < len(lett); i++ {
		fmt.Printf("%c: %d\n", lett[i], contaMinu[i])
	}
}
