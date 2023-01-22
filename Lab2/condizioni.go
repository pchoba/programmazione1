package main

import (
	"fmt"
	"math"
)

func main() {
	var i int
	fmt.Print("Inserire un valore: ")
	fmt.Scanf("%d", &i)
	if i == 10 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	if i != 10 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	if i != 10 && i != 20 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	if i != 10 || i != 20 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	if i >= 10 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	if i >= 10 && i <= 20 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	if i > 10 && i < 20 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	if i >= 10 && i < 20 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	if i < 10 || i > 20 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	if (i >= 10 && i <= 20) || (i >= 30 && i <= 40) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	if i%4 == 0 && i%100 != 0 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	if i%2 == 1 && (i >= 0 && i >= 100) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	if math.Abs(float64(i)-10) <= 1e-6 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

}
