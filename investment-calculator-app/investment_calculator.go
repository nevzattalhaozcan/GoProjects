package main

import (
	"fmt"
	"math"
)

func main() {

	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	//fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	//fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	//fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)

	// FV=PV×(1+r)^n
	//futureValue := investmentAmount * math.Pow(( 1 + expectedReturnRate),years)

	// Real Future Value=FV×(1/(1+inflationRate)^n)
	//realFutureValue := futureValue * (1 / math.Pow((1 + inflationRate), years))

	futureValue, realFutureValue := calculateFutureValues(investmentAmount, expectedReturnRate, inflationRate, years)

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Real Future Value: %.2f\n", realFutureValue)

	outputText(formattedFV)
	outputText(formattedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, inflationRate, years float64) (float64, float64) {

	futureValue := investmentAmount * math.Pow((1+expectedReturnRate), years)
	realFutureValue := futureValue * (1 / math.Pow((1+inflationRate), years))

	return futureValue, realFutureValue
}
