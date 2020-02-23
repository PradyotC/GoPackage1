package mypack1
import "fmt"
import "math"
func Square(num int) {
res:= num * num
fmt.Println("square = ", res)
}
func SquareRoot(num int){
res:= math.Sqrt(float64(num))
fmt.Println("square root = ",res)
}
func Cube(num int){
res:= num * num * num
fmt.Println("cube = ", res)
}