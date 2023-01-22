package main

import (
	"fmt"
	"strconv"
)

func main() {
	var orario string
	var min, ore int
	fmt.Println("Insere un orario (hh:mm): ")
	fmt.Scan(&orario)
	if string(orario[1]) == ":" {
		ore, _ = strconv.Atoi(string(orario[0]))
		min, _ = strconv.Atoi(string(orario[2:]))
	} else if string(orario[2]) == ":" {
		ore, _ = strconv.Atoi(string(orario[:2]))
		min, _ = strconv.Atoi(string(orario[3:]))
	}
	if ore >= 6 && ore <= 11 {
		fmt.Println("mattino")
	} else if (ore == 5 && min >= 30) && (ore == 12 && min < 30) {
		fmt.Println("mattino")
	} else {
		fmt.Println("non mattino")
	}
}
