package main

import "fmt"

func main(){
	var loan_amount, intrest_rate, num_of_year, tottal_amount, monthly_amount float32
	fmt.Println("Enter loan Amount :")
	fmt.Scanln(&loan_amount)
	fmt.Println("Enter the intrest rate :")
	fmt.Scanln(&intrest_rate)
	fmt.Println("The number of years")
	fmt.Scanln(&num_of_year)

	tottal_amount = (num_of_year * loan_amount)+(num_of_year*loan_amount*(intrest_rate/100.00))
	monthly_amount = tottal_amount/(num_of_year*12)

	fmt.Println("Tottal amount to be paid :", tottal_amount)
	fmt.Println("Tottal intrest :", (tottal_amount -(num_of_year * loan_amount)))
	fmt.Println("Monthly Amount to be paid : ",monthly_amount)
}