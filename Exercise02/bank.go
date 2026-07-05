package main

import (
	"fmt"
)

func main() {
	var accountBalance float64 = 1000

	fmt.Println("Welcome to Go Bank")
	fmt.Println("What you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposity money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println(accountBalance)
	} else if choice == 2 {
		fmt.Print("Your deposit: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)

		if depositAmount <= 0 {
			fmt.Println("Invalid Amount. Must be greater then 0.")
			return
		}

		accountBalance += depositAmount //accountBalance = accountBalance + depositAmount

		fmt.Println("balance updated! new amount: ", accountBalance)
	} else if choice == 3 {
		fmt.Print("Withdraw amount: ")
		var withdrawAmount float64
		fmt.Scan(&withdrawAmount)

		if withdrawAmount <= 0 {
			fmt.Println("Invalid Withdraw. Must be greater then 0.")
			return
		}

		if withdrawAmount > accountBalance {
			fmt.Println("Insufficient funds")
			return
		}

		accountBalance -= withdrawAmount //accountBalance = accountBalance - withdrawAmount

		fmt.Println("balance updated! new amount: ", accountBalance)
	} else {
		fmt.Println("Exiting...")
	}
}
