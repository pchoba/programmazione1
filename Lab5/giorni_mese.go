package main
import "fmt"

func hasUpper(s string) bool {
	for i:=0; i<len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			return true
		}
	}
	return false
	}

func firstUpper(s string) int {
	for i:=0; i<len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			return i
		}
	}
	return -1
	}

	func lastUpper(s string) int {
		for i:=len(s)-1; i>=0; i-- {
			if s[i] >= 'A' && s[i] <= 'Z' {
				return i
			}
		}
		return -1
		}

func main() {
	fmt.Println(lastUpper("ldAwq"))
}
