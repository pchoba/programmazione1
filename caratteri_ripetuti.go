//data una stringa stabilire se ci sono caratteri ripetuti
package main
import "fmt"

CaratteriRipetuti(s string) bool {
    for i:=0; i<=len(s); i++ {
        for j:=i + 1; j< len(s); j++ {
            if s[i] == s[j] {
                return true
            }    
        }
    }
    return false
}
