package main

import (
	"fmt"
	"strconv"
)

func hasUpper(s string) bool {
	/* La funzione riceve una stringa come parametro e restituisce
	true se la stringa contiene almeno una lettera latina maiuscola (tra 'A' e 'Z'), false altrimenti. */
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			return true
		}
	}
	return false
}

func firstUpper(s string) int {
	/* La funzione riceve una stringa come parametro e restituisce la
	posizione della prima lettera latina maiuscola (tra 'A' e 'Z'), se la stringa ne contiene almeno una, -1 altrimenti. */
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			return i
		}
	}
	return -1
}

func lastUpper(s string) int {
	/* Scrivere una funzione lastUpper(s string) int. La funzione riceve una stringa come parametro e
	restituisce la posizione dell'ultima lettera latina maiuscola (tra 'A' e 'Z'), se la stringa ne contiene
	almeno una, -1 altrimenti. */
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] >= 'A' && s[i] <= 'Z' {
			return i
		}
	}
	return -1
}

func countDigitsLettersOthers(s string) (int, int, int) {
	/* La funzione riceve una stringa come parametro, conta quante cifre, quante lettere e quanti caratteri
	che non sono né cifre né lettere contiene e restituisce i tre risultati in questo ordine. */
	var cifre, letters, others int
	for i := 0; i < len(s); i++ {
		_, err := strconv.Atoi(string(s[i]))
		if err == nil {
			cifre++
			continue
		}
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] >= 'Z') {
			letters++
			continue
		}
		others++
	}
	return cifre, letters, others
}

func isPalindrome(s string) bool {
	/*La funzione riceve una stringa come parametro e restituisce true se la stringa è palindroma,
	e false altrimenti. Una parola è palindroma se può essere letta  sia da sinistra a destra che da
	destra a sinistra. Ad esempio "osso" e "ingegni" sono palindrome, ma anche "" e "t". */
	var count int
	for i := 0; i < len(s); i++ {
		if s[i] == s[len(s)-1-i] {
			count++
		}
	}
	if count == len(s) {
		return true
	}
	return false
}

func main() {
	var s string
	fmt.Scan(&s)
	if hasUpper(s) {
		fmt.Println("La stringa contiene almeno una maiuscola")
		fmt.Printf("La prima maiuscola è in posizione %d\n", firstUpper(s))
		fmt.Printf("L'ultima maiuscola è in posizione %d\n", lastUpper(s))
	}
	cifre, lettere, others := countDigitsLettersOthers(s)
	fmt.Printf("La stringa contiene %d cifre, %d lettere e %d altri caratteri\n", cifre, lettere, others)
	if isPalindrome(s) {
		fmt.Println("La stringa è palindroma")
	}
}
