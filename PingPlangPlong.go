package main
import "fmt"

func main(){
	var number int
	fmt.Println("Enter Number")
	fmt.Scanln(&number)
	if (number % 3 == 0) || (number % 5 == 0){
		fmt.Println("PlingPlang")
	} else if (number % 7 == 0) || (number % 3 == 0) {
		fmt.Println("PlongPling")
	} else if (number % 5 == 0) || (number % 7 == 0) {
		fmt.Println("PlangPlong")
	} else if number % 3 == 0 {
		fmt.Println("Pling")
	} else if number % 5 == 0 {
		fmt.Println("Plang")
	} else if number % 7 == 0 {
		fmt.Println("Plong")
	} else {
		fmt.Println("No match : ", number)
	}
}