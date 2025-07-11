package main
import "fmt"
	
    

func main() {
	
    revenue := getuserInput("Enter Revenue: ")
	expenses := getuserInput("Enter Expenses:")
	taxRate := getuserInput("Enter Tax Rate: ")

	EBT, profit, ratio := Calculaes(revenue, expenses, taxRate)
	// EBT := revenue - expenses
	outputTextNumber("Earnings Before Tax: %.1f \n", EBT)

	// profit := EBT * (1 - taxRate/100)
	outputTextNumber("Profit: %.1f \n", profit)

	// ratio := EBT / profit
	outputTextNumber("Profit Ratio:%.2f", ratio)

}

func getuserInput(Text string)  float64 {
	var userInput float64
	fmt.Println(Text)
	fmt.Scan(&userInput)
	return userInput
}

func outputTextNumber(Text string, Number float64) {
	fmt.Printf(Text, Number)
}
func Calculaes(revenue, expenses, taxRate float64) (EBT, profit, ratio float64) {
	EBT = revenue - expenses
	profit = EBT * (1 - taxRate/100)
	ratio = EBT / profit
	return EBT, profit, ratio
}