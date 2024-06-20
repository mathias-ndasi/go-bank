package main

import "fmt"

func main() {
	var accountBalance float64 = 1000

	fmt.Println("Welcome to Go Bank!!!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

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
			}
		default:
			fmt.Println("Exiting. Thank you for Banking with us :) !!!")
			return
		}

		fmt.Println("")
	}
}
