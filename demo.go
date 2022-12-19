package main
import "fmt"
package main
import (
    "fmt"
    "log"
    )
// func mulFloatNumbers(numb1, numb2 float32) float32 {
// 	return numb1 * numb2
//   }
  
//   func main() {
// 	 fmt.Println("1.5 X 5.0 = ", mulFloatNumbers(6.2, 3.2))
//   }

func swap(strng1, strng2 string) (string, string) {
	return strng2, strng1
 }
 
 
 func main() {
	 
	lName, fName := swap("Priyanka", "Bachhav")
	
	fmt.Println(lName, fName)
 }




func main() {
    fmt.Println("Logging examples: ")
    log.Println("We are logging in Golang!")
    //log.Panic("A panic call ...!")
    //log.Fatal("Exception occured!")
    log.Println("Are we reaching here?")
}