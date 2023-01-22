package main

import "fmt"

func main() {
	var gg1, str1, end1, gg2, str2, end2 int
	fmt.Print("appuntamento 1 (gg, start, end): ")
	fmt.Scan(&gg1, &str1, &end1)
	fmt.Print("appuntamento 2 (gg, start, end): ")
	fmt.Scan(&gg2, &str2, &end2)
	if gg1 != gg2 {
		fmt.Println("non si sovrappongono")
	} else if str2 > str1 && str2 < end1 {
		fmt.Println("si sovrappongono")
	}
}
