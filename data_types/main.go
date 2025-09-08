package main
import "fmt" 
func main(){
	// Floating Point
	// var a float32
	// a = 3.14
	// fmt.Println(a)

	// Complex Numbers
	a:= 3 + 4i
	b:=complex(5,7)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(real(a))
	fmt.Println(imag(a))
	// Addition of Complex Numbers
	fmt.Println(a + b)
} 	
