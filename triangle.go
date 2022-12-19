package main
import "fmt"

func main(){
	var x, y, z int
	fmt.Println("Enter the value of 1st Side :")
	fmt.Scanln(&x)
	fmt.Println("Enter the value of 2nd Side :")
	fmt.Scanln(&y)
	fmt.Println("Enter the value of 3rd Side :")
	fmt.Scanln(&z)
	if(x == y) && (x == z) {
		fmt.Println("Triangle is Equilateral triangle")
	} else if (x == y) && (x == z) && (y == z){
		fmt.Println("Triangle is isosceles triangle")
	} else {
		fmt.Println("Triangle is scalene triangle")
	}
}