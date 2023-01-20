package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	type Capoluogo struct {
		Nome        string
		Sigla       string
		Regione     string
		Popolazione int
		Superficie  int
		Densita     int
		Altitudine  int
	}
	var capolouogo_gen Capoluogo
	scanner := bufio.NewScanner(os.Stdin)
	var dati []string
	var capoluoghi []Capoluogo
	for scanner.Scan() {
		text := scanner.Text()
		dati = strings.Split(text, ",")
		capolouogo_gen.Nome = dati[0]
		capolouogo_gen.Sigla = dati[1]
		capolouogo_gen.Regione = dati[2]
		capolouogo_gen.Popolazione, _ = strconv.Atoi(string(dati[3]))
		capolouogo_gen.Superficie, _ = strconv.Atoi(string(dati[4]))
		capolouogo_gen.Densita, _ = strconv.Atoi(string(dati[5]))
		capolouogo_gen.Altitudine, _ = strconv.Atoi(string(dati[6]))
		capoluoghi = append(capoluoghi, capolouogo_gen)
	}
	capoluoghi = capoluoghi[1:]
	for i := 0; i < len(capoluoghi); i++ {
		if capoluoghi[i].Nome == "MILANO" {
			fmt.Println(capoluoghi[i])
		}
	}
}
