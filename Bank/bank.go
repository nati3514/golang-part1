package main
import "fmt"
import "os"
import "strconv"
 import "errors"

const accountBalanceFile= "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 1000, errors.New("Failed to read balance file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse balance value")
	}

	
	return balance, nil
	
}

func witeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}

func main() {
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-------------")
		
	}

	fmt.Println("Welcome to Go Bank")

	for {
	fmt.Println("What do want to do?")	
	fmt.Println("1, Check Balance")
	fmt.Println("2, Deposit")
	fmt.Println("3, Withdraw")
	fmt.Println("4, Exit")
	var choice int
	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
      fmt.Println("Your balance is", accountBalance)
	case 2:
		var depositAmount float64
		fmt.Print("Enter deposit amount: ")
		fmt.Scan(&depositAmount)

		if depositAmount < 0 {
			fmt.Println("The deposit amount must be greater than 0")
			continue
		}

		accountBalance += depositAmount
		fmt.Println("Your new balance is", accountBalance)
		witeBalanceToFile(accountBalance)
	case 3:
		var withdrawAmount float64
		fmt.Print("Enter withdraw amount: ")
		fmt.Scan(&withdrawAmount)

		if withdrawAmount > accountBalance {
			fmt.Println("Insufficient Balance")
			continue
		}
		
			accountBalance -= withdrawAmount
			fmt.Println("Your new balance is", accountBalance)
			witeBalanceToFile(accountBalance)
	case 4:
		fmt.Println("Exiting...")
		return
	default:
		fmt.Println("Invalid choice")
		return
	}
 // fmt.Println("Thank you for using Go Bank")
}
}
