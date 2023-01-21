package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var max int
	var slice1, slice2 []string
	if len(os.Args) == 1 || len(os.Args) == 2 {
		fmt.Println("Inserire DUE nomi di file")
	} else {
		f1, _ := os.Open(os.Args[1])
		f2, _ := os.Open(os.Args[2])
		defer f1.Close()
		defer f2.Close()
		scanner1 := bufio.NewScanner(f1)
		scanner2 := bufio.NewScanner(f2)
		for scanner1.Scan() {
			text1 := scanner1.Text()
			slice1 = strings.Split(text1, " ")
		}
		for scanner2.Scan() {
			text2 := scanner2.Text()
			slice2 = strings.Split(text2, " ")
		}
		if len(slice1) >= len(slice2) {
			max = len(slice1)
		} else {
			max = len(slice2)
		}
		for i := 0; i < max; i++ {
			if i > len(slice1)-1 {
				fmt.Println(slice2[i])
			} else if i > len(slice2)-1 {
				fmt.Println(slice1[i])
			} else {
				fmt.Println(slice1[i])
				fmt.Println(slice2[i])
			}
		}
	}

}
