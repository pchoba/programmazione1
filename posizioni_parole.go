package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var parole_per_riga, temp []string
	var posiz_parole map[string]string
	posiz_parole = make(map[string]string)
	var strI string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		temp = strings.Split(text, " ")
		parole_per_riga = append(parole_per_riga, temp...)
	}
	//fmt.Println(parole_per_riga)
	for i := 0; i < len(parole_per_riga); i++ {
		strI = strconv.Itoa(i)
		posiz_parole[parole_per_riga[i]] += strI
		posiz_parole[parole_per_riga[i]] += " "
	}
	fmt.Println(posiz_parole)
}
