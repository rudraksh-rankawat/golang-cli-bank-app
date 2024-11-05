package main

import (
	"fmt"
	"os"
)

var balanceFile = "balance.txt"

func readFromFile() float64 {
	
}

func writeToFile(balance float64) {
	data := fmt.Sprintf("%f", balance)
	os.WriteFile(balanceFile, []byte(data), 0644)
	// log.Println("Balance successfully recorded")
}

func main() {
	fmt.Println("Welcome to the CLI Bank...")
	var accountBalance float64 = 1000
	for {

		fmt.Println("Please choose an option:")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter option: ")
		fmt.Scan(&choice)

		if choice <= 0 || choice > 4 {
			fmt.Printf("-- Invalid Choice --\n\n")
			continue
		}

		if choice == 1 {
			fmt.Print("Here is the remaining balance in your account: ")
			fmt.Printf("%v\n\n", accountBalance)
			continue

		} else if choice == 2 {
			var deposit float64
			print("Enter the amount to be deposited: ")
			fmt.Scan(&deposit)
			if deposit <= 0 {
				fmt.Printf("-- Please enter the amount above 0 --\n\n")
				continue
			}
			accountBalance += deposit
			writeToFile(accountBalance)
			fmt.Printf("Here is your updated balance: ")
			fmt.Printf("%f\n\n", accountBalance)

		} else if choice == 3 {

			var withdrawed float64
			print("Enter the amount to be withdrawn: ")
			fmt.Scan(&withdrawed)
			if withdrawed > accountBalance {
				fmt.Printf("-- Please enter the amount less than the current balance --\n")
				fmt.Printf("-- Current Balance: %f --\n\n", accountBalance)
				continue
			}
			accountBalance -= withdrawed
			writeToFile(accountBalance)
			fmt.Printf("Here is your updated balance: %f\n\n", accountBalance)
		} else {

			fmt.Printf("Thank You for choosing our services !!\n\n")
			break
		}

	}

}
