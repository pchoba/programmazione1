package main
import "fmt"
func main() {
	var a, b float64
	const gradi_triangolo = 180.0
	fmt.Print("angolo uno e due: ")
	fmt.Scan(&a, &b)
	fmt.Println("ampiezza terzo angolo: ", (gradi_triangolo - a - b))
}
