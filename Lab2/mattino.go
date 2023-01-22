package main

import (
	"fmt"
	"strconv"
)

func main() {
	var ora string
	var oraint int
	fmt.Print("Inserire un orario (hh:mm): ")
	fmt.Scanf("%s", &ora)
	if string(ora[2]) == ":" {
		oraint, _ = strconv.Atoi(ora[:2])
	} else if string(ora[1]) == ":" {
		oraint, _ = strconv.Atoi(string(ora[0]))
	}
	if oraint >= 6 && oraint < 13 {
		fmt.Println("mattino")
	} else {
		fmt.Println("non mattino")
	}
}
