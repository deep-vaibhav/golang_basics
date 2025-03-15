package main

import (
	"fmt"

	"example.com/bank/fileops"
)

func main() {
	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("Choose an option:")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		balance, err := fileops.ReadFloatFromFile("balance.txt")

		if err != nil {
			fmt.Println("ERROR")
			fmt.Println(err)
			fmt.Println("--------")
		}

		var choice int
		fmt.Scan(&choice)
		fmt.Println("Your choice: ", choice)

		switch choice {
		case 1:
			fmt.Println("Balance: ", balance)
		case 2:
			fmt.Print("Deposit amount: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount!")
				continue
			}

			balance += depositAmount
			fmt.Println("Balance updated! New amount: ", balance)
			fileops.WriteFloatToFile(balance, "balance.txt")
		case 3:
			fmt.Print("Withdraw amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount!")
				continue
			}

			if withdrawAmount > balance {
				fmt.Println("Insufficient balance!")
				continue
			}

			balance -= withdrawAmount
			fmt.Println("Balance updated! New amount: ", balance)
			fileops.WriteFloatToFile(balance, "balance.txt")
		default:
			fmt.Println("Thank you!")
			return
		}
	}
}
