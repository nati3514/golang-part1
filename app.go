package main
import "fmt"
import "math"

const inflationRate = 2.5

func main() {
	
	var investmentAmount float64
	years := 10.0
	expectedAmount := 5.5

    outputText("Enter Investment Amount:")
	fmt.Scan(&investmentAmount)
	outputText("Enter Expected Amount:")
	fmt.Scan(&expectedAmount)
	outputText("Enter Years")
	fmt.Scan(&years)

	futureValue, futureRateValue := CalculaeFutureValue(investmentAmount, expectedAmount, years)
	

	// futureValue := investmentAmount * math.Pow(1 + expectedAmount/100, years)
	fmt.Println("Future Value: ", futureValue)

	

	// futureRateValue := futureValue / math.Pow(1 + inflationRate/100, years)
	fmt.Println("Future Rate Value: ", futureRateValue)


}

func outputText(Text string) {
	fmt.Println(Text)
}

func CalculaeFutureValue(investmentAmount, expectedAmount, years float64) (futureValue float64, futureRateValue float64) {
	futureValue = investmentAmount * math.Pow(1 + expectedAmount/100, years)
	futureRateValue = futureValue / math.Pow(1 + inflationRate/100, years)

	return futureValue, futureRateValue

}
