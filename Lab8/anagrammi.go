package main

import (
	"fmt"
	"os"
	"reflect"
)

func IsAnagram(s1, s2 string) bool {
	map1 := make(map[string]int)
	map2 := make(map[string]int)
	for i := 0; i < len(s1); i++ {
		map1[string(s1[i])]++
	}
	for j := 0; j < len(s2); j++ {
		map2[string(s2[j])]++
	}
	if reflect.DeepEqual(map1, map2) {
		return true
	} else {
		return false
	}
}
func main() {
	if IsAnagram(os.Args[1], os.Args[2]) {
		fmt.Printf("%s e %s sono anagrammi\n", os.Args[1], os.Args[2])
	} else {
		fmt.Printf("%s e %s non sono anagrammi\n", os.Args[1], os.Args[2])
	}
}
