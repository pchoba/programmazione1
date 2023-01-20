package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func contaCifre(s string, numCifre *[10]int) (haCifre bool) {
	for i := 0; i < len(s); i++ {
		_, err := strconv.Atoi(string(s[i]))
		if err != nil {
			return haCifre == true
		} else {
			continue
		}
	}
}

func main() {
	var s string
	var n, one, two, three, four, five, six, seven, eight, nine, zero, temp int
	for {
		fmt.Scan(&s)
		if s == "stop" {
			break
		} else {
			for i := 0; i < len(s); i++ {
				if unicode.IsNumber(rune(s[i])) {
					n++
					temp, _ = strconv.Atoi(string(s[i]))
					switch temp {
					case 0:
						zero++
					case 1:
						one++
					case 2:
						two++
					case 3:
						three++
					case 4:
						four++
					case 5:
						five++
					case 6:
						six++
					case 7:
						seven++
					case 8:
						eight++
					case 9:
						nine++
					}
				}
			}
		}
	}
	if n != 0 {
		fmt.Printf("In %d stringhe c'era almeno una cifra\n", n)
	} else {
		fmt.Println("Nella sequenza di stringhe non c'erano cifre")
	}
	if zero != 0 {
		fmt.Printf("Nella sequenza di stringhe c'erano %d zero\n", zero)
	}
	if one != 0 {
		fmt.Printf("Nella sequenza di stringhe c'erano %d uno\n", one)
	}
	if two != 0 {
		fmt.Printf("Nella sequenza di stringhe c'erano %d due\n", two)
	}
	if three != 0 {
		fmt.Printf("Nella sequenza di stringhe c'erano %d tre\n", three)
	}
	if four != 0 {
		fmt.Printf("Nella sequenza di stringhe c'erano %d quattro\n", four)
	}
	if five != 0 {
		fmt.Printf("Nella sequenza di stringhe c'erano %d cinque\n", five)
	}
	if six != 0 {
		fmt.Printf("Nella sequenza di stringhe c'erano %d sei\n", six)
	}
	if seven != 0 {
		fmt.Printf("Nella sequenza di stringhe c'erano %d sette\n", seven)
	}
	if eight != 0 {
		fmt.Printf("Nella sequenza di stringhe c'erano %d otto\n", eight)
	}
	if nine != 0 {
		fmt.Printf("Nella sequenza di stringhe c'erano %d nove\n", nine)
	}
}
