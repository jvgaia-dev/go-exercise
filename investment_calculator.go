package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	//fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	//fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	//fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFutureFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedFutureFRV := fmt.Sprintf("Future value (adjusted inflation): %.2f\n", futureRealValue)
	//fmt.Printf("Future Value: %.2f\nFuture value (adjusted inflation): %.2f", futureValue, futureRealValue)
	fmt.Print(formattedFutureFV, formattedFutureFRV)
}

func outputText(text string) {
	fmt.Printf(text)
}
