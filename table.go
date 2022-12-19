package main
import "fmt"

func main(){
	var n int
	fmt.Println("Enter a number")
	fmt.Scanln(&n)
	for i := 1; i <= 10; i++ {
		fmt.Println(n,"*", i, "=", n * i)
	}	
}