package main

import (
	"fmt"
)

func main() {
	//var revenue float64
	//var expenses float64
	//var taxRate float64

	//fmt.Print("Revenue: ")
	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Percentage of Tax Rate: ")

	//fmt.Print("Expenses: ")
	//fmt.Scan(&expenses)

	//fmt.Print("Percentage of Tax Rate: ")
	//fmt.Scan(&taxRate)

	ebt, profit, ratio := taxCalculate(revenue, expenses, taxRate)

	fmt.Println("-------------")
	fmt.Println("Earnings before tax: ", ebt)
	fmt.Println("Earnings after tax: ", profit)
	fmt.Println("Ratio: ", ratio)
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

func taxCalculate(revenue float64, expenses float64, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}
