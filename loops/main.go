package main 
import "fmt"
func main(){
	for i:=1;i<=5;i++{ 
		fmt.Println("Iteration:",i)	
	}
	// while loop
	j:=1
	for j<=10{
		fmt.Println("Value of j:",j)
		j++
	}

	// // infinite loop 
	// for{ 
	// 	fmt.Println("Inside Infinite Loop")
	// }


	for i:=range 22{
		fmt.Println("Value from array:",i)
	}
	fmt.Println("Loop Ended")
}