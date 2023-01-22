package main

import "fmt"

func main() {
	var guasto1, guasto2, guasto3 int
	fmt.Print("Inserire codici di tre guasti: ")
	fmt.Scan(&guasto1, &guasto2, &guasto3)
	if guasto1 == 1 {
		fmt.Println("caricamento acqua")
	} else if guasto1 == 2 {
		fmt.Println("scarico acqua")
	} else if guasto1 == 3 {
		fmt.Println("sistema di riscaldamento")
	}
	if guasto2 == 1 {
		fmt.Println("caricamento acqua")
	} else if guasto2 == 2 {
		fmt.Println("scarico acqua")
	} else if guasto2 == 3 {
		fmt.Println("sistema di riscaldamento")
	}
	if guasto3 == 1 {
		fmt.Println("caricamento acqua")
	} else if guasto3 == 2 {
		fmt.Println("scarico acqua")
	} else if guasto3 == 3 {
		fmt.Println("sistema di riscaldamento")
	}

}
