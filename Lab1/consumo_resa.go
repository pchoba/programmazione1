package main

import "fmt"

func main() {
	var km, carb float64
	fmt.Print("Distanza percorsa (in km): ")
	fmt.Scan(&km)
	fmt.Print("Quantità di carburante utilizzato (in l): ")
	fmt.Scan(&carb)
	cons_medio := carb / km
	resa := km / carb
	fmt.Println("consumo medio:", cons_medio, "l/km")
	fmt.Println("resa media:", resa, "km/l")
}
