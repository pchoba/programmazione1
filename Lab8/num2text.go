package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var num int
	var temp string
	numeri := make(map[int]string)
	fmt.Print("Un numero: ")
	fmt.Scan(&num)
	strnum := strconv.Itoa(num)
	if num < 0 {
		fmt.Println("Scegli un numero maggiore di zero")
		os.Exit(0)
	} else {
		for i := 0; i < len(strnum); i++ {
			temp1, _ := strconv.Atoi(string(strnum[i]))
			if numeri[temp1] == "" {
				fmt.Printf("parola per %d? ", temp1)
				fmt.Scan(&temp)
				numeri[temp1] = temp
			}
		}
	}
	for i := 0; i < len(strnum); i++ {
		temp1, _ := strconv.Atoi(string(strnum[i]))
		if i == len(strnum)-1 {
			fmt.Print(numeri[temp1])
			fmt.Println("")
		} else {
			fmt.Print(numeri[temp1], "-")
		}

	}
}
