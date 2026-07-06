package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() float64 {
	data, _ := os.ReadFile(accountBalanceFile)
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)

	return balance
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	var accountBalance float64 = getBalanceFromFile()

	fmt.Println("Welcome to Go Bank")

	for {
		fmt.Println("--------------------------")
		fmt.Println("What you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposity money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println(accountBalance)

		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid Amount. Must be greater then 0.")
				continue
			}

			accountBalance += depositAmount //accountBalance = accountBalance + depositAmount
			writeBalanceToFile(accountBalance)

			fmt.Println("balance updated! new amount: ", accountBalance)

		case 3:
			fmt.Print("Withdraw amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid Withdraw. Must be greater then 0.")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient funds")
				continue
			}

			accountBalance -= withdrawAmount //accountBalance = accountBalance - withdrawAmount
			writeBalanceToFile(accountBalance)

			fmt.Println("balance updated! new amount: ", accountBalance)

		case 4:
			fmt.Println("Exiting...")
			fmt.Println("Thanks for choosing our bank")
			return
			//break

		default:
			fmt.Println("Invalid input")
		}
	}
}
