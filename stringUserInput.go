package main
import (
	"fmt"
	"strings"

)

func main(){
	var place string

	fmt.Println("We are checking Whether [a, e, p] are present or not")
	fmt.Println("Enter your favourite place : ")
	fmt.Scanln(&place)

	if (strings.Contains(place, "a")) && (strings.Contains(place, "e")) && (strings.Contains(place, "p")){
		fmt.Println("All are present")
	} else if (strings.Contains(place, "a")) || (strings.Contains(place, "e")) || (strings.Contains(place, "p")) {
		fmt.Println("One or more present")
	} else{
		fmt.Println("None are present")
	}
	

}
// "os"
// "bufio"
// scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan()
// 	place := scanner.Text()

















