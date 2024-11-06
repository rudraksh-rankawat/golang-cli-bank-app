package main

import (
	"fmt"
	"example1.com/bank/file_utility"
)

var balanceFile = "balance.txt"


func main() {
	fmt.Println("Welcome to the CLI Bank...")
	accountBalance, err := file_utility.ReadFloatFromFile(balanceFile)
	if err != nil {
		fmt.Println("--------")
		fmt.Println("ERROR!")
		fmt.Println(err)
		panic("Unable to fetch the balance...")
	}

	for {

		printMain()

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
			file_utility.WriteFloatToFile(accountBalance, balanceFile)
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
			file_utility.WriteFloatToFile(accountBalance, balanceFile)
			fmt.Printf("Here is your updated balance: %f\n\n", accountBalance)
		} else {

			fmt.Printf("Thank You for choosing our services !!\n\n")
			break
		}

	}

}
