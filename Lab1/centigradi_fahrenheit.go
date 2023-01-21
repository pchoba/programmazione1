package main

import "fmt"

func main() {
	var temp1, temp2 float64
	fmt.Print("temperatura in centigradi? ")
	fmt.Scan(&temp1)
	temp2 = 32 + ((temp1 * 9) / 5)
	fmt.Println(temp1, "C = ", temp2, "F")
}
