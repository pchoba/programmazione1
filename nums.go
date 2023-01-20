package main
import "fmt"
func main() {
	var x, y, z int
	x = 1
	y = x + (2 * z) + x
	x = y + x
	z = x + (2 * y) + x
	x,y = z,x
	y,z,x = x,y,z+1
	fmt.Println(x, y, z)  
}