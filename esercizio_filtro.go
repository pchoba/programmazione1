package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		if i%2 == 1 {
			temp, _ := strconv.Atoi(os.Args[i+1])
			for j := 0; j < temp; j++ {
				fmt.Print(os.Args[i])
			}
			fmt.Print(" ")
		}
	}
	fmt.Println("")
}
