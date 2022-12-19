package main
import "fmt"
func main(){
	var celcius, fahrenheit, fahrenheit2, celcius2 float32
	var choice int
	var n bool
	n = true
	for n {
		fmt.Println("Select Option \n0. To Convert celcius to Fahrenheit \n1. To Convert Fahrenheit to celcius \n2.Exit")
		fmt.Scanln(&choice)
		switch choice {
		case 0:
			fmt.Println("Enter Temperature in celcius : ")
			fmt.Scanln(&celcius)
			fahrenheit = ((celcius*9)/5) + 32
			fmt.Println("Fahrenheit :", fahrenheit)	
			fmt.Println("********************************")
		case 1:
			fmt.Println("Enter Temperature in fahrenheit : ")
			fmt.Scanln(&fahrenheit2)
			celcius2 = (fahrenheit2 - 32) / 1.8
			fmt.Println("Celcius :", celcius2)
			fmt.Println("*********************************")	
		default:
			n = false
		}
	}
	
}