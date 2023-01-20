package main

import "fmt"

const DIM = 69

func main() {
	var ciao [DIM]int
	for i := 0; i < DIM; i++ {
		ciao[i] = i + 1
	}
	fmt.Println(ciao)

	// inverse array
	for i := 0; i < DIM/2; i++ {
		ciao[DIM-i-1], ciao[i] = ciao[i], ciao[DIM-i-1]
	}
	fmt.Println(ciao)
	// first to last
	ciao[0], ciao[DIM-1] = ciao[DIM-1], ciao[0]
	fmt.Println(ciao)

}
