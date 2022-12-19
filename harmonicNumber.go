package main
import "fmt"

func main(){
	var num1 int
	var sum float32
	//sum = 1.0
	fmt.Println("Enter Number of terms to add :")
	fmt.Scanln(&num1)
	fmt.Printf("%T", num1)
	for i := 1 ; i <= num1; i++ {
		sum = sum + 1.0/float32(i)
	}
	fmt.Println("Harmonic Number :", sum)
}