package main

import "fmt"

func main() {
	var x1, x2, y1, y2, x3, y3 int
	fmt.Print("P1: ")
	fmt.Scan(&x1, &y1)
	fmt.Print("P2: ")
	fmt.Scan(&x2, &y2)
	fmt.Print("P3: ")
	fmt.Scan(&x3, &y3)
	if x3 > x1 && x3 < x2 {
		if y3 > y1 && y3 < y2 {
			fmt.Println("P3 interno")
		} else if y3 == y1 || y3 == y2 {
			fmt.Println("P3 perimetrale")
		} else {
			fmt.Println("P3 esterno")
		}
	} else if x3 == x1 || x3 == x2 {
		fmt.Println("P3 perimetrale")
	} else {
		fmt.Println("P3 esterno")
	}

}
