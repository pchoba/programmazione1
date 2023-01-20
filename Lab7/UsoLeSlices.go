package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	var slice1, slice2 []string

	slice1 = os.Args[1:]
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Inserire almeno tre stringhe nella slice2: ")
	for scanner.Scan() {
		slice_temp := strings.Split(scanner.Text(), " ")
		slice2 = append(slice2, slice_temp...)
	}
	slice1 = append(slice1, slice2...)
	fmt.Println(slice1)
	sort.Strings(slice1)
	fmt.Println(slice1)
	slice1 = slice1[:len(slice1)-1]
	fmt.Println(slice1)
	slice1 = append(slice1[:2], slice1[4:]...)
	fmt.Println(slice1)
	//slice3 := []string{"a", "b", "c"}
}
