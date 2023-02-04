package main

import (
	"fmt"
	"strconv"
)

func punti(s string) int {
	var somma int
	for i := 0; i < len(s); i++ {
		cifra, err := strconv.Atoi(string(s[i]))
		if err == nil {
			if cifra == 3 {
				somma += 10
			} else if cifra == 2 || cifra == 4 || cifra == 5 || cifra == 6 || cifra == 7 {
				somma += 0
			} else {
				return -1
			}
		} else {
			if s[i] == 'J' {
				somma += 2
			} else if s[i] == 'Q' {
				somma += 3
			} else if s[i] == 'K' {
				somma += 4
			} else if s[i] == 'A' {
				somma += 11
			} else {
				return -1
			}
		}
	}
	return somma
}

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(punti(s))
}
