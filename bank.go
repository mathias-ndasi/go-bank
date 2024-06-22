package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"

	fileOperations "example.com/bank/file-operations"
)

const accountBalanceFileName = "account_balance.txt"

func main() {
	var accountBalance, error = fileOperations.GetFloatFromFile(accountBalanceFileName)

	if error != nil {
		fmt.Println("ERROR:", error)
		fmt.Println("----------------------------------------------------------------")
		panic(error)
	}

	fmt.Println("Welcome to Go Bank!!!")
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())

	for {
		presentOptions()
		var choice int

		fmt.Print("Your choice: ")
		fmt.Scanln(&choice)

		if choice <= 0 || choice > 4 {
			fmt.Printf("Invalid choice\n\n")
			continue
		}

		switch choice {
		case 1:
			fmt.Println("Checking balance...")
			fmt.Println("Your account balance is: ", accountBalance)
		case 2:
			fmt.Println("Depositing money...")
			fmt.Print("How much do you want to deposit: ")

			var depositAmount float64
			fmt.Scanln(&depositAmount)

			if depositAmount <= 0 {
				fmt.Printf("Invalid deposit amount. Amount must be greater than zero\n\n")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Your account balance is: ", accountBalance)

			fileOperations.WriteFloatToFile(accountBalance, accountBalanceFileName)
		case 3:
			fmt.Println("Withdrawing money...")
			fmt.Print("How much do you want to withdraw: ")

			var withdrawalAmount float64
			fmt.Scanln(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Printf("Invalid withdrawal amount. Amount must be greater than zero\n\n")
				continue
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("Insufficient funds")
			} else {
				accountBalance -= withdrawalAmount
				fmt.Println("Your account balance is: ", accountBalance)

				fileOperations.WriteFloatToFile(accountBalance, accountBalanceFileName)
			}
		default:
			fmt.Println("Exiting. Thank you for Banking with us :) !!!")
			return
		}

		fmt.Println("")
	}
}
