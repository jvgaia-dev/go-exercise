package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Percentage of Tax Rate: ")
	fmt.Scan(&taxRate)
	taxRate = taxRate / 100

	earningsBeforeTax := revenue - expenses
	taxAmount := earningsBeforeTax * taxRate
	earningsAfterTax := earningsBeforeTax - taxAmount
	ratio := earningsBeforeTax / earningsAfterTax

	fmt.Println("Earnings before tax: ", earningsBeforeTax)
	fmt.Println("Earnings after tax: ", earningsAfterTax)
	fmt.Println("Ratio: ", ratio)
}
